- slice is a reference to a contiguous segment of an underlying array 
- it has a length (no. of elements it contains) and a capacity (no. of elements in the underlying array, starting from the first element in the slice)
- slices are dynamic, meaning their size can change during runtime

creating a slice:
we can create a slice in several ways.
1. Using a Slice Literal
2. Using the make Function
3. Slicing an Array


Slice Length and Capacity:
1. Length: Number of elements in the slice (len(slice)).
2. Capacity: Maximum number of elements the slice can hold without resizing (cap(slice)).


Practice Exercise
Write a program that:

1. Creates a slice of integers.
2. Appends new elements to the slice.
3. Removes an element from the slice.
4. Prints the final slice.