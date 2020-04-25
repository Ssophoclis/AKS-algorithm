## AKS-algorithm
The implementation is taken from https://en.wikipedia.org/wiki/AKS_primality_test.

## Python:
Implementation of the AKS primality test algorithm in python. To test if a number n is prime, type ``` print(aks(n)) ```. For example, for n= 1009, ``` print(aks(1009))```.

## Go:
Implementation of the AKS primality test algorithm in Golang. To test if a number n is prime, in the main function, initialize the variable n to be the number you want to test. For example, for n= 1009, ```var n int64 = 1009```.

## TODO:
- [ ] Fix python implementation.
- [ ] Make Go implementation faster by using goroutines to the fastPoly function.


## Efficiency:
Note that this implementation of the algorithm is not as efficient as it could be. This implementation will detect if a number is coprime pretty fast. Although, if it is a prime, it will take quite some time. The main bottleneck is the fastPoly function and, more specifically, the multi function. In fastPoly we have to calculate (x+a)<sup>n</sup>, for multiple a. Currently, I am using fast modular exponentiation for polynomials to do this operation.
