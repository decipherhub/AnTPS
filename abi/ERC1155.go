// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ERC1155MetaData contains all meta data concerning the ERC1155 contract.
var ERC1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIELD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SILVER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SWORD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"THORS_HAMMER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040518060600160405280602781526020016200243c602791396200003c816200004b60201b60201c565b506001600381905550620003a8565b80600290816200005c9190620002c4565b5050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620000dc57607f821691505b602082108103620000f257620000f162000097565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620001567fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000119565b62000162868362000119565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620001ac620001a6620001a0846200017a565b62000183565b6200017a565b9050919050565b5f819050919050565b620001c7836200018c565b620001df620001d682620001b3565b84845462000125565b825550505050565b5f90565b620001f5620001e7565b62000202818484620001bc565b505050565b5b8181101562000229576200021d5f82620001eb565b60018101905062000208565b5050565b601f82111562000278576200024281620000f8565b6200024d846200010a565b810160208510156200025d578190505b620002756200026c856200010a565b83018262000207565b50505b505050565b5f82821c905092915050565b5f6200029a5f19846008026200027d565b1980831691505092915050565b5f620002b4838362000289565b9150826002028217905092915050565b620002cf8262000060565b67ffffffffffffffff811115620002eb57620002ea6200006a565b5b620002f78254620000c4565b620003048282856200022d565b5f60209050601f8311600181146200033a575f841562000325578287015190505b620003318582620002a7565b865550620003a0565b601f1984166200034a86620000f8565b5f5b8281101562000373578489015182556001820191506020850194506020810190506200034c565b868310156200039357848901516200038f601f89168262000289565b8355505b6001600288020188555050505b505050505050565b61208680620003b65f395ff3fe608060405234801561000f575f80fd5b50600436106100e7575f3560e01c80634e1273f41161008a578063d562e20411610064578063d562e20414610259578063e3e55f0814610277578063e985e9c514610295578063f242432a146102c5576100e7565b80634e1273f4146101ef5780635b2725ed1461021f578063a22cb4651461023d576100e7565b806313dc989f116100c657806313dc989f1461017b5780632eb2c2d6146101995780633e4bee38146101b557806340c10f19146101d3576100e7565b8062fdd58e146100eb57806301ffc9a71461011b5780630e89341c1461014b575b5f80fd5b610105600480360381019061010091906114f3565b6102e1565b6040516101129190611540565b60405180910390f35b610135600480360381019061013091906115ae565b610336565b60405161014291906115f3565b60405180910390f35b6101656004803603810190610160919061160c565b610417565b60405161017291906116c1565b60405180910390f35b6101836104a9565b6040516101909190611540565b60405180910390f35b6101b360048036038101906101ae91906118d1565b6104ae565b005b6101bd610555565b6040516101ca9190611540565b60405180910390f35b6101ed60048036038101906101e891906114f3565b610559565b005b61020960048036038101906102049190611a5c565b610590565b6040516102169190611b89565b60405180910390f35b61022761069d565b6040516102349190611540565b60405180910390f35b61025760048036038101906102529190611bd3565b6106a2565b005b6102616106b8565b60405161026e9190611540565b60405180910390f35b61027f6106bd565b60405161028c9190611540565b60405180910390f35b6102af60048036038101906102aa9190611c11565b6106c2565b6040516102bc91906115f3565b60405180910390f35b6102df60048036038101906102da9190611c4f565b610750565b005b5f805f8381526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f7fd9b67a26000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061040057507f0e89341c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610410575061040f826107f7565b5b9050919050565b60606002805461042690611d0f565b80601f016020809104026020016040519081016040528092919081815260200182805461045290611d0f565b801561049d5780601f106104745761010080835404028352916020019161049d565b820191905f5260205f20905b81548152906001019060200180831161048057829003601f168201915b50505050509050919050565b600381565b5f6104b7610860565b90508073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141580156104fc57506104fa86826106c2565b155b156105405780866040517fe237d922000000000000000000000000000000000000000000000000000000008152600401610537929190611d4e565b60405180910390fd5b61054d8686868686610867565b505050505050565b5f81565b610575826003548360405180602001604052805f81525061095b565b60035f81548092919061058790611da2565b91905055505050565b606081518351146105dc57815183516040517f5b0599910000000000000000000000000000000000000000000000000000000081526004016105d3929190611de9565b60405180910390fd5b5f835167ffffffffffffffff8111156105f8576105f76116e5565b5b6040519080825280602002602001820160405280156106265781602001602082028036833780820191505090505b5090505f5b84518110156106925761066261064a82876109f090919063ffffffff16565b61065d8387610a0390919063ffffffff16565b6102e1565b82828151811061067557610674611e10565b5b6020026020010181815250508061068b90611da2565b905061062b565b508091505092915050565b600481565b6106b46106ad610860565b8383610a16565b5050565b600281565b600181565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f610759610860565b90508073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415801561079e575061079c86826106c2565b155b156107e25780866040517fe237d9220000000000000000000000000000000000000000000000000000000081526004016107d9929190611d4e565b60405180910390fd5b6107ef8686868686610b7f565b505050505050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036108d7575f6040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016108ce9190611e3d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603610947575f6040517f01a8351400000000000000000000000000000000000000000000000000000000815260040161093e9190611e3d565b60405180910390fd5b6109548585858585610c85565b5050505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036109cb575f6040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016109c29190611e3d565b60405180910390fd5b5f806109d78585610d31565b915091506109e85f87848487610c85565b505050505050565b5f60208202602084010151905092915050565b5f60208202602084010151905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a86575f6040517fced3e100000000000000000000000000000000000000000000000000000000008152600401610a7d9190611e3d565b60405180910390fd5b8060015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610b7291906115f3565b60405180910390a3505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610bef575f6040517f57f447ce000000000000000000000000000000000000000000000000000000008152600401610be69190611e3d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1603610c5f575f6040517f01a83514000000000000000000000000000000000000000000000000000000008152600401610c569190611e3d565b60405180910390fd5b5f80610c6b8585610d31565b91509150610c7c8787848487610c85565b50505050505050565b610c9185858585610d61565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614610d2a575f610ccd610860565b90506001845103610d19575f610cec5f86610a0390919063ffffffff16565b90505f610d025f86610a0390919063ffffffff16565b9050610d128389898585896110f7565b5050610d28565b610d278187878787876112a6565b5b505b5050505050565b60608060405191506001825283602083015260408201905060018152826020820152604081016040529250929050565b8051825114610dab57815181516040517f5b059991000000000000000000000000000000000000000000000000000000008152600401610da2929190611de9565b60405180910390fd5b5f610db4610860565b90505f5b8351811015610fb6575f610dd58286610a0390919063ffffffff16565b90505f610deb8386610a0390919063ffffffff16565b90505f73ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614610f0e575f805f8481526020019081526020015f205f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610eba57888183856040517f03dee4c5000000000000000000000000000000000000000000000000000000008152600401610eb19493929190611e56565b60405180910390fd5b8181035f808581526020019081526020015f205f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610fa357805f808481526020019081526020015f205f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f9b9190611e99565b925050819055505b505080610faf90611da2565b9050610db8565b506001835103611071575f610fd45f85610a0390919063ffffffff16565b90505f610fea5f85610a0390919063ffffffff16565b90508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f628585604051611062929190611de9565b60405180910390a450506110f0565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb86866040516110e7929190611ecc565b60405180910390a45b5050505050565b5f8473ffffffffffffffffffffffffffffffffffffffff163b111561129e578373ffffffffffffffffffffffffffffffffffffffff1663f23a6e6187878686866040518663ffffffff1660e01b8152600401611157959493929190611f53565b6020604051808303815f875af192505050801561119257506040513d601f19601f8201168201806040525081019061118f9190611fbf565b60015b611213573d805f81146111c0576040519150601f19603f3d011682016040523d82523d5f602084013e6111c5565b606091505b505f81510361120b57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016112029190611e3d565b60405180910390fd5b805181602001fd5b63f23a6e6160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461129c57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016112939190611e3d565b60405180910390fd5b505b505050505050565b5f8473ffffffffffffffffffffffffffffffffffffffff163b111561144d578373ffffffffffffffffffffffffffffffffffffffff1663bc197c8187878686866040518663ffffffff1660e01b8152600401611306959493929190611fea565b6020604051808303815f875af192505050801561134157506040513d601f19601f8201168201806040525081019061133e9190611fbf565b60015b6113c2573d805f811461136f576040519150601f19603f3d011682016040523d82523d5f602084013e611374565b606091505b505f8151036113ba57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016113b19190611e3d565b60405180910390fd5b805181602001fd5b63bc197c8160e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461144b57846040517f57f447ce0000000000000000000000000000000000000000000000000000000081526004016114429190611e3d565b60405180910390fd5b505b505050505050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61148f82611466565b9050919050565b61149f81611485565b81146114a9575f80fd5b50565b5f813590506114ba81611496565b92915050565b5f819050919050565b6114d2816114c0565b81146114dc575f80fd5b50565b5f813590506114ed816114c9565b92915050565b5f80604083850312156115095761150861145e565b5b5f611516858286016114ac565b9250506020611527858286016114df565b9150509250929050565b61153a816114c0565b82525050565b5f6020820190506115535f830184611531565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61158d81611559565b8114611597575f80fd5b50565b5f813590506115a881611584565b92915050565b5f602082840312156115c3576115c261145e565b5b5f6115d08482850161159a565b91505092915050565b5f8115159050919050565b6115ed816115d9565b82525050565b5f6020820190506116065f8301846115e4565b92915050565b5f602082840312156116215761162061145e565b5b5f61162e848285016114df565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561166e578082015181840152602081019050611653565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61169382611637565b61169d8185611641565b93506116ad818560208601611651565b6116b681611679565b840191505092915050565b5f6020820190508181035f8301526116d98184611689565b905092915050565b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61171b82611679565b810181811067ffffffffffffffff8211171561173a576117396116e5565b5b80604052505050565b5f61174c611455565b90506117588282611712565b919050565b5f67ffffffffffffffff821115611777576117766116e5565b5b602082029050602081019050919050565b5f80fd5b5f61179e6117998461175d565b611743565b905080838252602082019050602084028301858111156117c1576117c0611788565b5b835b818110156117ea57806117d688826114df565b8452602084019350506020810190506117c3565b5050509392505050565b5f82601f830112611808576118076116e1565b5b813561181884826020860161178c565b91505092915050565b5f80fd5b5f67ffffffffffffffff82111561183f5761183e6116e5565b5b61184882611679565b9050602081019050919050565b828183375f83830152505050565b5f61187561187084611825565b611743565b90508281526020810184848401111561189157611890611821565b5b61189c848285611855565b509392505050565b5f82601f8301126118b8576118b76116e1565b5b81356118c8848260208601611863565b91505092915050565b5f805f805f60a086880312156118ea576118e961145e565b5b5f6118f7888289016114ac565b9550506020611908888289016114ac565b945050604086013567ffffffffffffffff81111561192957611928611462565b5b611935888289016117f4565b935050606086013567ffffffffffffffff81111561195657611955611462565b5b611962888289016117f4565b925050608086013567ffffffffffffffff81111561198357611982611462565b5b61198f888289016118a4565b9150509295509295909350565b5f67ffffffffffffffff8211156119b6576119b56116e5565b5b602082029050602081019050919050565b5f6119d96119d48461199c565b611743565b905080838252602082019050602084028301858111156119fc576119fb611788565b5b835b81811015611a255780611a1188826114ac565b8452602084019350506020810190506119fe565b5050509392505050565b5f82601f830112611a4357611a426116e1565b5b8135611a538482602086016119c7565b91505092915050565b5f8060408385031215611a7257611a7161145e565b5b5f83013567ffffffffffffffff811115611a8f57611a8e611462565b5b611a9b85828601611a2f565b925050602083013567ffffffffffffffff811115611abc57611abb611462565b5b611ac8858286016117f4565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611b04816114c0565b82525050565b5f611b158383611afb565b60208301905092915050565b5f602082019050919050565b5f611b3782611ad2565b611b418185611adc565b9350611b4c83611aec565b805f5b83811015611b7c578151611b638882611b0a565b9750611b6e83611b21565b925050600181019050611b4f565b5085935050505092915050565b5f6020820190508181035f830152611ba18184611b2d565b905092915050565b611bb2816115d9565b8114611bbc575f80fd5b50565b5f81359050611bcd81611ba9565b92915050565b5f8060408385031215611be957611be861145e565b5b5f611bf6858286016114ac565b9250506020611c0785828601611bbf565b9150509250929050565b5f8060408385031215611c2757611c2661145e565b5b5f611c34858286016114ac565b9250506020611c45858286016114ac565b9150509250929050565b5f805f805f60a08688031215611c6857611c6761145e565b5b5f611c75888289016114ac565b9550506020611c86888289016114ac565b9450506040611c97888289016114df565b9350506060611ca8888289016114df565b925050608086013567ffffffffffffffff811115611cc957611cc8611462565b5b611cd5888289016118a4565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611d2657607f821691505b602082108103611d3957611d38611ce2565b5b50919050565b611d4881611485565b82525050565b5f604082019050611d615f830185611d3f565b611d6e6020830184611d3f565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611dac826114c0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611dde57611ddd611d75565b5b600182019050919050565b5f604082019050611dfc5f830185611531565b611e096020830184611531565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f602082019050611e505f830184611d3f565b92915050565b5f608082019050611e695f830187611d3f565b611e766020830186611531565b611e836040830185611531565b611e906060830184611531565b95945050505050565b5f611ea3826114c0565b9150611eae836114c0565b9250828201905080821115611ec657611ec5611d75565b5b92915050565b5f6040820190508181035f830152611ee48185611b2d565b90508181036020830152611ef88184611b2d565b90509392505050565b5f81519050919050565b5f82825260208201905092915050565b5f611f2582611f01565b611f2f8185611f0b565b9350611f3f818560208601611651565b611f4881611679565b840191505092915050565b5f60a082019050611f665f830188611d3f565b611f736020830187611d3f565b611f806040830186611531565b611f8d6060830185611531565b8181036080830152611f9f8184611f1b565b90509695505050505050565b5f81519050611fb981611584565b92915050565b5f60208284031215611fd457611fd361145e565b5b5f611fe184828501611fab565b91505092915050565b5f60a082019050611ffd5f830188611d3f565b61200a6020830187611d3f565b818103604083015261201c8186611b2d565b905081810360608301526120308185611b2d565b905081810360808301526120448184611f1b565b9050969550505050505056fea2646970667358221220fc53a6878edb1df1b8611a1f7a4ac562c6030dd70e913c73038ccf8b75fbaf6164736f6c6343000814003368747470733a2f2f67616d652e6578616d706c652f6170692f6974656d2f7b69647d2e6a736f6e",
}

// ERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155MetaData.ABI instead.
var ERC1155ABI = ERC1155MetaData.ABI

// ERC1155Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC1155MetaData.Bin instead.
var ERC1155Bin = ERC1155MetaData.Bin

// DeployERC1155 deploys a new Ethereum contract, binding an instance of ERC1155 to it.
func DeployERC1155(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC1155, error) {
	parsed, err := ERC1155MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC1155Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC1155{ERC1155Caller: ERC1155Caller{contract: contract}, ERC1155Transactor: ERC1155Transactor{contract: contract}, ERC1155Filterer: ERC1155Filterer{contract: contract}}, nil
}

// ERC1155 is an auto generated Go binding around an Ethereum contract.
type ERC1155 struct {
	ERC1155Caller     // Read-only binding to the contract
	ERC1155Transactor // Write-only binding to the contract
	ERC1155Filterer   // Log filterer for contract events
}

// ERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155Session struct {
	Contract     *ERC1155          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155CallerSession struct {
	Contract *ERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155TransactorSession struct {
	Contract     *ERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155Raw struct {
	Contract *ERC1155 // Generic contract binding to access the raw methods on
}

// ERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155CallerRaw struct {
	Contract *ERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// ERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155TransactorRaw struct {
	Contract *ERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155 creates a new instance of ERC1155, bound to a specific deployed contract.
func NewERC1155(address common.Address, backend bind.ContractBackend) (*ERC1155, error) {
	contract, err := bindERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155{ERC1155Caller: ERC1155Caller{contract: contract}, ERC1155Transactor: ERC1155Transactor{contract: contract}, ERC1155Filterer: ERC1155Filterer{contract: contract}}, nil
}

// NewERC1155Caller creates a new read-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Caller(address common.Address, caller bind.ContractCaller) (*ERC1155Caller, error) {
	contract, err := bindERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Caller{contract: contract}, nil
}

// NewERC1155Transactor creates a new write-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155Transactor, error) {
	contract, err := bindERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Transactor{contract: contract}, nil
}

