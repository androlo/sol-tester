contract TestUtils {

    // The null address: 0
    address constant ADDRESS_NULL = 0;
    // The null bytes32: 0
    bytes32 constant BYTES32_NULL = 0;
    // The null string: ""
    string constant STRING_NULL = "";

    // ASCII codes as uint8
    uint8 constant ZERO = uint8(byte('0'));
    uint8 constant CH_a = uint8(byte('a'));

    // ASCII codes as byte
    byte constant MINUS = byte('-');
    byte constant SPACE = byte(' ');

    // Strings

    // Strings equal
    function strEqual(string memory a, string memory b) constant internal returns (bool res) {
        assembly {
                let lA := mload(a)
                let lB := mload(b)
                jumpi(tag_finalize, not(eq(lA, lB)))
            tag_compare:
                {
                        let i := 0
                        let words := add(div(lA, 32), gt(mod(lA, 32), 0))
                        let offsetA := add(a, 32)
                        let offsetB := add(b, 32)
                    tag_loop:
                        {
                            let offset := mul(i, 32)
                            i := add(i, 1)
                            res := eq(mload(add(offsetA, offset)), mload(add(offsetB, offset)))
                        }
                        jumpi(tag_loop, and(lt(i, words), res) )
                }
            tag_finalize:
        }
    }

    // true if the string is empty, otherwise false
    function strEmpty(string memory str) constant internal returns (bool result) {
        assembly {
            result := not(mload(str))
        }
    }

    // Integer to ASCII string
    function itoa(int n, uint8 radix) internal constant returns (string memory str) {
        if (n == 0 || radix < 2 || radix > 16)
            return "0";

        bytes memory bts = new bytes(256);
        uint i;
        bool neg = false;
        if (n < 0) {
            n = -n;
            neg = true;
        }
        var m = uint(n);
        while (m > 0) {
            bts[i++] = _utoa(m % radix); // Turn it to ascii.
            m /= radix;
        }
        // Reverse
        uint size = i;
        uint j = 0;
        bytes memory rev;
        if (neg) {
            size++;
            j = 1;
            rev = new bytes(size);
            rev[0] = MINUS;
        }
        else
            rev = new bytes(size);

        for (; j < size; j++)
            rev[j] = bts[size - j - 1];
        str = string(rev);
    }

    // Unsigned integer to ASCII string
    function utoa(uint n, uint8 radix) internal constant returns (string memory str) {
        if (n == 0 || radix < 2 || radix > 16) {
            return "0";
        }
        bytes memory bts = new bytes(256);
        uint i;
        while (n > 0) {
            bts[i++] = _utoa(n % radix); // Turn it to ascii.
            n /= radix;
        }
        // Reverse
        bytes memory rev = new bytes(i);
        for (uint j = 0; j < i; j++)
            rev[j] = bts[i - j - 1];
        str = string(rev);
    }

    function _utoa(uint u) internal constant returns (byte) {
        if (u < 10)
            return byte(u + ZERO);
        else if (u < 16)
            return byte(u - 10 + CH_a);
        else
            return 0;
    }

    // Boolean
    function ttoa(bool val) internal constant returns (string) {
        if (val)
            return "true";
        else
            return "false";
    }

    // Address
    function rtoa(address addr) internal constant returns (string memory str) {
        bytes memory bts = new bytes(40);
        bytes20 b = bytes20(addr);
        for (uint i = 0; i < 20; i++) {
            uint bu = uint(b[i]);
            bts[2*i + 1] = _utoa(bu % 16);
            bts[2*i] = _utoa((bu / 16) % 16);
        }
        str = string(bts);
    }

    // Bytes32
    function btoa(bytes32 b) internal constant returns (string memory str) {
        bytes memory bts = new bytes(64);
        for (uint i = 0; i < 32; i++) {
            uint bu = uint(b[i]);
            bts[2*i + 1] = _utoa(bu % 16);
            bts[2*i] = _utoa((bu / 16) % 16);
        }
        str = string(bts);
    }

    function tagString(string value, string tag) internal constant returns (string) {

        bytes memory valueB = bytes(value);
        bytes memory tagB = bytes(tag);

        uint vl = valueB.length;
        uint tl = tagB.length;

        bytes memory newB = new bytes(vl + tl + 2);

        uint i;
        uint j;

        for (i = 0; i < tl; i++)
            newB[j++] = tagB[i];
        newB[j++] = ':';
        newB[j++] = ' ';
        for (i = 0; i < vl; i++)
            newB[j++] = valueB[i];

        return string(newB);
    }

    function tagInt(int value, string tag) internal constant returns (string) {
        var nstr = itoa(value, 10);
        return tagString(nstr, tag);
    }

    function tagUint(uint value, string tag) internal constant returns (string) {
        var nstr = utoa(value, 10);
        return tagString(nstr, tag);
    }

    function tagBool(bool value, string tag) internal constant returns (string) {
        var nstr = ttoa(value);
        return tagString(nstr, tag);
    }

    function tagAddress(address value, string tag) internal constant returns (string) {
        var nstr = rtoa(value);
        return tagString(nstr, tag);
    }

    function tagBytes32(bytes32 value, string tag) internal constant returns (string) {
        var nstr = btoa(value);
        return tagString(nstr, tag);
    }

    function appendTagged(string tagged, string str) internal constant returns (string) {

        bytes memory taggedB = bytes(tagged);
        bytes memory strB = bytes(str);

        uint sl = strB.length;
        uint tl = taggedB.length;

        bytes memory newB = new bytes(sl + tl + 3);

        uint i;
        uint j;

        for (i = 0; i < sl; i++)
            newB[j++] = strB[i];
        newB[j++] = ' ';
        newB[j++] = '(';
        for (i = 0; i < tl; i++)
            newB[j++] = taggedB[i];
        newB[j++] = ')';

        return string(newB);
    }

    function appendTagged(string tagged0, string tagged1, string str) internal constant returns (string) {

        bytes memory tagged0B = bytes(tagged0);
        bytes memory tagged1B = bytes(tagged1);
        bytes memory strB = bytes(str);

        uint sl = strB.length;
        uint t0l = tagged0B.length;
        uint t1l = tagged1B.length;

        bytes memory newB = new bytes(sl + t0l + t1l + 5);

        uint i;
        uint j;

        for (i = 0; i < sl; i++)
            newB[j++] = strB[i];
        newB[j++] = ' ';
        newB[j++] = '(';
        for (i = 0; i < t0l; i++)
            newB[j++] = tagged0B[i];
        newB[j++] = ',';
        newB[j++] = ' ';
        for (i = 0; i < t1l; i++)
            newB[j++] = tagged1B[i];
        newB[j++] = ')';

        return string(newB);
    }

}

contract Test is TestUtils {

    event TestEvent(bool indexed result, string message);

    function assert(bool value, string message) internal {
        if (value)
            TestEvent(true, "");
        else
            TestEvent(false, message);
    }

}