contract DebugASM {

    // 0x3576493CC0A51DE8F7D9532F252B361B03C593648AABA260682ACCD2B51E6897
    event TestEvent(bool indexed result, bytes32 indexed expected, bytes32 indexed value);

    // Stack: [ ... , jumpdest, value, expected ]
    function _assertStack() internal {
        assembly {
                dup2
                dup2
                eq
                0x3576493CC0A51DE8F7D9532F252B361B03C593648AABA260682ACCD2B51E6897
                0x0
                0x0
                log4
                jump
        }
    }

}