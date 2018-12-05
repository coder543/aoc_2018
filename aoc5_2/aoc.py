from copy import deepcopy
polymer = list(open('input').read())


def elements(polymer):
    poly = deepcopy(polymer)
    poly = [p.lower() for p in poly]
    poly = sorted(poly)
    elements = list(set(poly))
    return elements


def react(polymer):
    # np = new_polymer
    np = [polymer[0]]
    for c in polymer[1:]:
        if len(np) > 0 and c.upper() == np[-1].upper() and c != np[-1]:
            np.pop()
            continue
        np.append(c)
    return np


elem = elements(polymer)
shortest = polymer
letter = None
for e in elem:
    poly = deepcopy(polymer)
    poly = list(filter(lambda c: c != e and c != e.upper(), poly))
    poly = react(poly)
    if len(poly) < len(shortest):
        shortest = poly
        letter = e

print(letter, len(shortest))
