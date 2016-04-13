import {Test} from "./Test.sol";

contract TestUtilTagsTest is Test {

    function testTagString() {
        assert(strEqual(tagString("abcd", "tag"), "tag: abcd"), "tagString failed.");
    }

    function testTagInt() {
        assert(strEqual(tagInt(-317, "tag"), "tag: -317"), "tagInt failed.");
    }

    function testTagUint() {
        assert(strEqual(tagUint(56781, "tag"), "tag: 56781"), "tagUint failed.");
    }

    function testTagBool() {
        assert(strEqual(tagBool(true, "tag"), "tag: true"), "tagBool failed.");
    }

    function testTagAddress() {
        address addr = 0x692a70d2e424a56d2c6c27aa97d1a86395877b3a;
        assert(strEqual(tagAddress(addr, "tag"), "tag: 692a70d2e424a56d2c6c27aa97d1a86395877b3a"), "tagAddress failed.");
    }

    function testTagBytes32() {
        bytes32 b = 0xb1591967aed668a4b27645ff40c444892d91bf5951b382995d4d4f6ee3a2ce03;
        assert(strEqual(tagBytes32(b, "tag"), "tag: b1591967aed668a4b27645ff40c444892d91bf5951b382995d4d4f6ee3a2ce03"), "tagBytes32 failed.");
    }

}