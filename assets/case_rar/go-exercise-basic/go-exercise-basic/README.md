## Atention

### Exercise 1 :

- You can find the solution in the file `calculate_area.go`
- Calculate the area of a circle, a square and a rectangle using formulas :
  - Circle : `area = pi * radius * radius`
  - Square : `area = side * side`
  - Rectangle : `area = length * width`
  - Triangle : `area = 0.5 * base * height`
- You can run the program with `go run calculate_area.go`
- You can run the test with `go test -run '^TestCalcShapeArea$' go-exercise`

### Exercise 2 :

- You can find the solution in the file `contact_list.go`
- Create a contact list with the following fields :
  - First name
  - Last name
  - Phone number
  - Email
- Manipulate the contact list with the following functions :
  - Add a contact
  - Remove a contact
  - Update a contact
  - Search a contact
  - List all contacts
- You can run the program with `go run contact_list.go`
- You can run the test with `go test -run '^(TestContactList_AddContact|TestContactList_GetContact|TestContactList_UpdateContact|TestContactList_DeleteContact)$' go-exercise`
