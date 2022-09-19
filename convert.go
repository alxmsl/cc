package main

import (
	"fmt"
	"github.com/alxmsl/cc/service"
	"github.com/spf13/cobra"
	"log"
)

var (
	amount   float64
	from, to string

	convertCmd = &cobra.Command{
		Use: "convert",
		Short: `
Coins converter using CoinMarketCap API [https://coinmarketcap.com/api/documentation/v1/#section/Introduction]

Tool automatically finds credentials based on the application environment [CMC_API_KEY], and it uses those credentials 
to authenticate to CoinMarketCap API. Credentials could be obtained at [https://pro.coinmarketcap.com/signup/]
`,
		Run: func(cmd *cobra.Command, args []string) {
			var (
				svc      = service.New()
				res, err = svc.Convert(amount, from, to)
			)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(res)
		},
	}
)

func init() {
	convertCmd.PersistentFlags().Float64Var(&amount, "amount", .0, "converted amount")
	convertCmd.PersistentFlags().StringVar(&from, "from", "", "converted from")
	convertCmd.PersistentFlags().StringVar(&to, "to", "", "converted to")

	rootCmd.AddCommand(convertCmd)
}
