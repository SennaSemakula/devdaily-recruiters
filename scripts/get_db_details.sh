#!/bin/bash

DB_NAME=$1

if [ -z $DB_NAME ]
then
    echo "Please provide database name!"
    exit 1
fi

echo "Listing databases..."
echo "@@____________________@@"
psql -d postgres -c "\l"

echo "Listing current users"
echo "@@____________________@@"
psql -d "postgres" -c "\du"

echo "Listing current schemas"
echo "@@____________________@@"
psql -d "postgres" -c "\dn" 