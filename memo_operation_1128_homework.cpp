#include <iostream>
#include <cstring>

using namespace std;

int main()
{
	// Quesetion one printf first address
	// “hello” 字符常量返回的就是首地址，所以可以赋值给指针变量
	char* str = NULL;
	str = (char*)malloc(sizeof(char*)*14);
	strncpy(str, "Hello,World.", 13);
	printf("Address of First Character is %p", str);
	
	// Quesetion two using point chage the value of last character
	// "" 长字符，末尾包含/0 ‘’就是一个char
	*(str+11) = '!';
	
	// Quesetion three malloc 1024 bytes memeory and free memory
	char* big_str = NULL;
	big_str = (char*)malloc(sizeof(char*)*1024);
	char* p = str;
	char* big_p = big_str+9;
	while(*p){
		*(big_p++) = *(p++);
	}
	
	free(str);
	free(big_str);
	
   return 0;
}