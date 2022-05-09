# golang-training
golang exercises of the course " Learn How To Code: Google's Go (golang) Programming Language"

## Ninja level 10

### EX1
- get this code working (https://go.dev/play/p/j-EA6003P0):
1. using func literal, aka, anonymous self-executing func
2. a buffered channel

### EX2
- Get this code running:
1. https://play.golang.org/p/oB-p3KMiH6
2. https://play.golang.org/p/_DBRueImEq

### EX3
Starting with this code (https://go.dev/play/p/sfyu4Is3FG), pull the values off the channel using a for range loop

### EX4
Starting with this code (https://go.dev/play/p/MvL6uamrJP), pull the values off the channel using a select statement

### EX5
Show the comma ok idiom starting with this code (https://go.dev/play/p/YHOMV9NYKK).

### EX6
Write a program that
- puts 100 numbers to a channel
- pull the numbers off the channel and print them

### EX7
Write a program that
- launches 10 goroutines
  - each goroutine adds 10 numbers to a channel
- pull the numbers off the channel and print them

### EX8
Using Fan in pattern, write a program that
- generate 100 numbers
  - send even to one channel
  - send odd to another
- pull the numbers off the channels and collect them in the same channel
- print the final numbers

### EX9
Using Fan out pattern, write a program that
- generate 100 numbers
- pull the numbers off the channel and create a goroutine for every number to do anything with them
  - send the result to result channel
- print all results