# Summation of primes
#
# The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
# Find the sum of all the primes below two million.

# https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
# man I can't believe I actually read through this
# turns out, C doesn't have easy lists so Python'll do
# run time: 11113s~ 

# generate numbers
numbers = list(); size = 2000000
for i in range(2,size):
    numbers.append(i)

# Sieve
sum = 0
for number in numbers:
    for sub in numbers:
        if number < sub and sub % number == 0:
            numbers.remove(sub)
    sum += number

print(numbers)
print("Sum of primes;",sum)
#end res: 142913828922
