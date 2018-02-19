import string
import collections

def is_pangram(sentence):
    counter_dict = collections.Counter(sentence.lower())
    for c in string.ascii_lowercase:
        if c not in counter_dict:
            return False
    return True

