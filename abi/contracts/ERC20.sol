// contracts/GLDToken.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract ANTPSERC20 is ERC20 {
    constructor(uint256 initialSupply) ERC20("ANTPS TOKEN", "ANTPS") {
        _mint(msg.sender, initialSupply);
    }

    function mint(address _to, uint256 _amount) external {
        _mint(_to, _amount);
    }
}