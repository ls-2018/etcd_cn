import os

for i in range(100000):
    print(i)
    os.system("etcdctl put a%s b%s" % (i, i))
