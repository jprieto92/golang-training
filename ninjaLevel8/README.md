# golang-training
golang exercises of the course " Learn How To Code: Google's Go (golang) Programming Language"

## Ninja level 8

### EX1
Starting with this code (https://go.dev/play/p/_fQUGm9Utvl), marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.

### EX2
Starting with this code (https://go.dev/play/p/b_UuCcZag9), unmarshal the JSON into a Go data structure. Hint:
- https://mholt.github.io/json-to-go/
- code:
  - code setup (just fyi, not needed for exercise):https://play.golang.org/p/nWPP5Z2Q4e

### EX3
Starting with this code (https://go.dev/play/p/BVRZTdlUZ_), encode to JSON the []user sending the results to Stdout. Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})

### EX4
Starting with this code (https://go.dev/play/p/H_q75mpmHW), sort the []int and []string for each person.

### EX5
Starting with this code (https://go.dev/play/p/BVRZTdlUZ_), sort the []user by
- age
- last

Also sort each []string “Sayings” for each user
- print everything in a way that is pleasant.
  