// NewERC1155Filterer creates a new log filterer instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155Filterer, error) {
	contract, err := bindERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155Filterer{contract: contract}, nil
}

// bindERC1155 binds a generic wrapper to an already deployed contract.
func bindERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.ERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transact(opts, method, params...)
}

// GOLD is a free data retrieval call binding the contract method 0x3e4bee38.
//
// Solidity: function GOLD() view returns(uint256)
func (_ERC1155 *ERC1155Caller) GOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "GOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GOLD is a free data retrieval call binding the contract method 0x3e4bee38.
//
// Solidity: function GOLD() view returns(uint256)
func (_ERC1155 *ERC1155Session) GOLD() (*big.Int, error) {
	return _ERC1155.Contract.GOLD(&_ERC1155.CallOpts)
}

// GOLD is a free data retrieval call binding the contract method 0x3e4bee38.
//
// Solidity: function GOLD() view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) GOLD() (*big.Int, error) {
	return _ERC1155.Contract.GOLD(&_ERC1155.CallOpts)
}

// SHIELD is a free data retrieval call binding the contract method 0x5b2725ed.
//
// Solidity: function SHIELD() view returns(uint256)
func (_ERC1155 *ERC1155Caller) SHIELD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "SHIELD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIELD is a free data retrieval call binding the contract method 0x5b2725ed.
//
// Solidity: function SHIELD() view returns(uint256)
func (_ERC1155 *ERC1155Session) SHIELD() (*big.Int, error) {
	return _ERC1155.Contract.SHIELD(&_ERC1155.CallOpts)
}

