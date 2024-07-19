// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract ANTPSERC1155 is ERC1155 {
    uint256 public constant GOLD = 0;
    uint256 public constant SILVER = 1;
    uint256 public constant THORS_HAMMER = 2;
    uint256 public constant SWORD = 3;
    uint256 public constant SHIELD = 4;
    uint256 id;

    constructor() ERC1155("https://game.example/api/item/{id}.json") {
        id = 1;
    }

    function mint(address _to, uint256 _amount) external {
        _mint(_to, id, _amount, "");
        id++;
    }
}