#!/usr/bin/python3

data = open('input').read().split('\n')
fabric = [[0 for _ in range(1000)] for _ in range(1000)]

for claim in data:
    claim = claim.replace(',', ' ').replace(':', '')
    claim = claim.replace(' @', '').replace('#', '')
    claim = claim.replace('x', ' ').split(' ')
    claim_num = int(claim[0])
    left = int(claim[1])
    top = int(claim[2])
    width = int(claim[3])
    height = int(claim[4])
    for row in range(top, top + height):
        for col in range(left, left + width):
            fabric[row][col] += 1

sqin = 0
for row in fabric:
    for val in row:
        if val > 1:
            sqin += 1

print(sqin)
