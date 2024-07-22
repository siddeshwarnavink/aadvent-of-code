#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <sys/types.h>

#include "node.h"

node_list *create() {
  node_list *list = (node_list *)malloc(sizeof(node_list));
  list->list = (node **)malloc(sizeof(node *));
  list->size = 0;
  return list;
}

void add(node_list *list, node *n) {
  list->size++;
  list->list = realloc(list->list, list->size * sizeof(node *));
  list->list[list->size - 1] = n;
}

node *find(node_list *list, char *label) {
  node *n;
  for (int i = 0; i < list->size; i++) {
    if (strcmp(list->list[i]->label, label) == 0) {
      n = list->list[i];
      return n;
    }
  }
  return n;
}

void removeNode(node_list *list, char *label) {
 int n = 0;
  for (int i = 0; i < list->size; i++) {
    if (strcmp(list->list[i]->label, label) == 0) {
      n = i;
    }
  }

  if (n > 0) {
    free(list->list[n]);

    // move all elements one step back
    for (int i = n; i < list->size - 1; i++) {
      list->list[i] = list->list[i + 1];
    }

    list->size -= 1;
    list->list[list->size] = NULL;
  }
}

void clean(node_list *list) {
  for(int i = 0; i < list->size; i++) {
    free(list->list[i]);
  }
  free(list->list);
  free(list);
}

node *from(char *line) {
  node *n = (node *)malloc(sizeof(node));

  // get label
  strncpy(n->label, line, 3);
  n->label[3] = '\0';

  // get left label
  char *leftLabelStr = strstr(line, "(");
  strncpy(n->leftLabel, leftLabelStr + 1, 3);
  n->leftLabel[3] = '\0';

  // get right label
  char *rightLabelStr = strstr(line, ", ");
  strncpy(n->rightLabel, rightLabelStr + 2, 3);
  n->rightLabel[3] = '\0';

  return n;
}

void display(node n) {
  int left = n.left == NULL ? 0 : 1;
  int right = n.right == NULL ? 0 : 1;
  printf("label='%s' left='%s'(%d) right='%s'(%d)\n", n.label, n.leftLabel, left,  n.rightLabel, right);
}
