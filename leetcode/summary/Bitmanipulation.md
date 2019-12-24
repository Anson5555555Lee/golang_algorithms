#### Number Reverse
Base-10: 
````text
ans = ans*10 + n%10
n /= 10
````
Base-2 is exactly the same if n is positive:
````text
ans = ans*2 + n %2
n /= 2

or use bit operators:

ans = (ans << 1) | (n & 1)
n >>= 1
````
But need to handle leading zeros:
````text
00110 => 01100 not 110
````