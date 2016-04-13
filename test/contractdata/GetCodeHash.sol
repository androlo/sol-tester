library GetCodeHash {
    function at(address a) constant returns (bytes32 codeHash) {
        assembly {
            let code := mload(0x40)
            let size := extcodesize(a)
            jumpi(end, eq(size, 0))
            extcodecopy(a, code, 0, size)
            mstore(0x40, add(code, size))
            codeHash := sha3(code, size)
            end:
        }
    }
}