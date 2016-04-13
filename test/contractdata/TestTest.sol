import {Test} from "./Test.sol";

contract TestTest is Test {

   function testTestSuccess(){
       assert(true, "");
   }

   function testTestFail(){
       assert(false, "Test failed.");
   }
}