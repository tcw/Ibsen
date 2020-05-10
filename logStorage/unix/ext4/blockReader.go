package ext4

import (
	"bufio"
	"errors"
	"github.com/tcw/ibsen/logStorage"
	"io"
	"log"
	"os"
	"sync"
)

func OpenFileForRead(fileName string) (*os.File, error) {

	f, err := os.OpenFile(fileName,
		os.O_RDONLY, 0400)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return f, nil
}

func FastForwardToOffset(file *os.File, offset int64) error {
	var offsetFound int64 = -1
	for {
		if offsetFound == offset {
			return nil
		}
		bytes := make([]byte, 8)
		n, err := file.Read(bytes)
		if err == io.EOF {
			return errors.New("no offset in block")
		}
		if err != nil {
			return errors.New("error")
		}
		if n != 8 {
			log.Println("offset incorrect")
			return errors.New("offset incorrect")
		}
		offsetFound = int64(fromLittleEndian(bytes))

		n, err2 := file.Read(bytes)
		if n != 8 {
			log.Println("offset incorrect")
			return errors.New("offset incorrect")
		}
		size := fromLittleEndian(bytes)
		if err2 != nil {
			log.Println(err2)
			return err2
		}
		_, err = file.Seek(int64(size), 1)
		if err != nil {
			println(err)
			return err
		}
	}
}

func ReadLogBlockFromOffsetNotIncluding(file *os.File, logChan chan logStorage.LogEntryBatch, wg *sync.WaitGroup, offset uint64, batchSize int) error {

	err := FastForwardToOffset(file, int64(offset))
	if err != nil {
		return err
	}

	partialBatch, _, err := ReadLogInBatchesToEnd(file, nil, logChan, wg, batchSize)
	if err != nil {
		return err
	}
	wg.Add(1)
	logChan <- logStorage.LogEntryBatch{Entries: partialBatch.Entries}
	return nil
}

func ReadLogInBatchesToEnd(file *os.File, partialBatch []logStorage.LogEntry, logChan chan logStorage.LogEntryBatch,
	wg *sync.WaitGroup, batchSize int) (logStorage.LogEntryBatch, bool, error) {

	hasSent := false
	reader := bufio.NewReader(file)
	bytes := make([]byte, 8)
	var entryBatch []logStorage.LogEntry
	if partialBatch != nil {
		entryBatch = append(entryBatch, partialBatch...)
	}
	for {

		if entryBatch != nil && len(entryBatch)%batchSize == 0 {
			wg.Add(1)
			logChan <- logStorage.LogEntryBatch{Entries: entryBatch}
			hasSent = true
			entryBatch = nil
		}

		n, err := io.ReadFull(reader, bytes)

		if err == io.EOF {
			return logStorage.LogEntryBatch{Entries: entryBatch}, hasSent, nil
		}
		if err != nil {
			log.Println(err)
			return logStorage.LogEntryBatch{}, false, err
		}
		if n != 8 {
			log.Println("offset incorrect")
			return logStorage.LogEntryBatch{}, false, errors.New("offset incorrect")
		}
		offset := int64(fromLittleEndian(bytes))

		n, err2 := io.ReadFull(reader, bytes)
		if n != 8 {
			log.Println("entry size incorrect")
			return logStorage.LogEntryBatch{}, false, errors.New("entry size incorrect")
		}
		size := fromLittleEndian(bytes)
		if err2 != nil {
			log.Println(err2)
			return logStorage.LogEntryBatch{}, false, err2
		}
		entry := make([]byte, size)
		n, err3 := io.ReadFull(reader, entry)
		if err3 != nil {
			log.Println(err3)
			return logStorage.LogEntryBatch{}, false, err3
		}
		if n != int(size) {
			log.Println("entry incorrect")
			return logStorage.LogEntryBatch{}, false, errors.New("entry incorrect")
		}
		entryBatch = append(entryBatch, logStorage.LogEntry{
			Offset:   uint64(offset),
			ByteSize: int(size),
			Entry:    entry,
		})
	}
}

func ReadLogFromOffsetNotIncluding(file *os.File, c chan logStorage.LogEntry, wg *sync.WaitGroup, offset uint64) error {
	err := FastForwardToOffset(file, int64(offset))
	if err != nil {
		return err
	}

	err = ReadLogToEnd(file, c, wg)
	if err != nil {
		return err
	}
	return nil
}

func ReadLogToEnd(file *os.File, c chan logStorage.LogEntry, wg *sync.WaitGroup) error {
	reader := bufio.NewReader(file)
	bytes := make([]byte, 8)
	for {
		n, err := io.ReadFull(reader, bytes)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}
		if n != 8 {
			log.Println("offset incorrect")
			return errors.New("offset incorrect")
		}
		offset := fromLittleEndian(bytes)
		n, err2 := io.ReadFull(reader, bytes)
		if n != 8 {
			log.Println("entry size incorrect")
			return errors.New("entry size incorrect")
		}
		size := fromLittleEndian(bytes)
		if err2 != nil {
			log.Println(err2)
			return err2
		}
		entry := make([]byte, size)
		n, err3 := io.ReadFull(reader, entry)
		if err3 != nil {
			log.Println(err3)
			return err3
		}
		if n != int(size) {
			log.Println("entry incorrect")
			return errors.New("entry incorrect")
		}
		wg.Add(1)
		c <- logStorage.LogEntry{
			Offset:   offset,
			ByteSize: int(size),
			Entry:    entry,
		}
	}
}
