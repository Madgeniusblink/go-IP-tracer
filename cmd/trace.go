/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP address",
	Long:  `Trace the IP address`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
				showData(ip)
			}
		} else {
			fmt.Println("Please provide IP to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func showData(ip string) {
	fmt.Println("Showing data")
	url := fmt.Sprintf("https://ipinfo.io/%s/geo", ip)
	responseByte := getData(url)
	data := Ip{}
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		fmt.Println("unable to unmarshal the response")
	}

	fmt.Println("DATA FOUND:")
	fmt.Println(data)
}

func getData(url string) []byte {
	response, err := http.Get(url)
	fmt.Println("Getting data")
	if err != nil {
		fmt.Println(err)
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return responseByte
}
