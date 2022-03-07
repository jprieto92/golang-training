# golang-training
golang exercises of the course " Learn How To Code: Google's Go (golang) Programming Language"

## Ninja level 5

### EX1
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
- first name
- last name
- favorite ice cream flavors

Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

### EX2
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

### EX3
1. Create a new type: vehicle.
- The underlying type is a struct.
- The fields:
  - doors
  - color
2. Create two new types: truck & sedan.
- The underlying type of each of these new types is a struct.
- Embed the “vehicle” type in both truck & sedan.
- Give truck the field “fourWheel” which will be set to bool.
- Give sedan the field “luxury” which will be set to bool. solution
3. Using the vehicle, truck, and sedan structs:
- using a composite literal, create a value of type truck and assign values to the fields;
- using a composite literal, create a value of type sedan and assign values to the fields.
4. Print out each of these values.
5. Print out a single field from each of these values.

### EX4
Create and use an anonymous struct.