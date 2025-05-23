/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

var (
	//TelegramBot Token
	TeleToken = os.Getenv("TELE_TOKEN")
)

// tBotCmd represents the tBot command
var tBotCmd = &cobra.Command{
	Use:     "tBot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("tBot %s started", appVersion)

		tBot, err := tele.NewBot(tele.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check Telegram Token env variable. %s", err)
			return
		}

		tBot.Handle(
			telebot.OnText, func(m telebot.Context) error {

				log.Print(m.Message().Payload, m.Text())
				payload := m.Message().Payload

				switch payload {
				case "hello":
					err = m.Send(fmt.Sprintf("Hello I'm tBot %s!", appVersion))
				}

				return err
			})

		tBot.Start()
	},
}

func init() {
	rootCmd.AddCommand(tBotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tBotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tBotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
