#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *trimStr(char *s) {
  int s_len = strlen(s);
  int front = 0, rear = s_len - 1;

  while (s[front] == ' ') {
    front++;
  }

  while (s[rear] == ' ') {
    rear--;
  }

  int o_len = strlen(s) - (front + (s_len - rear));
  char *o = (char *)malloc(o_len * sizeof(char));
  int i = 0;

  while (front < rear) {
    o[i] = s[front];
    front++;
    i++;
  }

  return o;
}

void extractData(char *line) {
  line = strstr(line, ":");
  line = strstr(line, " ");
  line = trimStr(line);

  char *token;
  int i = 0;

  const char c[] = " ";
  token = strtok(line, c);

  while (token != NULL) {
    int val = atoi(token);
    printf("%d\n", val);

    token = strtok(NULL, c);
    i++;
  }
}

int main() {
  FILE *filePtr = fopen("day6.txt", "r");
  char *line;

  if (filePtr == NULL) {
    printf("Failed to read the file");
    exit(1);
  }

  fgets(line, 50, filePtr);
  extractData(line);

  fclose(filePtr);

  return 0;
}
