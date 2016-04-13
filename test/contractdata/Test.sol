contract Test {

    event TestEvent(bool indexed result, string message);

    function assert(bool value, string message) {
        if (value)
            TestEvent(true, "");
        else
            TestEvent(false, message);
    }
}