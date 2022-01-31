In an arbitrary sized rectangle filled with random number of blocks, find out the largest sized square inside it (from most top-left corner) and draw it on stdout.

## The rules

The program must take in filename as an argument. The file passed to the program contains metadata in the first line and the actual (possible) rectangle after that,
e.g.:

```
9.O*
........O...
............
.O..........
............
....O.......
............
..O.........
........O...
............
```

The correct solution to the example would be:

```
........O...
.....******.
.O...******.
.....******.
....O******.
.....******.
..O..******.
........O...
............
```

While this would be incorrect, due to the `from most top-left corner` rule:

```
........O...
......******
.O....******
......******
....O.******
......******
..O...******
........O...
............
```

Each line is null-terminated and example boards can be found in `boards`-folder. The program has to be implemented in `main.go`. Once ready, open a pull-request. The program has to pass the pipeline in order to be successful.

## The format of the metadata-line

First characters (0-9) describes the height of the rectangle, followed by character used for empty space, character for block and character for drawing the square. For instance,

`12+o-`

Would state that the rectangle's height is 12 rows, uses `+` for empty space, `o` for the blocks and `-` for the square. Given the previous file-example, with meta-line `9+o-`, the solved outcome would be:

```
++++++++o+++
+++++------+
+o+++------+
+++++------+
++++o------+
+++++------+
++o++------+
++++++++o+++
++++++++++++
```
