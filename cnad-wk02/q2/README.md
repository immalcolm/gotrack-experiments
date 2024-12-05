# Morsecode 
Write a Go program that prompts the user to key in a number, convert each digit according to the diagram below. Each Morse code of a digit has to be separated with 3 spaces. The following shows a sample output of the program.

```
Enter a number to convert:9
The morse code for 9 is ----.
>>> 
Enter a number to convert:101
The morse code for 101 is .----   -----   .----
>>>
```

## Thinking process
- explore using maps
- consider invalid inputs by user
- readup on exists 
- consider to display the total number of conversions


## Morse Code 0-9
- "0": "-----"
- "1": ".----"
- "2": "..---"
- "3": "...--"
- "4": "....-"
- "5": "....."
- "6": "-...."
- "7": "--..."
- "8": "---.."
- "9": "----."