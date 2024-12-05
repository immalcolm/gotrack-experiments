# 
messing around with arrays

## 

Arrays
- fixed in size
- can't easily sort them during runtime
- `var <array name> [<sizeof array>]<array type>= `

Slices
- dynamic sizing, grow/shrink
- reference to underlying array
- built in ops 
- `cap` (capacity of the slice in memory)
    - usually doubled to reduce frequence reallocation
- `len` increases with each added element.
- `copy` copy(<new slice>, <original slice >)

**Tips**
- When appending large number of elements, pre-allocate a large slice using `make` to avoid frequent reallocations

**Packages**
- `"sort"`
`sort.Ints`
`sort.Strings`


| Operation	| Length (len)	| Capacity (cap) |	Action |
|-----------|---------------|----------------|---------|
|s := make([]int, 0, 2)	|0	|2	|Slice created|
|append(s, 1)	| 1	| 2	| Add 1 element, within capacity|
|append(s, 2)	| 2 | 2	| Add 2nd element, still within capacity|
|append(s, 3)	| 3	| 4	| Exceeded capacity, new array allocated|
|append(s, 4)	| 4	| 4	| Add 4th element, within new capacity|
