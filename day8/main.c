#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "node.h"

int main() {
  FILE *file = fopen("day8.txt", "r");
  char line[25];

  fgets(line, sizeof(line), file);
  
  char *seq = (char *)malloc(sizeof(line));
  int seq_len = strlen(line);
  strncpy(seq, line, seq_len - 1);
  seq[seq_len - 1] = '\0';
  seq_len--;
  
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

    // display(*n);
  }

  // Start traversing
  node *ptr = find(list, "AAA");
  if (ptr == NULL) {
    printf("AAA not found!\n");
    fclose(file);
    free(seq);
    clean(list);
    return 1;
  }
  
  int i = 0, steps = 0;
  while (strcmp(ptr->label, "ZZZ") != 0) {
    if (seq[i] == 'L') {
      ptr = ptr->left;
    } else {
      ptr = ptr->right;
    }

    i = (i + 1) % seq_len;
    steps++;
  }

  printf("steps = %d", steps);

  fclose(file);
  free(seq);
  clean(list);

  return 0;
}
