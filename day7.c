#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct hand {
  char *cards;
  int bid;
  int rank;
} hand;

void display(hand h) {
  printf("cards='%s', bid=%d, rank=%d\n", h.cards, h.bid, h.rank);
}

void extractData(char *line, hand *o) {
  const char s[] = " ";
  hand h;

  char *token = strtok(line, s);
  strcpy(h.cards, token);

  token = strtok(NULL, s);
  h.bid = atoi(token);

  *o = h;
}

int main() {
  FILE *filePtr = fopen("day7.txt", "r");
  char line[50];
  hand *h = (hand *)malloc(sizeof(hand));

  while (fgets(line, sizeof(line), filePtr)) {
    extractData(line, h);
    display(*h);
  }

  fclose(filePtr);
  free(h);

  return 0;
}
