# NINTENDO CHALLENGE

Int this repo i will attempt to solve the nintendo challenge as given by:

## Challenge

The SETI program has received a series of messages from Alpha Centauri. The most frequent message seems to be: 541a4231 5d324646 27219a26 12497b0e 724eddcb 0e131617 9521bedf 55544dc7
It is not known why these messages are encoded, but there is a good chance that the Alpha Centaurians are trying to evaluate our cognitive abilities before establishing advanced contact.
Our best engineers are working to decode these messages, and they've already succeeded at identifying the program that the Centaurians use to encode the messages. This program takes a size and a list of numbers as inputs. It then outputs the encoded message (see the pseudo-code below).
But so far, no one has been able to decode the messages. We are well aware that this task is by far the hardest that we’ve encountered, and that only a true NERD will be able to pull it off!

Here is a pseudo-code version of the encoding program:
 
READ size
READ size / 16 integers in array a
WRITE size / 16 zeros in array b

For i from 0 to size - 1:
    For j from 0 to size - 1:
        b[(i+j)/32] ^= ((a[i/32] >> (i%32)) & (a[j/32 + size/32] >> (j%32)) & 1) << ((i+j)%32)

PRINT b

You can see a C++ version of the program [Here](cpp/encoder.cpp).
You can see a Go version of the program [Here](go/encoder.go).

The goal is to determine the series of numbers entered (array a) from the encoded output of the program (array b). The numbers – input and output – should be displayed in hexadecimal, 8 characters padded with 0 (for example, 42 would be displayed as 0000002a).
If you pass the output of your program as input to the encoder above, you should obtain the input provided to your program.

If there are several possible decoded values, you should display all the possibilities in alphabetical order.

## Approach

Since I know  nothing about c++ I will use golang to model and solve the problem first
