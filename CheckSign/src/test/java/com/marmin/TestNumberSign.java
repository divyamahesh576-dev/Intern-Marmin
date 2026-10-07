package com.marmin;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Assertions;

class TestNumberSign {

    NumberSign num = new NumberSign();

    @Test
	void Test1(){
        String Expected = num.Numbers(-1);
        String Actual = "negative";

		Assertions.assertEquals(Expected,Actual);
    }

    @Test
	void Test2(){
        String Expected = num.Numbers(2);
        String Actual = "positive";

		Assertions.assertEquals(Expected,Actual);
    }

    @Test
	void Test0(){
        String Expected = num.Numbers(0);
        String Actual = "zero";

		Assertions.assertEquals(Expected,Actual);
    }

}
