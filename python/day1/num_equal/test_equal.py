from num_equal.equal import equal


def test_10_equal_10():
    assert equal(10, 10) == True

def test_10_not_equal_1():
    assert equal(10, 1) == False

def test_minus_10_equal_minus_10():
    assert equal(-10, -10) == True

def test_number_10_not_equals_string_10():
    assert equal(10, "10") == False

def test_string_10_not_equals_number_10():
    assert equal("10",10) == False

def test__not_equals_number_10():
    assert equal("10",10) == False

def test__both_none():
    assert equal(None, None) == True