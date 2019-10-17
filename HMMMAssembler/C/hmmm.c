#include "hmmm.h"
#define LIN 20
int eval(char* s, node* regs, node* mem, stack* a, char* lines[20]);
int main(int argc, char* argv[]){
	stack* st=(stack*)malloc(sizeof(stack));
	stack_init(st);
	node* regs=(node*)malloc(sizeof(node));
	regs=createList(15, 'r');
	node* mem=(node*)malloc(sizeof(node));
	mem=createList(15, 'm');
	char lines[LIN][LIN];
	if(argc<2){
		int i=0;
		while(1){
			printf("ASSEMBLE!>");
			char line[20];
			fgets(line, sizeof(line), stdin);
			strncpy(lines[i++], line, 20);
			eval(line, regs, mem, st, lines);
			displayList(regs);
		}
	}else{
		char const* const fileName=argv[1];
		FILE* file=fopen(fileName, "r");
		char line[20];
		int i=0;
		while(fgets(line, sizeof(line), file)){
			strncpy(lines[i++], line, 20);
		}
		for(int i=0; i<20; i++){
			if(eval(clean(lines[i]), regs, mem, st, lines)==2){
				break;
			}
		}
		//displayList(regs);
		//displayList(mem);
		fclose(file);
	}
	return 0;
}
int eval(char* s, node* regs, node* mem, stack* a, char* lines[20]){
	if(strcmp("N", s)==0){
		return 2;
	}
	char** v=split_line(s);
	switch(v[0][0]){
		case 'l':
			if(v[0][4]=='r'){
				loadr(v[1], v[2], regs, mem);
				break;
			}else{
				loadn(v[1], v[2], regs, mem);
				break;
			}
		case 'o':
			or(v[1], v[2], v[3], regs);
			break;
		case 'x':
			xor(v[1], v[2], v[3], regs);
			break;
		case 's':
			if(v[0][1]=='u'){
				sub(v[1], v[2], v[3], regs);	
			}else if(v[0][1]=='e'){
				setn(v[1], v[2], regs);
			}else if(v[0][1]=='t'&&v[0][5]=='r'){
				storer(v[1], v[2], regs, mem);
			}else{
				storen(v[1], v[2], regs, mem);
			}
			break;
		case 'a':
			if(strcmp("and", v[0])==0){
				and(v[1], v[2], v[3], regs);
			}else{
			if(strcmp("addn", v[0])==0){
				addn(v[1], v[2], regs);
			}else{
				add(v[1], v[2], v[3], regs);
			}
			}
			break;
		case 'm':
			if(strcmp("mov", v[0])==0){
				copy(v[1], v[2], regs);
			}else{
			if(strcmp("mul", v[0])==0){
				mul(v[1], v[2], v[3], regs);
			}else{
				mod(v[1], v[2], v[3], regs);
			}}
			break;
		case 'c':
			copy(v[1], v[2], regs);
			break;
		case 'd':
			divide(v[1], v[2], v[3], regs);
			break;
		case 'n':
			if(strcmp("neg", v[0])==0){
				neg(v[1], v[2], regs);
			} //else: nop
			break;
		case 'e':
			break;
		case 'h':
			return 2;
			break;
		case 'r':
			read(v[1], regs);
			break;
		case 'w':
			write(v[1], regs);
			break;
		case 'p':
			if(v[0][1]=='u'){
				node* n=search(regs, v[1], 'n');
				push(a, n->val);
			}else{
				node *n=search(regs, v[1], 'n');
				n->val=pop(a);
			}
			break;
		case 'j':
			break;
		default:
			printf("Error: opcode ""%s"" not found\n", v[0]);
			break;
	}
	return 0;
}

