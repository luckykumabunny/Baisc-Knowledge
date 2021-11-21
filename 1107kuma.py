def max(a, b, c):
    if a > b:
        if a > c:
            return a
        return c
    elif a <= b:
        if b > c:
            return b
        return c

def max_v2(a, b, c):
    d = a if a>b else b
    return d if d>c else c

# print(max(1,2,3))
# print(max_v2(1,2,3))

a = 4
b = 3
a = a/2
b = b >> 1
print(a)
print(b)

c= 10# {"1":2}
if c:
    print(True)
d = [1,2,3,4]
d = d[:3]
print(d)

for i in range(0, 4):
    print(i)

# if __name__ == "__main__":

    