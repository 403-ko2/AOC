# AOC
Advent through the years

## 2025
Language: Golang

The way I am formatting the files is to have one main .go file with the main package and function. Each file after that will be seperated by the day and its connected input in the following format: 
```
day<#>.go < day<#>_input.txt
```
This way I can follow go syntax protocols and I wont have to see the "errors" in my terminal editor due to having multiple main functions since originally I made a new main for each file to run individually. Now each file will just house its specific function and I will put all of the days in the main file. Obviously to run any input successfully you would have to comment out all other days' functions and just use the day specified for that input.
