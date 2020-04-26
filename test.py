import math
import array as arr
from fractions import gcd

def fastPoly(base,power,r):
    """Use fast modular exponentiation for polynomials to raise them to a big power.
    """
    x = arr.array('d',[],)
    for i in range(r):
        x.append(0)
    x[(0)] = 1 
    n = power
    
    while power > 0:
        if power % 2 == 1: 
            x = multi(x,base,n,r)
        base = multi(base,base,n,r)
        power = power // 2        
    return(x)


def multi(a,b,n,r):
    """Function used by fastPoly to multiply two polynomials together.
    """ 
    x = arr.array('d',[])
    for i in range(r):
        x.append(0)
    for i in range(len(a)):
        for j in range(len(b)):
            x[(i+j) % r ] += a[(i)] * b[(j)] 
            #x[(i+j) % r] = x[(i+j) % r] % n 
    for i in range(r,len(x)):
            x=x[:-1]
    return(x)

print(fastPoly([1,1],31,29))