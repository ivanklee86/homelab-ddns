package cmd

import (
	"fmt"
	"github.com/rdegges/go-ipify"
	"github.com/spf13/cobra"
	"log"
	"net"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates domain records with your home's public IP.",
	Long:  `Pings third-party API for IP and updates configured DNS records.`,
	Run: func(cmd *cobra.Command, args []string) {
		rawIp, err := ipify.GetIp()
		if err != nil {
			er(err)
		}
		log.Printf("IP found!  Ipify public IP: %s", rawIp)

		ip := net.ParseIP(rawIp)
		log.Printf("Parsed IP: %s", ip.String())

		fmt.Fprintf(cmd.OutOrStdout(), "Your IP is %s.", ip.String())
	},
}
