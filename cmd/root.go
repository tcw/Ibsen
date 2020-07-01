package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tcw/ibsen/client"
	"github.com/tcw/ibsen/server"
	"github.com/tcw/ibsen/tools"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	useHttp        bool
	host           string
	serverPort     int
	maxBlockSizeMB int
	cpuprofile     string
	memprofile     string

	rootCmd = &cobra.Command{
		Use:   "ibsen",
		Short: "Ibsen is a kubernetes friendly streaming platform",
		Long: `Ibsen builds on the shoulders of giants. Taking advantage of the recent advances in 
			distributed block storage and unix's philosophy of simplicity'`,
		TraverseChildren: true,
	}

	cmdServer = &cobra.Command{
		Use:              "server",
		Short:            "server",
		Long:             `server`,
		TraverseChildren: true,
	}

	cmdServerStart = &cobra.Command{
		Use:              "start [data directory]",
		Short:            "start",
		Long:             `start`,
		TraverseChildren: true,
		Args:             cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			absolutePath, err2 := filepath.Abs(args[0])
			if err2 != nil {
				log.Fatal(err2)
			}
			ibsenServer := server.IbsenServer{
				DataPath:     absolutePath,
				UseHttp:      useHttp,
				Host:         host,
				Port:         serverPort,
				MaxBlockSize: maxBlockSizeMB,
				CpuProfile:   "",
				MemProfile:   "",
			}
			ibsenServer.Start()
		},
	}

	cmdClient = &cobra.Command{
		Use:   "client [from directory]",
		Short: "client",
		Long:  `client`,
	}

	cmdClientCreate = &cobra.Command{
		Use:   "create [topic]",
		Short: "create",
		Long:  `create`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ibsenClient := startClient()
			ibsenClient.CreateTopic(args[0])
		},
	}

	cmdClientRead = &cobra.Command{
		Use:   "read [topic]",
		Short: "read",
		Long:  `read`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ibsenClient := startClient()
			ibsenClient.ReadTopic(args[0])
		},
	}

	cmdClientWrite = &cobra.Command{
		Use:   "write [topic]",
		Short: "write",
		Long:  `write`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ibsenClient := startClient()
			ibsenClient.WriteTopic(args[0])
		},
	}

	cmdClientBench = &cobra.Command{
		Use:   "bench",
		Short: "bench",
		Long:  `benchmark`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	cmdCat = &cobra.Command{
		Use:   "cat [from file/directory]",
		Short: "cat",
		Long:  `cat`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tools.ReadTopic(args[0])
		},
	}

	cmdCheck = &cobra.Command{
		Use:   "check [from directory]",
		Short: "check",
		Long:  `check`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}
)

func startClient() client.IbsenClient {
	return client.Start(host + ":" + strconv.Itoa(serverPort))
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	cmdServer.Flags().BoolVarP(&useHttp, "http", "u", true, "config file (default is current directory)")
	cmdServer.Flags().Lookup("http").NoOptDefVal = "true"
	rootCmd.Flags().IntVarP(&serverPort, "port", "p", 50001, "config file (default is current directory)")
	rootCmd.Flags().StringVarP(&host, "host", "l", "localhost", "config file (default is current directory)")
	cmdServer.Flags().IntVarP(&maxBlockSizeMB, "maxBlockSize", "b", 100, "config file (default is current directory)")
	cmdServer.Flags().StringVarP(&cpuprofile, "cpuprofile", "c", "", "config file (default is current directory)")
	cmdServer.Flags().StringVarP(&memprofile, "memprofile", "m", "", "config file (default is current directory)")
	err := cmdServer.Flags().MarkHidden("cpuprofile")
	if err != nil {
		log.Fatal(err)
	}
	err = cmdServer.Flags().MarkHidden("memprofile")
	if err != nil {
		log.Fatal(err)
	}

	if useHttp && serverPort == 50001 {
		serverPort = 5001
	}
	rootCmd.AddCommand(cmdServer, cmdClient, cmdCat, cmdCheck)
	cmdServer.AddCommand(cmdServerStart)
	cmdClient.AddCommand(cmdClientCreate, cmdClientRead, cmdClientWrite, cmdClientBench)
}
