#!/bin/bash
# Modified version of bash script written by Saad Ismail - me@saadismail.net
# https://raw.githubusercontent.com/saadismail/useful-bash-scripts/master/db.sh
    echo "Please enter the NAME of the new database (example: database1)"
    read dbname
    echo "Please enter the database CHARACTER SET! (example: latin1, utf8, ...)"
    read charset
    echo "Creating new database, enter mysql root password when prompted..."
    mysql -uroot -p -e "CREATE DATABASE ${dbname} /*\!40100 DEFAULT CHARACTER SET ${charset} */;"
    echo "Database successfully created!"
    echo "Showing existing databases, enter mysql root password when prompted..."
    mysql -uroot -p -e "show databases;"
    echo ""
    echo "Please enter the NAME of the new database user! (example: user1)"
    read username
    echo "Creating new user..."
    echo "Please enter the PASSWORD for the new database user! Enter mysql root password when prompted"
    read userpass
    mysql -uroot -p -e "CREATE USER ${username}@localhost IDENTIFIED BY '${userpass}';"
    echo "User successfully created!"
    echo ""
    echo "Granting ALL privileges on ${dbname} to ${username}! Enter mysql root password when prompted"
    mysql -uroot -p -e "GRANT ALL PRIVILEGES ON ${dbname}.* TO '${username}'@'localhost';"
    mysql -uroot -p -e "FLUSH PRIVILEGES;"
    echo "You're good now :)"
    exit