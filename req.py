import os

for i in range(10000):
    os.system("etcdctl put a%s b%s" % (i, i))
