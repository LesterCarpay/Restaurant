A small project to implement the back-end of a simple restaurant management app. 
It mainly serves for practice of Go and databases.

Currently, the program allows the user to add ingredients and menu items to a database through a terminal-based interface. Menu items have a description and contain ingredients that the user specified. Potential future features are:
- menu item categories
- generating a restaurant menu
- an order system, where the user can input and process a customer order
- a stock system where the program keeps track of the restaurant stock, the user can add stock and orders consume stock
- supplier information could be added to the database

The terminal interface was implemented through a simple custom dialogue loop. The database was implemented as a PostgreSQL relational database, with tables for ingredients, menu items and menu_item_ingredients, the latter keeping track of which menu items contained which ingredients. A [tutorial](https://blog.logrocket.com/building-simple-app-go-postgresql/) about how to implement a simple to-do list formed the basis for the implementation of the project. The complexity was increased by focusing on restaurants rather than a to-do list and interface was changed from a web page to the terminal.

To do:
- manage menu items should just be one menu, with options add menu item, delete menu item, change menu item description, add ingredients to menu item etc.
- removing ingredients should also remove all menu item ingredient relations from the table
- password should be hidden
- each menu should have the ability to cancel the operation
