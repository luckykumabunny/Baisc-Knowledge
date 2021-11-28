#include <iostream>
using namespace std;

int main()
{
   string A = "Hello,World.";
	cout << "the starting locstion of A is " << &A; 
	
	int len = A.length(); 
	A.erase(len-1);
	string B = "!";
	int len2 = A.length(); 
	A.insert(len2, B);
	cout << "the new A is " << A; 
	
	char C[13] = "Hello,World!";
	char* D = C;
	std::allocator<char> alloc;
	char* buffer = alloc.allocate(1024);
	
	buffer[0] = C[0];
	int i = 9;
	while(*D){
		buffer[i] = *(D++);
	}
   return 0;
}