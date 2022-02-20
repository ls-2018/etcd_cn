import os


def get_files(path):
    for item in os.listdir(path):
        dir_file = os.path.join(path, item)
        if os.path.isfile(dir_file):
            yield dir_file
        else:
            for xx in get_files(dir_file):
                yield xx


main_set = set()
for file in get_files('.'):
    if file.endswith('py'):
        continue
    try:
        with open(file, encoding='utf8')as f:
            if 'package main' in f.read():
                main_set.add(os.path.dirname(file))
    except Exception:
        pass

for item in main_set:
    print('go build %s' % item)
