#!/bin/sh

# ANSI escape codes for terminal output
YELLOW='\033[1;33m'
RED='\033[0;31m'
NO_COLOR='\033[0m'  # Reset text color to default

# Load environment variables from .env file
if [ -f .env ]; then
  echo "${YELLOW}.env file found, using existing file...${NO_COLOR}"
  export $(grep -v '^#' .env | xargs)
else
  if [ -f .env.example ]; then
    echo "${YELLOW}.env file not found, creating from .env.example...${NO_COLOR}"
    cp .env.example .env
    export $(grep -v '^#' .env | xargs)
  else
    echo "${RED}.env.example file not found. Cannot create .env file. Exiting...${NO_COLOR}"
    exit 1
  fi
fi

# Check if MySQL client is installed (assuming you are using Homebrew)
if ! command -v mysql &> /dev/null; then
  echo "${YELLOW}MySQL client not found, installing...${NO_COLOR}"
  brew install mysql
fi

# Dump MySQL database
echo "${YELLOW}Dumping MySQL database to ./database/init.sql...${NO_COLOR}"
# mysqldump -u $DB_USER -p $DB_PASSWORD -h $DB_HOST $DB_NAME > ./database/init.sql 2>&1
mysqldump -uroot -proot -h127.0.0.1  --protocol=tcp --databases  --skip-comments golang > ./database/init.sql 
# Check for errors during the dump
if [ $? -ne 0 ]; then
  echo "${RED}Failed to dump data. Exiting...${NO_COLOR}"
  exit 1
fi

# Modify init.sql (adjust as needed)
echo "${YELLOW}Modifying init.sql...${NO_COLOR}"
# ...

echo "Database dump completed successfully."
echo "init.sql modified."
