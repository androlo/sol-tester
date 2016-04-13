import {CodeSigner} from "./CodeSigner.sol";
import {Test} from "./Test.sol";

contract CodeSignerTest is Test {

    function testSign() {
        CodeSigner cs = new CodeSigner();
        cs.getHash(this);

    }

}