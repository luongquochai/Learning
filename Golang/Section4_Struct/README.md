NOTE in section 4:
- **&variable**: Give me the memory address of the value this variable is poiting at
- ***pointer**: Give me the value this memory address is pointing at
- Big gotcha:

- **Value Types**: int, float, bool, string, struct
- **Reference Types**: slices, map, channels, pointer, functions

*Value types: Use pointers to change these things in a function*
*Reference types: Dont worry about pointers with these*

**NOTE**
*So remember with our struct when we passed it off to the function,*
*go internally, made **a copy the entire struct***
*And so when we modified it in the function, we were **modifying the copy**,*
***not the original struct***

*When we make a slice in go, go will automatically create that ***little slice data structure that has the length, capacity and a pointer to the array in one address in memory***. And then that data structure is going to have a pointer over to the actual underlying array, which it will be existing at a completely separate address in memory.*

Example:
 - mySlice := []string{"Hi","There","how","are","you?"}

**SLICE**                    
| ptr to head   |---------->|"Hi"|"There"|"how"|"are"|"you?"|**array**
| capacity      |
| length        |

**Example:**
-------------
- When we call a function and pass my slice into it, go is still behaving as a pass by value language. So Go is still making a copy of that underlying of that value. So when we call the update slice function and pass in our slice, we take the slice data structure and copy it off to another address in memory.
- Even though the slice data structure is copied, it is *still pointing at the original array in memory*.
**THIS IS REFERENCE TYPE**
-------------