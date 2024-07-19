#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/*enum HandType {*/
/*  FiveOfKind = 1,*/
/*  FourOfKind = 2,*/
/*  FullHouse = 3,*/
/*  ThreeOfKind = 4,*/
/*  TwoPair = 5,*/
/*  OnePair = 6,*/
/*  High = 7*/
/*};*/

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

  char *cards = (char *)malloc(sizeof(token));
  strcpy(cards, token);
  h.cards = cards;

  token = strtok(NULL, s);
  h.bid = atoi(token);
  h.rank = 0;

  *o = h;
}

void rankHand(hand *h) {
  // determine the hand type
  for (int i = 0; h->cards[i] != '\0'; i++) {
    printf("%c\t", h->cards[i]);
  }
  printf("\n");
}

int main() {
  FILE *filePtr = fopen("day7.txt", "r");
  char line[50];

  hand **hands = malloc(sizeof(hand *));
  int hands_size = 0;

  if (hands == NULL) {
    printf("memory got burr");
    fclose(filePtr);
    return 1;
  }

  while (fgets(line, sizeof(line), filePtr)) {
    hand *h = (hand *)malloc(sizeof(hand));
    extractData(line, h);
    rankHand(h);

    hands = realloc(hands, (hands_size + 1) * sizeof(hand *));
    hands[hands_size] = h;

    hands_size++;
  }

  for (int i = 0; i < hands_size; i++) {
    display(*hands[i]);
    free(hands[i]->cards);
    free(hands[i]);
  }

  fclose(filePtr);

  return 0;
}
