package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/locona/proxyhealth/pkg/fileutil"
	"github.com/locona/proxyhealth/pkg/proxy"
	"github.com/spf13/cobra"
)

const successProxyFile = "success_proxy_list"
const defaultProxyFile = "proxy_list"

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use: "run",
	Run: func(cmd *cobra.Command, args []string) {
		proxyFilePath := defaultProxyFile
		if len(args) == 1 {
			proxyFilePath = defaultProxyFile
		}
		Run(proxyFilePath)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func Run(proxyFile string) {
	proxyFQDNList, err := proxyFQDNList(proxyFile)
	if err != nil {
		panic(err)
	}

	fp, _, err := fileutil.Recreate(successProxyFile)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	for i := range proxyFQDNList {
		err := proxy.Health(proxyFQDNList[i])
		if err == nil {
			fmt.Fprintln(fp, proxyFQDNList[i])
		}

		if err != nil {
			log.Printf("Failure Proxy: %v \n", proxyFQDNList[i])
		}
	}
}

func proxyFQDNList(proxyFile string) ([]string, error) {
	fp, err := os.Open(proxyFile)
	defer fp.Close()
	if err != nil {
		return []string{}, err
	}

	scanner := bufio.NewScanner(fp)
	res := make([]string, 0)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, nil
}
