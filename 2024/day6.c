#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void trimStr(char *s, char **o) {
  int s_len = strlen(s);
  int front = 0, rear = s_len - 1;

  while (s[front] == ' ') {
    front++;
  }

  while (s[rear] == ' ') {
    rear--;
  }

  int o_len = strlen(s) - (front + (s_len - rear)) + 1;
  *o = (char *)malloc(o_len * sizeof(char));
  int i = 0;

  while (front <= rear) {
    (*o)[i] = s[front];
    front++;
    i++;
  }

  (*o)[i] = '\0';
}

int size = 0;

int *extractData(char *line) {
  char *trimLine;

  line = strstr(strstr(line, ":") + 1, " ");
  trimStr(line, &trimLine);

  char *token;
  int i = 0;
  int *o = (int *)malloc(1 * sizeof(int));

  if (o == NULL) {
    printf("memory got burr\n");
    exit(1);
  }

  const char c[] = " ";
  token = strtok(trimLine, c);

  while (token != NULL) {
    int val = atoi(token);

    o = (int *)realloc(o, (i + 1) * sizeof(int));
    if (o == NULL) {
      printf("memory got burr\n");
      exit(1);
    }

    o[i] = val;

    token = strtok(NULL, c);
    i++;
  }

  size = i;

  free(trimLine);
  free(token);

  return o;
}

int main() {
  FILE *filePtr = fopen("day6.txt", "r");
  char *line;
  int *timeData = NULL, *distanceData = NULL;
  int result = 1;

  if (filePtr == NULL) {
    printf("Failed to read the file");
    exit(1);
  }

  fgets(line, 50, filePtr);
  timeData = extractData(line);

  fgets(line, 50, filePtr);
  distanceData = extractData(line);

  for (int i = 0; i < size; i++) {
    int count = 0;
    int holdTime = 1;

    while (holdTime < timeData[i]) {
      int distance = holdTime * (timeData[i] - holdTime);

      if (distance > distanceData[i]) {
        count++;
      }
      holdTime++;
    }

    result *= count;
  }

  printf("%d", result);

  fclose(filePtr);
  free(timeData);
  free(distanceData);

  return 0;
}
