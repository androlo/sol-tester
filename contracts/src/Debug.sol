contract Debug {
    event RuntimeEvent(bool indexed result, string message);
    event StackEvent(uint indexed type, bytes32 indexed expected, bytes32 indexed value)
}