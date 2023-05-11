Compute Pascal's triangle up to a given number of rows. Return a list of lists, where each list represents a successive row from Pascal's triangle.

In Pascal's Triangle each number is computed by adding the numbers to the right and left of the current position in the previous row. For example, Pascal's triangle with 5 rows can be visualized as:

```
    1
   1 1
  1 2 1
 1 3 3 1
1 4 6 4 1
# ... etc
```

## Test Cases

```
pascals(1)  // [[1]]
pascals(2)  // [[1], [1, 1]]
pascals(3)  // [[1], [1, 1], [1, 2, 1]]
pascals(4)  // [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1]]
pascals(5)  // [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]
```
