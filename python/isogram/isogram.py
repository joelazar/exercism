import collections

def is_isogram(string):
    if len(string) != 0:
        counter = collections.Counter(string.lower().replace(" ","").replace("-",""))
        return counter.most_common(1)[0][1] == 1
    else:
        return True
