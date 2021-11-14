package cmd

import (
	"fmt"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/tcw/ibsen/api"
	"github.com/tcw/ibsen/consensus"
	"github.com/tcw/ibsen/errore"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	host           string
	serverPort     int
	maxBlockSizeMB int

	entries int

	cpuProfile string
	memProfile string

	rootCmd = &cobra.Command{
		Use:              "ibsen",
		Short:            "Ibsen is a simple log streaming system",
		Long:             `Ibsen is a simple log streaming system build on unix's philosophy of simplicity'`,
		TraverseChildren: true,
	}

	cmdServer = &cobra.Command{
		Use:              "server [data directory]",
		Short:            "start a ibsen server",
		Long:             `server`,
		TraverseChildren: true,
		Args:             cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			inMemory := false
			var afs *afero.Afero
			absolutePath := "/tmp/data"
			if len(args) == 0 {
				var fs = afero.NewMemMapFs()
				afs = &afero.Afero{Fs: fs}
				inMemory = true
			} else {
				var err error
				if len(args) != 1 {
					fmt.Println("No file path to ibsen storage was given, use -i to run in-memory or specify path")
					return
				}
				absolutePath, err = filepath.Abs(args[0])
				if err != nil {
					log.Fatal(err)
				}
				var fs = afero.NewOsFs()
				afs = &afero.Afero{Fs: fs}
			}
			writeLock := absolutePath + string(os.PathSeparator) + ".writeLock"
			lock := consensus.NewFileLock(afs, writeLock, time.Second*20)
			ibsenServer := api.IbsenServer{
				Lock:         lock,
				InMemory:     inMemory,
				Afs:          afs,
				RootPath:     absolutePath,
				MaxBlockSize: maxBlockSizeMB,
				CpuProfile:   cpuProfile,
				MemProfile:   memProfile,
			}
			lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, serverPort))
			if err != nil {
				log.Println(errore.SprintTrace(err))
				return
			}
			err = ibsenServer.Start(lis)
			if err != nil {
				log.Fatalf(errore.SprintTrace(err))
			}
		},
	}

	cmdFile = &cobra.Command{
		Use:              "file",
		Short:            "access files directly from file",
		Long:             `file`,
		TraverseChildren: true,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	serverPort, _ = strconv.Atoi(getenv("IBSEN_PORT", strconv.Itoa(50001)))
	host = getenv("IBSEN_HOST", "0.0.0.0")
	maxBlockSizeMB, _ = strconv.Atoi(getenv("IBSEN_MAX_BLOCK_SIZE", "100"))
	entries, _ = strconv.Atoi(getenv("IBSEN_ENTRIES", "1000"))
	rootCmd.Flags().IntVarP(&serverPort, "port", "p", serverPort, "config file (default is current directory)")
	rootCmd.Flags().StringVarP(&host, "host", "l", "0.0.0.0", "config file (default is current directory)")
	cmdServer.Flags().IntVarP(&maxBlockSizeMB, "maxBlockSize", "m", maxBlockSizeMB, "Max MB in log files")
	cmdServer.Flags().StringVarP(&cpuProfile, "cpuProfile", "z", "", "config file")
	cmdServer.Flags().StringVarP(&memProfile, "memProfile", "y", "", "config file")
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
