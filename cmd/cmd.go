package cmd

import (
	"decipher.com/tps/benchmark"
	"decipher.com/tps/config"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "antps",
	Short: "EVM Blockchain Benchmark Application",
	Long:  "This is a command line application for benchmarking",
}

func Execute() {
	config.LoadAddresses("config/config.yml")
	if err := rootCmd.Execute(); err != nil {
		log.Printf("Execute err: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(updateConfig)
	rootCmd.AddCommand(erc20MintCmd)
	rootCmd.AddCommand(erc20TransferCmd)
	rootCmd.AddCommand(erc721MintCmd)
	rootCmd.AddCommand(erc721TransferCmd)
	rootCmd.AddCommand(erc1155MintCmd)
	rootCmd.AddCommand(erc1155TransferCmd)
	rootCmd.AddCommand(nativeTransferCmd)
	rootCmd.AddCommand(multiTransferCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize contracts",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark.InitAccount(1)
		ERC20, ERC721, ERC1155 := benchmark.InitContract()
		benchmark.UpdateAddress(ERC20, ERC721, ERC1155)
	},
}

var updateConfig = &cobra.Command{
	Use:   "updateConfig",
	Short: "Update Network Config",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark.UpdateConfig(args[0])
	},
}

var erc20MintCmd = &cobra.Command{
	Use:   "erc20mint",
	Short: "Mint ERC20 tokens",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark.InitAccount(config.Total)
		benchmark.ERC20Mint(config.Total, config.Rate, config.ERC20ADDRESS)
	},
}

var erc20TransferCmd = &cobra.Command{
	Use:   "erc20transfer",
	Short: "Transfer ERC20 tokens",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark.InitAccount(config.Total)
		benchmark.ERC20Transfer(config.Total, config.Rate, config.ERC20ADDRESS)
	},
}

var erc721MintCmd = &cobra.Command{
	Use:   "erc721mint",
	Short: "Mint ERC721 tokens",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark.InitAccount(config.Total)
		benchmark.ERC721Mint(config.Total, config.Rate, config.ERC721ADDRESS)
	},
}

var erc721TransferCmd = &cobra.Command{
	Use:   "erc721transfer",
	Short: "Transfer ERC721 tokens",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark.InitAccount(config.Total)
		benchmark.ERC721Transfer(config.Total, config.Rate, config.ERC721ADDRESS)
	},
}

var erc1155MintCmd = &cobra.Command{
	Use:   "erc1155mint",
	Short: "Mint ERC1155 tokens",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark.InitAccount(config.Total)
		benchmark.ERC1155Mint(config.Total, config.Rate, config.ERC1155ADDRESS)
	},
}

var erc1155TransferCmd = &cobra.Command{
	Use:   "erc1155transfer",
	Short: "Transfer ERC1155 tokens",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark.InitAccount(config.Total)
		benchmark.ERC1155Transfer(config.Total, config.Rate, config.ERC1155ADDRESS)
	},
}

var nativeTransferCmd = &cobra.Command{
	Use:   "nativetransfer",
	Short: "Transfer Native Coins",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark.InitAccount(config.Total)
		benchmark.NativeTransfer(config.Total, config.Rate)
	},
}

var multiTransferCmd = &cobra.Command{
	Use:   "multitransfer",
	Short: "Transfer Native Coins ",
	Run: func(cmd *cobra.Command, args []string) {
		benchmark.InitAccount(config.Total)
		benchmark.MultiTransfer(config.Total)
	},
}
