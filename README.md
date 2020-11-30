```
         |
        -+-
         A
        /=\               /\  /\    ___  _ __  _ __ __    __
      i/ O \i            /  \/  \  / _ \| '__|| '__|\ \  / /
      /=====\           / /\  /\ \|  __/| |   | |    \ \/ /
      /  i  \           \ \ \/ / / \___/|_|   |_|     \  /
    i/ O * O \i                                       / /
    /=========\        __  __                        /_/    _
    /  *   *  \        \ \/ /        /\  /\    __ _  ____  | |
  i/ O   i   O \i       \  /   __   /  \/  \  / _` |/ ___\ |_|
  /=============\       /  \  |__| / /\  /\ \| (_| |\___ \  _
  /  O   i   O  \      /_/\_\      \ \ \/ / / \__,_|\____/ |_|
i/ *   O   O   * \i
/=================\
       |___|
cred: https://www.asciiart.eu/
```
# AOC-2020
My collection of solutions to www.adventofcode.com

Solutions are made in Go, using the helper [aocurl](https://github.com/IAmBullsaw/aocurl) to retrieve and use the inputs.

## Running the days
The template is runnable with `aocurl year day -io | go run main.go testcases.go` and will first run all tests, and if they pass it will run the solution on the puzzle

If the tests fail, it will exit with exit code 1.

## Snippet to copy the template for us lazy people
`mkdir day01 && cp template/* day01/`