// SHIELD is a free data retrieval call binding the contract method 0x5b2725ed.
//
// Solidity: function SHIELD() view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) SHIELD() (*big.Int, error) {
	return _ERC1155.Contract.SHIELD(&_ERC1155.CallOpts)
}

// SILVER is a free data retrieval call binding the contract method 0xe3e55f08.
//
// Solidity: function SILVER() view returns(uint256)
func (_ERC1155 *ERC1155Caller) SILVER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "SILVER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SILVER is a free data retrieval call binding the contract method 0xe3e55f08.
//
// Solidity: function SILVER() view returns(uint256)
func (_ERC1155 *ERC1155Session) SILVER() (*big.Int, error) {
	return _ERC1155.Contract.SILVER(&_ERC1155.CallOpts)
}

// SILVER is a free data retrieval call binding the contract method 0xe3e55f08.
//
// Solidity: function SILVER() view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) SILVER() (*big.Int, error) {
	return _ERC1155.Contract.SILVER(&_ERC1155.CallOpts)
}

// SWORD is a free data retrieval call binding the contract method 0x13dc989f.
//
// Solidity: function SWORD() view returns(uint256)
func (_ERC1155 *ERC1155Caller) SWORD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "SWORD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SWORD is a free data retrieval call binding the contract method 0x13dc989f.
//
// Solidity: function SWORD() view returns(uint256)
func (_ERC1155 *ERC1155Session) SWORD() (*big.Int, error) {
	return _ERC1155.Contract.SWORD(&_ERC1155.CallOpts)
}

