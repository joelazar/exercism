def verify(isbn):
    isbnlist = list(isbn)
    if len(isbnlist) not in [10, 13]:
        return False
    if "-" in isbnlist:
        for i in sorted([1, 5, 11], reverse=True):
            del isbnlist[i]
    summa = 0
    for i in range(1, 11):
        try:
            summa += int(isbnlist[i-1])*(11-i)
        except ValueError:
            if i == 10 and isbnlist[i-1] == "X":
                summa += 10*(11-i)
            else:
                return False
    return summa%11 == 0
