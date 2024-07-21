#include <stdio.h>
#include <string.h>

#include "node.h"

node fromLine(char *line) {
  node n;

  // get label
  strncpy(n.label, line, 3);
  n.label[3] = '\0';

  // get left label
  char *leftLabelStr = strstr(line, "(");
  strncpy(n.leftLabel, leftLabelStr + 1, 3);
  n.leftLabel[3] = '\0';

  // get right label
  char *rightLabelStr = strstr(line, ", ");
  strncpy(n.rightLabel, rightLabelStr + 2, 3);
  n.rightLabel[3] = '\0';

  return n;
}

void display(node n) {
  printf("label='%s' left='%s' right='%s'\n", n.label, n.leftLabel,
         n.rightLabel);
}
