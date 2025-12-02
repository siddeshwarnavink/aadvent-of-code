#ifndef NODE_H
#define NODE_H

typedef struct node {
  char label[4];
  char leftLabel[4];
  char rightLabel[4];
  struct node *left;
  struct node *right;
} node;

typedef struct node_list {
  struct node **list;
  int size;
} node_list;

node *from(char *line);
void display(node n);
node_list *create();
void add(node_list *list, node *n);
node *find(node_list *list, char *label);
void removeNode(node_list *list, char *label);
void clean(node_list *list);

#endif
