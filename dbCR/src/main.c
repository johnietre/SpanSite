#include "libdb.h"
#include <stdio.h>

int Primec(int n) {
  if (n < 1) return 0;
  else if (n == 1) return 2;
  else if (n == 2) return 3;
  else if (n == 3) return 5;
  else if (n == 4) return 7;
  // prime keeps track of if the number is prime or not
  // counter keeps track of the number of primes encountered
  unsigned short prime = 1, counter = 4;
  // Increase by 2 so that there will only be odd numbers
  for (int i = 9; ; i += 2) {
    prime = 1;
    // Divide i by 2 (add one if result is even)
    int j = ((i / 2) % 2) ? i / 2 : (i / 2) + 1;
    for (; j > 1; j -= 2) {
      if (i % j == 0) {
        prime = 0;
        break;
      }
    }
    // Increment the counter if the number was prime
    counter += prime;
    if (counter == n) return i;
  }
}

int main(int argc, char *argv[]) {
  // printf("%i\n", Primec(10000));
  printf("%i\n", Primer(10000));
  return 0;
}