// SWORD is a free data retrieval call binding the contract method 0x13dc989f.
//
// Solidity: function SWORD() view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) SWORD() (*big.Int, error) {
	return _ERC1155.Contract.SWORD(&_ERC1155.CallOpts)
}

// THORSHAMMER is a free data retrieval call binding the contract method 0xd562e204.
//
// Solidity: function THORS_HAMMER() view returns(uint256)
func (_ERC1155 *ERC1155Caller) THORSHAMMER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "THORS_HAMMER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// THORSHAMMER is a free data retrieval call binding the contract method 0xd562e204.
//
// Solidity: function THORS_HAMMER() view returns(uint256)
func (_ERC1155 *ERC1155Session) THORSHAMMER() (*big.Int, error) {
	return _ERC1155.Contract.THORSHAMMER(&_ERC1155.CallOpts)
}

// THORSHAMMER is a free data retrieval call binding the contract method 0xd562e204.
//
// Solidity: function THORS_HAMMER() view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) THORSHAMMER() (*big.Int, error) {
	return _ERC1155.Contract.THORSHAMMER(&_ERC1155.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155.Contract.BalanceOfBatch(&_ERC1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155.Contract.BalanceOfBatch(&_ERC1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155.Contract.IsApprovedForAll(&_ERC1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155.Contract.IsApprovedForAll(&_ERC1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155.Contract.SupportsInterface(&_ERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155.Contract.SupportsInterface(&_ERC1155.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155 *ERC1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155 *ERC1155Session) Uri(arg0 *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155 *ERC1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_ERC1155 *ERC1155Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_ERC1155 *ERC1155Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.Mint(&_ERC1155.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_ERC1155 *ERC1155TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.Mint(&_ERC1155.TransactOpts, _to, _amount)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ERC1155 *ERC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransferFrom(&_ERC1155.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransferFrom(&_ERC1155.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ERC1155 *ERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalForAll(&_ERC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalForAll(&_ERC1155.TransactOpts, operator, approved)
}

// ERC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC1155 contract.
type ERC1155ApprovalForAllIterator struct {
	Event *ERC1155ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155ApprovalForAll represents a ApprovalForAll event raised by the ERC1155 contract.
type ERC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ERC1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155ApprovalForAllIterator{contract: _ERC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155ApprovalForAll)
				if err := _ERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) ParseApprovalForAll(log types.Log) (*ERC1155ApprovalForAll, error) {
	event := new(ERC1155ApprovalForAll)
	if err := _ERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ERC1155 contract.
type ERC1155TransferBatchIterator struct {
	Event *ERC1155TransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155TransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TransferBatch represents a TransferBatch event raised by the ERC1155 contract.
type ERC1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferBatchIterator{contract: _ERC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ERC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TransferBatch)
				if err := _ERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) ParseTransferBatch(log types.Log) (*ERC1155TransferBatch, error) {
	event := new(ERC1155TransferBatch)
	if err := _ERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ERC1155 contract.
type ERC1155TransferSingleIterator struct {
	Event *ERC1155TransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155TransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TransferSingle represents a TransferSingle event raised by the ERC1155 contract.
type ERC1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferSingleIterator{contract: _ERC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ERC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TransferSingle)
				if err := _ERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) ParseTransferSingle(log types.Log) (*ERC1155TransferSingle, error) {
	event := new(ERC1155TransferSingle)
	if err := _ERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ERC1155 contract.
type ERC1155URIIterator struct {
	Event *ERC1155URI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155URI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1155URI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155URI represents a URI event raised by the ERC1155 contract.
type ERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ERC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155URIIterator{contract: _ERC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ERC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155URI)
				if err := _ERC1155.contract.UnpackLog(event, "URI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) ParseURI(log types.Log) (*ERC1155URI, error) {
	event := new(ERC1155URI)
	if err := _ERC1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
