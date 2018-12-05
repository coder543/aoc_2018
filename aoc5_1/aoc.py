polymer = list(open('input').read())


def react(polymer):
    # np = new_polymer
    np = [polymer[0]]
    for c in polymer[1:]:
        if len(np) > 0 and c.upper() == np[-1].upper() and c != np[-1]:
            np.pop()
            continue
        np.append(c)
    return np


polymer = react(polymer)

print(len(polymer))
