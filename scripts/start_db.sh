#!/bin/bash
#set -euo pipefail
DB_PATH="/usr/local/var/postgres"
LOG_FILE="myDB.log"
# check if data dir exists
if [ ! -d $DB_PATH ]
then
	echo "${DB_PATH} does not exist. Initialising now..."
	sudo mkdir $DB_PATH
        sudo chmod 775 $DB_PATH
        sudo chown $(whoami) $DB_PATH
        initdb /usr/local/var/postgres
fi
# Start postgres in daemon mode
echo "Starting postgres server..."
pg_ctl -D "${DB_PATH}" -l $LOG_FILE start




