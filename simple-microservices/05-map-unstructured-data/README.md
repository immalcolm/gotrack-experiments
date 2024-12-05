# Map Unstructured Data
`var result map[string] interface{}`
Means that result is a variable of type map, with key of type string, and value of type interface (which can be of any type)

**The statement:**
`currRates.(map[string] interface{})`
asserts the currRates variable into a map type with keys of type string


`SGD := result["rates"].(map[string]interface{})["SGD"]` go
**Variable Declaration and Initialization:**
The variable SGD is declared and initialized using the `:=`short variable declaration syntax. This syntax both declares the variable and assigns it a value in one step.

**Accessing Nested Map:**
The expression `result["rates"`] accesses the value associated with the key "rates" in the result map. The result variable is expected to be a map with string keys and values of any type (map[string]interface{}).

**Type Assertion:** 
The expression `result["rates"].(map[string]interface{})` performs a type assertion. 
It asserts that the value associated with the key `"rates"` is of type `map[string]interface{}`. 
Type assertion is used to convert an interface value to a more specific type. If the assertion fails (i.e., the value is not of the expected type), a runtime panic will occur.

**Accessing Inner Map:** 
After the type assertion, the expression `result["rates"].(map[string]interface{})["SGD"]` accesses the value associated with the key `"SGD"` in the nested map. This value is expected to be the exchange rate for the Singapore Dollar (SGD).

```
+-------------------------------------+
| Start of main()                     |
+-------------------------------------+
                |
                v
+-------------------------------------+
| Define jsonString with JSON data    |
+-------------------------------------+
                |
                v
+-------------------------------------+
| Declare result as map[string]interface{} |
+-------------------------------------+
                |
                v
+-------------------------------------+
| Unmarshal jsonString into result    |
| using json.Unmarshal                |
+-------------------------------------+
                |
                v
+-------------------------------------+
| Access result["rates"]              |
| (type assertion to map[string]interface{}) |
+-------------------------------------+
                |
                v
+-------------------------------------+
| Access nested map value             |
| result["rates"]["SGD"]              |
+-------------------------------------+
                |
                v
+-------------------------------------+
| Assign value to SGD variable        |
+-------------------------------------+
                |
                v
+-------------------------------------+
| Print exchange rate for SGD         |
| using fmt.Printf                    |
+-------------------------------------+
                |
                v
+-------------------------------------+
| End of main()                       |
+-------------------------------------+
```