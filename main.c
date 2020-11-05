#include "libtest.h"
#include <stdio.h>

int Primec(int n) {
  int prime = 0;
  int counter = 4;
  for (int i = 0; ; i += 2) {
    int j = i / 2;
    if (j % 2 == 0) j++;
    for (; j > 1 && prime; j -= 2) {
      if (i % j) prime = 0;
    }
    counter += prime;
    if (counter == n) {
      return i;
    }
  }
}

int main(int argc, char *argv[]) {
  printf("%i\n", Primec(1000));
  printf("%i\n", Primer(1000));
  return 0;
}
