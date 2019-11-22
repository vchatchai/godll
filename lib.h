#include <string.h>
#include <stdio.h>
typedef struct{
    void* data;
    int size;
}GoMem;
void GetData(GoMem* gm);
void SetData(GoMem* gm);