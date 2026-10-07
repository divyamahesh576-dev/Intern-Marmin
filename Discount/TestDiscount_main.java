package com.marmin;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Assertions;

class TestDiscount_main {

    Discountnum num = new Discountnum();

    @Test
	void Test1(){
        int Expected = num.Numbers(500);
        int Actual = 450;

		Assertions.assertEquals(Expected,Actual);
    }

    @Test
	void Test2(){
        int Expected = num.Numbers(2);
        int Actual = 2;

		Assertions.assertEquals(Expected,Actual);
    }

    @Test
	void Test0(){
        int Expected = num.Numbers(0);
        int Actual = 0;

		Assertions.assertEquals(Expected,Actual);
    }

}
