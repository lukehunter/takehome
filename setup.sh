#!/bin/bash 
 
ARGS=3
 
if [ $# -ne "$ARGS" ] 
then 
    echo "you passed $# parameters" 
    echo "Usage: ./setup.sh dbname dbuser dbpass" 
    exit 
fi 
 
db_name=$1 
db_user=$2 
db_pass=$3 
 
#run mysql db setup script with paramenters 
 
#mysql –uroot -p -e "set @db_name=${db_name}; set @db_user=${db_user}; set @db_pass=${db_pass}; source initdb.sql;";
mysql –uroot -p -e "SET @db_name = ${db_name}; source initdb.sql;";
exit

# end of script.