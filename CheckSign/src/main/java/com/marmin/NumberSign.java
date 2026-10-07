package com.marmin;

 public class NumberSign
{
   public String Numbers( int n )
    {
        if(n>0){
            return "positive";
        }
        else if(n<0){
            return "negative";
        }
        else{
            return "zero";
        }
    }
}
