#include <stdio.h>
#include <stdlib.h>
#include <string.h>

enum HandType {
  FiveOfKind = 1,
  FourOfKind = 2,
  FullHouse = 3,
  ThreeOfKind = 4,
  TwoPair = 5,
  OnePair = 6,
  High = 7
};

typedef struct hand {
  char *cards;
  uint bid;
  uint power;
  enum HandType type;
} hand;

void display(hand h) {
  printf("cards='%s', bid=%d, type=%d, power=%d\n", h.cards, h.bid, h.type,
         h.power);
}

void extractData(char *line, hand *o) {
  const char s[] = " ";
  hand h;

  char *token = strtok(line, s);

  char *cards = (char *)malloc(sizeof(token));
  strcpy(cards, token);

  h.cards = cards;
  h.type = High;
  h.power = 0;

  token = strtok(NULL, s);
  h.bid = atoi(token);

  *o = h;
}

// int representation of each card
unsigned int getCardRank(char c) {
  switch (c) {
  case 'A':
    return 1;
  case 'K':
    return 2;
  case 'Q':
    return 3;
  case 'J':
    return 4;
  case 'T':
    return 5;
  case '9':
    return 6;
  case '8':
    return 7;
  case '7':
    return 8;
  case '6':
    return 9;
  case '5':
    return 10;
  case '4':
    return 11;
  case '3':
    return 12;
  case '2':
    return 13;
  default:
    return 0;
  }
}

// push value using the left shift operator
void pushValue(uint *state, uint value) { *state |= (1 << value); }

// check if value pushed using bitwise AND
int isValue(uint state, uint value) { return (state & (1 << value)) != 0; }

void calcHand(hand *h) {
  uint state = 0;
  int a = 0, a_val = 0, b = 0, b_val = 0;

  for (uint i = 0; h->cards[i] != '\0'; i++) {
    uint value = getCardRank(h->cards[i]);
    (*h).power += value;

    if (isValue(state, value)) {
      if (a_val == 0) {
        a_val = value;
        a = 2;
      } else if (a_val == value) {
        a++;
      } else if (b_val == 0) {
        b_val = value;
        b = 2;
      } else {
        b++;
      }
    } else {
      pushValue(&state, value);
    }
  }

  if (a == 5 && b == 0) {
    (*h).type = FiveOfKind;
  } else if (a == 4 && b == 0) {
    (*h).type = FourOfKind;
  } else if ((a == 3 && b == 2) || (a == 2 && b == 3)) {
    (*h).type = FullHouse;
  } else if (a == 3 && b == 0) {
    (*h).type = ThreeOfKind;
  } else if (a == 2 && b == 2) {
    (*h).type = TwoPair;
  } else if (a == 2 && b == 0) {
    (*h).type = OnePair;
  }
}

int main() {
  FILE *filePtr = fopen("day7.txt", "r");
  char line[50];

  hand **hands = malloc(sizeof(hand *));
  uint hands_size = 0;

  if (hands == NULL) {
    printf("memory got burr");
    fclose(filePtr);
    return 1;
  }

  while (fgets(line, sizeof(line), filePtr)) {
    hand *h = (hand *)malloc(sizeof(hand));

    extractData(line, h);
    calcHand(h);

    hands = realloc(hands, (hands_size + 1) * sizeof(hand *));
    hands[hands_size] = h;

    hands_size++;
  }

  for (uint i = 0; i < hands_size; i++) {
    display(*hands[i]);
    free(hands[i]->cards);
    free(hands[i]);
  }

  fclose(filePtr);

  return 0;
}
