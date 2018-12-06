from scipy.spatial import ConvexHull

points = open('input').read().split('\n')
points = [point.split(', ') for point in points]
points = [[int(point[0]), int(point[1])] for point in points]

hull = ConvexHull(points)
out = open('hull', 'w')
for v in hull.vertices:
    print(v, file=out)
