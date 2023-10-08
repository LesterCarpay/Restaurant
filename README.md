A small project to implement the back-end of a simple restaurant management app. 
It mainly serves for practice of Go and databases.

A [tutorial](https://blog.logrocket.com/building-simple-app-go-postgresql/) about how to implement a simple to-do list formed the basis of the project.
The complexity was increased by focusing on restaurants rather than a to-do list and interface was changed from a web page to the terminal.

To do:
- manage menu items should just be one menu, with options add menu item, delete menu item, change menu item description, add ingredients to menu item etc.
- removing menu items should also remove all their ingredient relations from the table
- all operations should be made safe in terms of ordering (i.e. use guaranteed ordering) by returning dictionaries with IDs and name/value etc instead of just the names
- password should be hidden
- each menu should have the ability to cancel the operation
- print the ingredient list in a readable format
