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

## nightlyscript.sh
This script is supposed to be edited with YOUR paths and then when run, it will:
* create `repo/day<current day>`
* copy `repo/template` into `repo/day<current day>`
* create `repo/day<current day>/Makefile`

This allows you to be ready when the gun goes off and only run `make` in `repo/day<current day>` to run your solution against the testcases, and if they pass, run your solution with the current input.
