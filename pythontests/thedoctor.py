for k in range(int(input())):
    i = int(input())
    if i % 10 != 0: print(i * (i - 1)//2)
    else: print(i)
