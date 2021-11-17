package manager

import (
	"errors"
	"github.com/spf13/afero"
	"github.com/tcw/ibsen/access"
	"github.com/tcw/ibsen/errore"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type TopicHandler struct {
	afs            *afero.Afero
	rootPath       string
	topic          access.Topic
	maxBlockSize   access.BlockSizeInBytes
	headBlockSize  access.BlockSizeInBytes
	logBlocks      access.Blocks
	logOffset      access.Offset
	logMutex       sync.Mutex
	indexBlocks    access.Blocks
	indexOffset    access.Offset
	indexMutex     int32
	logAccess      access.LogAccess
	logIndexAccess access.LogIndexAccess
	loaded         bool
}

func NewTopicHandler(afs *afero.Afero, rootPath string, topic access.Topic, maxBlockSize uint64) TopicHandler {
	return TopicHandler{
		afs:          afs,
		rootPath:     rootPath,
		maxBlockSize: access.BlockSizeInBytes(maxBlockSize),
		topic:        topic,
		logOffset:    0,
		logMutex:     sync.Mutex{},
		indexOffset:  0,
		indexMutex:   0,
		logAccess: access.ReadWriteLogAccess{
			Afs:      afs,
			RootPath: rootPath,
		},
		logIndexAccess: access.ReadWriteLogIndexAccess{
			Afs:          afs,
			RootPath:     rootPath,
			IndexDensity: 0.01,
		},
		loaded: false,
	}
}

func (t *TopicHandler) Write(entries access.Entries) (uint32, error) {
	t.logMutex.Lock()
	defer t.logMutex.Unlock()
	err := t.Load()
	if err != nil {
		return 0, errore.WrapWithContext(err)
	}
	if t.logBlocks.IsEmpty() {
		t.logBlocks.AddBlock(0)
	}
	offset, bytes, err := t.logAccess.Write(t.logBlocks.Head().LogFileName(t.rootPath, t.topic), entries, t.logOffset)
	if err != nil {
		return 0, errore.WrapWithContext(err)
	}
	t.logOffset = offset
	t.headBlockSize = t.headBlockSize + bytes
	if t.headBlockSize > t.maxBlockSize {
		t.logBlocks.AddBlock(access.Block(t.logOffset))
	}
	return uint32(bytes), nil
}

func (t *TopicHandler) Read(params access.ReadParams) (access.Offset, error) {
	err := t.Load()
	if err != nil {
		return 0, errore.WrapWithContext(err)
	}
	blocks := t.logBlocks.GetBlocks(t.logOffset)
	if len(blocks) == 0 {
		return 0, errore.WrapWithContext(errors.New("no blocks"))
	}
	var lastReadOffset access.Offset = 0
	for _, block := range blocks {
		logFileName := block.LogFileName(t.rootPath, t.topic)
		exists, err := t.afs.Exists(string(logFileName))
		if err != nil {
			return 0, errore.WrapWithContext(err)
		}
		if !exists {
			return lastReadOffset, nil
		}
		lastReadOffset, err = t.logAccess.ReadLog(logFileName, params, 0)
		if err != nil {
			return 0, errore.WrapWithContext(err)
		}
	}
	return lastReadOffset, nil
}

func (t *TopicHandler) Load() error {
	if !t.loaded {
		t.logMutex.Lock()
		defer t.logMutex.Unlock()
		if !t.loaded {
			err := t.lazyLoad()
			if err != nil {
				return errore.WrapWithContext(err)
			}
			t.loaded = true
			go t.indexScheduler()
		}
	}
	return nil
}

func (t *TopicHandler) lazyLoad() error {
	topicPath := t.rootPath + access.Sep + string(t.topic)
	exists, err := t.afs.Exists(topicPath)
	if err != nil {
		return errore.WrapWithContext(err)
	}
	if !exists {
		err = t.afs.Mkdir(topicPath, 640)
		if err != nil {
			return errore.WrapWithContext(err)
		}
	}
	blocks, err := t.logAccess.ReadTopicLogBlocks(t.topic)
	if err != nil {
		return errore.WrapWithContext(err)
	}
	t.logBlocks = blocks
	indexBlocks, err := t.logIndexAccess.ReadTopicIndexBlocks(t.topic)
	if err != nil {
		return errore.WrapWithContext(err)
	}
	t.indexBlocks = indexBlocks
	return nil
}

func (t *TopicHandler) indexScheduler() {
	for {
		err := t.updateIndex()
		if err != nil {
			log.Printf("index builder for topic %s has failed", t.topic)
		}
		time.Sleep(time.Second * 10)
	}
}

func (t *TopicHandler) updateIndex() error {
	if !atomic.CompareAndSwapInt32(&t.indexMutex, 0, 1) {
		return nil
	}
	defer atomic.CompareAndSwapInt32(&t.indexMutex, 1, 0)
	head := t.indexBlocks.Head()
	if head == 0 {
		return nil
	}
	index, err := t.logIndexAccess.Read(head.IndexFileName(t.rootPath, t.topic))
	if err != nil {
		return errore.WrapWithContext(err)
	}
	indexOffset := index.Head()
	if !indexOffset.IsEmpty() {
		_, err = t.logIndexAccess.WriteFromOffset(head.IndexFileName(t.rootPath, t.topic), indexOffset.ByteOffset)
		if err != nil {
			return errore.WrapWithContext(err)
		}
	}
	blocksWithoutIndex := t.logBlocks.Diff(t.indexBlocks)
	for _, block := range blocksWithoutIndex.BlockList {
		_, err := t.logIndexAccess.WriteFile(block.IndexFileName(t.rootPath, t.topic))
		if err != nil {
			return errore.WrapWithContext(err)
		}
	}
	return nil
}
