import {GetCodeHash} from "./GetCodeHash.sol";

contract CodeSigner {

    function getHash(address addr) constant returns (bytes32) {
        return GetCodeHash.at(addr);
    }

}