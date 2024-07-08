package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random passwords",
	Long:  "Generate random long passwords with customizable options",
	Run:   generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length of the generated password")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in the generated password")
	generateCmd.Flags().BoolP("special-chars", "s", false, "Include special characters in the generated password")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialCharacter, _ := cmd.Flags().GetBool("special-chars")

	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if isDigits {
		charSet += "0123456789"
	}

	if isSpecialCharacter {
		charSet += "!#$%&'()*+,-./:;<=>?@[]^_{|}~"
	}

	rand.Seed(time.Now().UnixNano()) // Initialize the random seed

	password := make([]byte, length)

	for i := range password {
		password[i] = charSet[rand.Intn(len(charSet))]
	}

	fmt.Println("Generated Password:")
	fmt.Println(string(password))
}
