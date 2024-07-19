# AnTPS

## About
AnTPS (Anti-TPS) is a tool designed to benchmark blockchain performance, specifically focusing on Transactions Per Second (TPS). In an era where TPS claims often become marketing tools rather than reliable metrics, AnTPS aims to provide an objective, transparent, and reproducible methodology for measuring blockchain performance.

## Key Features
- 🔍 **Objective Measurement**: Clearly defined criteria for TPS calculation
- 🌐 **Fully Support**: Support from infrastructure to measurement
- 🛠 **Multiple Scenarios**: Test different transaction types and loads
- 🔒 **Transparent Methodology**: Fully disclosed measurement process

## Team
AnTPS is developed by members of Decipher, the SNU Blockchain Research Center:<br>
[@rrhlrmrr](https://github.com/rrhlrmrr) [@bicoCrypto](https://github.com/bicoCrypto) [@SOLMIN LEE](https://github.com/solmingming)

## License
This project is licensed under the MIT License

## Getting Started

### Prerequisites
Ensure you have the following installed:
- Go (version 1.21.7 or higher)
- Python (version 3.8.10 or higher)
- Terraform (version 1.4.6 or higher)
- kubectl (version 1.28.2 or higher)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/decipherhub/AnTPS 
   cd AnTPS
   ```
2. Build the Project:
   ```bash
   make build-mac   # For macOS
   make build-linux # For Linux
   ```

### Usage
1. Start the blockchain nodes:
   ```bash
   make avalanche  # For Avalanche network
   make ethereum   # For Ethereum network
   make klaytn     # For Klaytn network
   ```

2. Deploy smart contracts:
   ```bash
   ./antps init
   ```

3. Run benchmarks:
   ```bash
   ./antps erc20mint      # Mint ERC20 tokens
   ./antps erc20transfer  # Transfer ERC20 tokens
   ./antps erc721mint     # Mint ERC721 tokens
   ./antps erc721transfer # Transfer ERC721 tokens
   ./antps erc1155mint    # Mint ERC1155 tokens
   ./antps erc1155transfer # Transfer ERC1155 tokens
   ./antps nativetransfer # Transfer native tokens (ETH, AVAX)
   ./antps multitransfer  # Transfer tokens from multiple accounts 
   ```

4. View results:
   ```bash
   make ava-output
   make eth-output
   ```
