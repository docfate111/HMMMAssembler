#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include "linkedlist.h"
void loadr(char* rX, char* rY, node* regs, node* mem){
	node* varA=NULL; node* varB=NULL; node* varC=NULL;
	varA=search(regs, rX, 'n');
	varB=search(regs, rY, 'n');
	char str[20];
	sprintf(str, "%d", varB->val);
	varC=search(mem, str, 'n');
	if(varC==NULL){
		printf("%s is not a memory address\n", str);
	}else{
		varA->val=varC->val;
	}
}
void loadn(char* rX, char* n, node* regs, node* mem){
	node* varA=NULL; node* varC=NULL;
	varA=search(regs, rX, 'n');
	varC=search(mem, n, 'n');
	if(varC==NULL){
		printf("There is no memory address %s\n", n);
	}
	varA->val=varC->val;
}
void storer(char* rX, char* rY, node* regs, node* mem){
	node* varA=NULL; node* varB=NULL; node* varC=NULL;
	varA=search(regs, rX, 'n');
	varB=search(regs, rY, 'n');
	char str[20];
	sprintf(str, "%d", varB->val);
	varC=search(mem, str, 'n');
	if(varC==NULL){
		varC=search(mem, "", 'n');
	}
	varC->name=str;
	varC->val=varA->val;
}
void storen(char* rX, char* n, node* regs, node* mem){
	node* varA=NULL; node* varC=NULL;
	varA=search(regs, rX, 'n');
	varC=search(mem, n, 'n');
	if(varC==NULL){
		varC=search(mem, "", 'n');
	}
	varC->name=n;
	varC->val=varA->val;
}
char* clean(char* line){
	if(isalpha(line[0])||isalnum(line[0])){
	char* o=(char*)malloc(20);
	for(int i=0; i<20; i++){
		if((line[i]!='\n'&&line[i]!='\t')&&line[i]!='\r'){
			o[i]=line[i];
		}else{
			o[i]='\0';
		}
	}
	return o;
	}else{
		return "N";
	}
}
node* setn(char* rX, char* st, node* regs){
	node* var=NULL;
	var=search(regs, rX,'n');
	var->val=atoi(st);
	return var;
}
node* read(char* rX, node* regs){
	node* var=NULL;
	var=search(regs, rX,'n');
	int n;
	printf("Enter value of %s:\n", var->name);
	scanf("%d", &n);
	var->val=n;
	return var;
}
node* addn(char* rX, char* st, node* regs){
	node* var=NULL;
	var=search(regs, rX,'n');
	var->val=var->val+atoi(st);
	return var;
}
node* copy(char* rX, char* st, node* regs){
	node* varA=NULL;
	varA=search(regs, rX, 'n');
	node* varB=NULL;
	varB=search(regs, st, 'n');
	varA->val=varB->val;
	return varA;
}
node* neg(char* rX, char* rY, node* regs){
	node* varA=NULL;
	varA=search(regs, rX, 'n');
	node* varB=NULL;
	varB=search(regs, rY, 'n');
	int n=(-1)*varB->val;
	varA->val=n;
	return varA;
}
node* add(char* rX, char* rY, char* rZ, node* regs){
	node* varA=NULL;
	varA=search(regs, rX, 'n');
	node* varB=NULL;
	varB=search(regs, rY, 'n');
	node* varC=NULL;
	varC=search(regs, rZ, 'n');
	int new=(varB->val)+(varC->val);
	varA->val=new;
	return varA;
}
node* sub(char* rX, char* rY, char* rZ, node* regs){
	node* varA=NULL;
	varA=search(regs, rX, 'n');
	node* varB=NULL;
	varB=search(regs, rY, 'n');
	node* varC=NULL;
	varC=search(regs, rZ, 'n');
	int new=(varB->val)-(varC->val);
	varA->val=new;
	return varA;
}
node* mul(char* rX, char* rY, char* rZ, node* regs){
	node* varA=NULL;
	varA=search(regs, rX, 'n');
	node* varB=NULL;
	varB=search(regs, rY, 'n');
	node* varC=NULL;
	varC=search(regs, rZ, 'n');
	int new=(varB->val)*(varC->val);
	varA->val=new;
	return varA;
}
node* xor(char* rX, char* rY, char* rZ, node* regs){
	node* varA=NULL;
	varA=search(regs, rX, 'n');
	node* varB=NULL;
	varB=search(regs, rY, 'n');
	node* varC=NULL;
	varC=search(regs, rZ, 'n');
	int new=(varB->val)^(varC->val);
	varA->val=new;
	return varA;
}
node* and(char* rX, char* rY, char* rZ, node* regs){
	node* varA=NULL;
	varA=search(regs, rX, 'n');
	node* varB=NULL;
	varB=search(regs, rY, 'n');
	node* varC=NULL;
	varC=search(regs, rZ, 'n');
	int new=(varB->val)&&(varC->val);
	varA->val=new;
	return varA;
}
node* or(char* rX, char* rY, char* rZ, node* regs){
	node* varA=NULL;
	varA=search(regs, rX, 'n');
	node* varB=NULL;
	varB=search(regs, rY, 'n');
	node* varC=NULL;
	varC=search(regs, rZ, 'n');
	int new=(varB->val)||(varC->val);
	varA->val=new;
	return varA;
}
node* mod(char* rX, char* rY, char* rZ, node* regs){
	node* varA=NULL;
	varA=search(regs, rX, 'n');
	node* varB=NULL;
	varB=search(regs, rY, 'n');
	node* varC=NULL;
	varC=search(regs, rZ, 'n');
	int new=(varB->val)%(varC->val);
	varA->val=new;
	return varA;
}
node* divide(char* rX, char* rY, char* rZ, node* regs){
	node* varA=NULL;
	varA=search(regs, rX, 'n');
	node* varB=NULL;
	varB=search(regs, rY, 'n');
	node* varC=NULL;
	varC=search(regs, rZ, 'n');
	int new=(varB->val)/(varC->val);
	varA->val=new;
	return varA;
}
node* write(char* rX, node* regs){
	node* var=NULL;
	var=search(regs, rX, 'v');
	printf("\n");
	return var;
}
char** split_line(char * s){
	char** res=NULL;
	char* p=strtok(s," ");
	int n_spaces=0;
	/* split string and append tokens to 'res' */
	while(p){
  		res=realloc(res, sizeof (char*) * ++n_spaces);
  		if(res==NULL)
    		exit (-1); /* memory allocation failed */
  		res[n_spaces-1]=p;
  		p=strtok(NULL," ");
	}
	/* realloc one extra element for the last NULL */
	res=realloc(res, sizeof (char*) * (n_spaces+1));
	res[n_spaces]=0;
	/* print the result 
	for (i = 0; i < (n_spaces+1); ++i)
  		printf ("res[%d] = %s\n", i, res[i]);
	free the memory allocated */
	return res;
	free(res);
}
