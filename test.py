import numpy as np
import sys


def main():
    sys.setrecursionlimit(150000)
    inp = """7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3"""
    inp = inp.split()
    corner_points = []
    x_coords = set()
    y_coords = set()

    for coord in inp:
        coord = coord.split(",")
        x = int(coord[0])
        y = int(coord[1])
        x_coords.add(x)
        y_coords.add(y)
        corner_points.append([x, y])
    x_coords = sorted(list(x_coords))
    y_coords = sorted(list(y_coords))
    x_dict = {}
    for i in range(len(x_coords)):
        x_dict[x_coords[i]] = i
    y_dict = {}
    for i in range(len(y_coords)):
        y_dict[y_coords[i]] = i
    old_corners = []
    for i in range(len(corner_points)):
        point = corner_points[i]
        x = point[0]
        y = point[1]
        old_corners.append(point)
        corner_points[i] = [x_dict[x], y_dict[y]]

    max_x = 0
    max_y = 0
    for point in corner_points:
        if point[0] > max_x:
            max_x = point[0]
        if point[1] > max_y:
            max_y = point[1]
    grid = np.zeros((max_x + 1, max_y + 1), dtype=np.uint8)
    for i in range(len(corner_points) - 1):
        point_a = corner_points[i]
        point_b = corner_points[i + 1]
        x_diff = point_b[0] - point_a[0]
        y_diff = point_b[1] - point_a[1]
        if x_diff != 0:
            y = point_a[1]
            x = point_a[0] if x_diff > 0 else point_b[0]
            x_diff = abs(x_diff)

            for x in range(x, x + x_diff + 1):
                grid[x][y] = 69
        else:
            x = point_a[0]
            y = point_a[1] if y_diff > 0 else point_b[1]
            y_diff = abs(y_diff)

            for y in range(y, y + y_diff + 1):
                grid[x][y] = 69

    point_a = corner_points[0]
    point_b = corner_points[-1]
    x_diff = point_b[0] - point_a[0]
    y_diff = point_b[1] - point_a[1]
    if x_diff != 0:
        y = point_a[1]
        x = point_a[0] if x_diff > 0 else point_b[0]
        x_diff = abs(x_diff)

        for x in range(x, x + x_diff + 1):
            grid[x][y] = 69
    else:
        x = point_a[0]
        y = point_a[1] if y_diff > 0 else point_b[1]
        y_diff = abs(y_diff)

        for y in range(y, y + y_diff + 1):
            grid[x][y] = 69
    for corner in corner_points:
        grid[corner[0]][corner[1]] = 69
    flood_fill(grid, 2, 1)
    largest_ar = 0
    for i in range(len(corner_points)):
        for j in range(i + 1, len(corner_points)):
            point_a = corner_points[i]
            point_d = corner_points[j]
            ar = area(old_corners[i], old_corners[j])
            if ar > largest_ar:
                x_min = min(point_a[0], point_d[0])
                y_min = min(point_a[1], point_d[1])
                x_max = max(point_a[0], point_d[0])
                y_max = max(point_a[1], point_d[1])

                point_a = [x_min, y_min]
                point_b = [x_max, y_min]
                point_c = [x_max, y_max]
                point_d = [x_min, y_max]
                perimeter = set()
                perimeter = perimeter | make_edge(point_a, point_b)
                perimeter = perimeter | make_edge(point_b, point_c)
                perimeter = perimeter | make_edge(point_c, point_d)
                perimeter = perimeter | make_edge(point_d, point_a)
                for point in perimeter:
                    x = point[0]
                    y = point[1]
                    if grid[x][y] == 0:
                        break
                else:
                    largest_ar = ar
    print(largest_ar)


def make_edge(point_a, point_b):
    x1 = point_a[0]
    x2 = point_b[0]
    y1 = point_a[1]
    y2 = point_b[1]
    edge = set()
    if x1 > x2:
        for x in range(x2, x1 + 1):
            edge.add((x, y1))
    elif x1 < x2:
        for x in range(x1, x2 + 1):
            edge.add((x, y1))
    elif y1 > y2:
        for y in range(y2, y1 + 1):
            edge.add((x1, y))
    else:
        for y in range(y1, y2 + 1):
            edge.add((x1, y))
    return edge


def dfs(grid, i, j):
    n = len(grid)
    m = len((grid[0]))

    if i < 0 or i >= n or j < 0 or j >= m or grid[i][j] > 0:
        return
    else:
        grid[i][j] = 69
        dfs(grid, i + 1, j)
        dfs(grid, i - 1, j)
        dfs(grid, i, j + 1)
        dfs(grid, i, j - 1)


def flood_fill(grid, i, j):
    if grid[i][j] > 0:
        return
    dfs(grid, i, j)


def area(point_a, point_b):
    return (abs(point_a[0] - point_b[0]) + 1) * (abs(point_a[1] - point_b[1]) + 1)


if __name__ == "__main__":
    main()
