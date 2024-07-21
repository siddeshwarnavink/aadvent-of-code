#ifndef NODE_H
#define NODE_H

typedef struct node {
  char label[4];
  char leftLabel[4];
  char rightLabel[4];
  struct node *left;
  struct node *right;
} node;

node fromLine(char *line);
void display(node n);

#endif
