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

# Initialize database (replace with your actual SQL script)
echo "${YELLOW}Initializing DB with ./database/init.sql...${NO_COLOR}"
mysql -umysql -pmysql -h127.0.0.1 --protocol=tcp < ./database/init.sql > mysql_output.txt 2>&1

# Check for errors in the initialization process
error_count=$(grep -c "ERROR" mysql_output.txt)
if [ $error_count -gt 0 ]; then
  echo "${RED}DB initialization encountered errors. Number of errors: $error_count${NO_COLOR}"
  echo "Error details:"
  grep "ERROR" mysql_output.txt
else
  echo "\n\n${YELLOW}DB initialization complete with NO errors.${NO_COLOR}"
fi

# Remove the output file
rm mysql_output.txt