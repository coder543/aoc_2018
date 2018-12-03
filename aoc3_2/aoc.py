#!/usr/bin/python3

from collections import namedtuple

data = open('input').read().split('\n')
claims = []
fabric = [[0 for _ in range(1000)] for _ in range(1000)]

Claim = namedtuple('Claim', ['id', 'left', 'top', 'width', 'height'])

for claim in data:
    claim = claim.replace(',', ' ').replace(':', '')
    claim = claim.replace(' @', '').replace('#', '')
    claim = claim.replace('x', ' ').split(' ')
    id = int(claim[0])
    left = int(claim[1])
    top = int(claim[2])
    width = int(claim[3])
    height = int(claim[4])
    claims.append(Claim(id, left, top, width, height))
    for row in range(top, top + height):
        for col in range(left, left + width):
            fabric[row][col] += 1


def overlapping(claim):
    bottom = claim.top + claim.height
    for row in fabric[claim.top:bottom]:
        right = claim.left + claim.width
        for val in row[claim.left:right]:
            if val != 1:
                return True
    return False


for claim in claims:
    if not overlapping(claim):
        print(claim.id)
        break
