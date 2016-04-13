import {Test} from "./Test.sol";

contract TestUtilToASCIITest is Test {

    // utoa

    function testUtoaValueIs0() {
        assert(strEqual(utoa(0, 10), "0"), "utoa failed.");
    }

    function testUtoaRadixLessThen2() {
        assert(strEqual(utoa(2345, 1), "0"), "utoa failed.");
    }

    function testUtoaRadixLargerThen16() {
        assert(strEqual(utoa(2345, 17), "0"), "utoa failed.");
    }

    function testUtoaRadix10Success1() {
        assert(strEqual(utoa(5, 10), "5"), "utoa failed.");
    }

    function testUtoaRadix10Success2() {
        assert(strEqual(utoa(452345, 10), "452345"), "utoa failed.");
    }

    function testUtoaRadix16Success1() {
        assert(strEqual(utoa(5, 16), "5"), "utoa failed.");
    }

    function testUtoaRadix16Success2() {
        assert(strEqual(utoa(2345234, 16), "23c912"), "utoa failed.");
    }

    function testUtoaRadix2Success1() {
        assert(strEqual(utoa(345234534, 2), "10100100100111101110001100110"), "utoa failed.");
    }

    // itoa

    function testItoaValueIs0() {
        assert(strEqual(itoa(0, 10), "0"), "itoa failed.");
    }

    function testItoaValueIsMinus0() {
        assert(strEqual(itoa(-0, 10), "0"), "itoa failed.");
    }

    function testItoaRadixLessThen2() {
        assert(strEqual(itoa(2345, 1), "0"), "itoa failed.");
    }

    function testItoaRadixLargerThen16() {
        assert(strEqual(itoa(2345, 17), "0"), "itoa failed.");
    }

    function testItoaRadix10Success1() {
        assert(strEqual(itoa(5, 10), "5"), "itoa failed.");
    }

    function testItoaRadix10SuccessNeg1() {
        assert(strEqual(itoa(-5, 10), "-5"), "itoa failed.");
    }

    function testItoaRadix10Success2() {
        assert(strEqual(itoa(452345, 10), "452345"), "itoa failed.");
    }

    function testItoaRadix10SuccessNeg2() {
        assert(strEqual(itoa(-452345, 10), "-452345"), "itoa failed.");
    }

    function testItoaRadix16Success1() {
        assert(strEqual(itoa(5, 16), "5"), "itoa failed.");
    }

    function testItoaRadix16SuccessNeg1() {
        assert(strEqual(itoa(-5, 16), "-5"), "itoa failed.");
    }

    function testItoaRadix16Success2() {
        assert(strEqual(itoa(2345234, 16), "23c912"), "itoa failed.");
    }

    function testItoaRadix16SuccessNeg2() {
        assert(strEqual(itoa(-2345234, 16), "-23c912"), "itoa failed.");
    }

    function testItoaRadix2Success1() {
        assert(strEqual(itoa(345234534, 2), "10100100100111101110001100110"), "itoa failed.");
    }

    function testItoaRadix2SuccessNeg1() {
        assert(strEqual(itoa(-345234534, 2), "-10100100100111101110001100110"), "itoa failed.");
    }

    // ttoa

    function testTtoaTrue() {
        assert(strEqual(ttoa(true), "true"), "ttoa failed.");
    }

    function testTtoaFalse() {
        assert(strEqual(ttoa(false), "false"), "ttoa failed.");
    }

    // rtoa

    function testRtoaSuccess() {
        address addr = 0x1234567812345678123456781234567812345678;
        assert(strEqual(rtoa(addr), "1234567812345678123456781234567812345678"), "rtoa failed.");
    }

    function testRtoaSuccess2() {
        address addr = 0x692a70d2e424a56d2c6c27aa97d1a86395877b3a;
        assert(strEqual(rtoa(addr), "692a70d2e424a56d2c6c27aa97d1a86395877b3a"), "rtoa failed.");
    }

    function testRtoaSuccess3() {
        address addr = 0x51b3;
        assert(strEqual(rtoa(addr), "00000000000000000000000000000000000051b3"), "rtoa failed.");
    }

    function testRtoaSuccess4() {
        address addr = 0x1;
        assert(strEqual(rtoa(addr), "0000000000000000000000000000000000000001"), "rtoa failed.");
    }

    function testRtoaSuccess5() {
        address addr = 0x01;
        assert(strEqual(rtoa(addr), "0000000000000000000000000000000000000001"), "rtoa failed.");
    }

    function testRtoaSuccess6() {
        address addr = 0x10;
        assert(strEqual(rtoa(addr), "0000000000000000000000000000000000000010"), "rtoa failed.");
    }

    function testRtoaSuccess7() {
        address addr = 0;
        assert(strEqual(rtoa(addr), "0000000000000000000000000000000000000000"), "rtoa failed.");
    }

    // btoa

    function testBtoaSuccess() {
        bytes32 bts = 0x1234432112344321123443211234432112344321123443211234432112344321;
        assert(strEqual(btoa(bts), "1234432112344321123443211234432112344321123443211234432112344321"), "btoa failed.");
    }

    function testBtoaSuccess2() {
        bytes32 bts = 0x1;
        assert(strEqual(btoa(bts), "0000000000000000000000000000000000000000000000000000000000000001"), "btoa failed.");
    }

    function testBtoaSuccess3() {
        bytes32 bts = 0x01;
        assert(strEqual(btoa(bts), "0000000000000000000000000000000000000000000000000000000000000001"), "btoa failed.");
    }

    function testBtoaSuccess4() {
        bytes32 bts = 0x10;
        assert(strEqual(btoa(bts), "0000000000000000000000000000000000000000000000000000000000000010"), "btoa failed.");
    }

    function testBtoaSuccess5() {
        bytes32 bts = 0x53a59a172925d8204acdfd12fe2035655aec9772ce03aa7a1de38b14af553259;
        assert(strEqual(btoa(bts), "53a59a172925d8204acdfd12fe2035655aec9772ce03aa7a1de38b14af553259"), "btoa failed.");
    }

}