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
  node_list *list = create();

  while (fgets(line, sizeof(line), file)) {
    node *n = from(line);
    add(list, n);
  }

  // Map nodes
  for (int i = 0; i < list->size; i++) {
    node *n = list->list[i];
    node *left = find(list, n->leftLabel);
    node *right = find(list, n->rightLabel);

    n->left = left;
    n->right = right;
 
    display(*n);
  }

  fclose(file);
  free(seq);
  clean(list);

  return 0;
}
