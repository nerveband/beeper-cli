package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/nerveband/beeper-cli/internal/api"
	"github.com/nerveband/beeper-cli/internal/output"
)

var sendCmd = &cobra.Command{
	Use:   "send --chat-id <chat-id> --message <text>",
	Short: "Send a message to a chat",
	Long:  `Send a new message to the specified Beeper chat.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		chatID, _ := cmd.Flags().GetString("chat-id")
		message, _ := cmd.Flags().GetString("message")

		if chatID == "" {
			return fmt.Errorf("--chat-id is required")
		}
		if message == "" {
			return fmt.Errorf("--message is required")
		}

		client := api.NewClient(cfg.APIURL)

		resp, err := client.SendMessage(chatID, message)
		if err != nil {
			return fmt.Errorf("failed to send message: %w", err)
		}

		formatted, err := output.FormatSendResponse(resp, getOutputFormat())
		if err != nil {
			return err
		}

		fmt.Print(formatted)
		return nil
	},
}

func init() {
	sendCmd.Flags().String("chat-id", "", "Chat ID to send message to")
	sendCmd.Flags().String("message", "", "Message text to send")
	rootCmd.AddCommand(sendCmd)
}
