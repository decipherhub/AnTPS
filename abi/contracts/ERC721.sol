// contracts/GameItem.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract ANTPSERC721 is ERC721URIStorage {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;
    uint256 id;

    constructor() ERC721("ANTPS NFT", "ANTPS") {
        id = 0;
    }

    function mint(address _to) external {
        _mint(_to, id);
        id++;
    }
}