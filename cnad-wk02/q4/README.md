The following lists are shown in Python code, convert each of them to Map in Go

Rainfall in millimetres and the total days of rain for the month in Singapore for the last three years is: 

yr2017 = [['January',220,18], ['February',10,2], ['March',20,3], ['April',10,3], 
  ['May',22,6], ['June',24,4], ['July',168,13], ['August',198,15],
  ['September',188,15], ['October',213,18], ['November',260,19], 
  ['December',278,19]] 

yr2018 = [['January',245,18], ['February',41,6], ['March',201,13], ['April',197,16], 
['May',175,15], ['June',172,14], ['July',166,14], ['August',198,18], ['September',185,15], ['October',216,16], ['November',256,19], ['December',278,20]] 

yr2019 = [['January',250,17], ['February',40,7], ['March',200,12], ['April',195,15], 
['May',175,15], ['June',170,13], ['July',165,13], ['August',190,15], ['September',180,14], ['October',210,16], ['November',250,19], ['December',275,19]] 


The following list contains the above lists for the yearly rainfall:

rainfall_list = [yr2017, yr2018, yr2019]

The following list contains total number of days of rainfall for each year:

rainfall_year_list = [['2017', 0], ['2018', 0], ['2019', 0]]

The following list contains total number of days of rainfall for each year:

rainfall_lowest_list = [['2017', ' '], ['2018', ' '], ['2019', ' ']]
 
Write a Go program to do the following:

(a)	update the total numbers of days of rainfall for each year in rainfall_year_list and display the list

(b)	update year and month in which there was lowest rainfall for each of the three years in rainfall_lowest_list and display the list	

You are required to make use of the above rainfall_list in your program. 

The following shows the required output from the program:
