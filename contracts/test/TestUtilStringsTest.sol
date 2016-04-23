import {Test} from "sol-tester/Test.sol";

contract TestUtilStringsTest is Test {

    function testStrEqualSuccess() {
        assert(strEqual("abcd", "abcd"), "strEqual failed.");
    }

    function testStrEqualMultiWordSuccess() {
        assert(strEqual(
            "abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd",
            "abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd"
        ), "strEqual failed");
    }

    function testStrEqualFailLengthNotEqual() {
        assert(!strEqual("abcd", "abcdef"), "strEqual failed.");
    }

    function testStrEqualFail() {
        assert(!strEqual("abcd", "abfd"), "strEqual failed.");
    }

    function testStrEqualMultiWordFail() {
        assert(!strEqual(
            "abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd",
            "abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcyabcdabcdabcdabcdabcd"
        ), "strEqual failed");
    }

    function testStrEmptySuccess() {
        assert(strEmpty(""), "strEmpty failed.");
    }

    function testStrEmptyFail() {
        assert(!strEmpty("g"), "strEmpty failed.");
    }

}