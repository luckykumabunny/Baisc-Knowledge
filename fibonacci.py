def cal_fib(n):
    a=0
    b=1
    for i in range(2,n):
        c=a+b
        a=b
        b=c
    return c

def recur_fib(n):
    if n==1:
        return 0
    if n==2:
        return 1  
    return recur_fib(n-1)+recur_fib(n-2)

memory = {}
def recur_fib_dynamic(n):
    if n in memory:
        return memory[n]
    if n==1:
        return 0
    if n==2:
        return 1
    recur_value = recur_fib(n-1)+recur_fib(n-2)
    memory[n] = recur_value
    return recur_value



print(cal_fib(6))
print(recur_fib(6))
print(recur_fib_dynamic(6))
