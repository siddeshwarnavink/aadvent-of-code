#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "node.h"

int main() {
  FILE *file = fopen("day8.txt", "r");
  char line[25];

  fgets(line, sizeof(line), file);
  char *seq = (char *)malloc(sizeof(line));
  strcpy(seq, seq);

  fgets(line, sizeof(line), file);

  while (fgets(line, sizeof(line), file)) {
    node n = fromLine(line);
    display(n);
  }

  fclose(file);
  free(seq);

  return 0;
}
