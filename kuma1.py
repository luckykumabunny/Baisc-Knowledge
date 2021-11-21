import random
# 遍历
activities = ['python', 'papa', 'guitar', 'facial', 'decoration', 'movie', 'muscle', 'delicious', 'shopping']
#  键值对
date = {"mor": 1, "noon":2, "eve":3}
date["mor"] = 4
print(date)

for key, val in date.items():
    date[key] = random.choice(activities)
print("first run")
print(date)

print("sec run")
date["morning"] = random.choice(activities)
date["noon"] = random.choice(activities)
print(date)

a = [1,2,3,4,5,6,8,9]

for i in range(0, len(a)):
    a[i] = a[i]*a[i]
print(a)

a = [1,2,3,4,5,6,8,9]

b = [dog * dog for dog in a]
print(b)

if "flower dog" is not "flower dog":
    print(True)