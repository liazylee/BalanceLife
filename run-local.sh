#!/bin/bash

# Color codes for terminal output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Display header
echo -e "${BLUE}==================================================${NC}"
echo -e "${BLUE}       BalanceLife API - Local Runner              ${NC}"
echo -e "${BLUE}==================================================${NC}"

# Configuration check
echo -e "${YELLOW}Checking configuration...${NC}"

# Check if using MongoDB X509 authentication
if grep -q "authMechanism=MONGODB-X509" config/.env; then
  CERT_PATH=$(grep "MONGODB_CERT_PATH" config/.env | cut -d "=" -f2)
  if [ ! -f "$CERT_PATH" ]; then
    echo -e "${RED}Error: MongoDB X509 certificate file not found at $CERT_PATH${NC}"
    echo -e "${YELLOW}Please ensure your certificate file is in the correct location${NC}"
    exit 1
  else
    echo -e "${GREEN}MongoDB X509 certificate found at $CERT_PATH${NC}"
  fi
elif grep -q "your-username:your-password" config/.env; then
  echo -e "${RED}Error: You need to update your MongoDB credentials in config/.env${NC}"
  exit 1
fi

if grep -q "your-redis-host.redislabs.com:10000" config/.env; then
  echo -e "${RED}Error: You need to update your Redis host in config/.env${NC}"
  exit 1
fi

if grep -q "your-redis-password" config/.env; then
  echo -e "${RED}Error: You need to update your Redis password in config/.env${NC}"
  exit 1
fi

echo -e "${GREEN}Configuration looks good!${NC}"

# Function to run the application
run_app() {
  echo -e "${YELLOW}Starting BalanceLife API with hot reloading...${NC}"
  if ! command -v air &> /dev/null; then
    echo -e "${YELLOW}Installing Air for hot reloading...${NC}"
    go install github.com/air-verse/air@latest
  fi
  air

}

# Function to test database connections
test_connections() {
  echo -e "${YELLOW}Testing database connections...${NC}"
  go run scripts/test_connections.go
  if [ $? -ne 0 ]; then
    echo -e "${RED}Connection tests failed. Please check the error messages above.${NC}"
    exit 1
  fi
}

# Main logic
case "$1" in
  run)
    run_app
    ;;
  
  test)
    test_connections
    ;;
  
  *)
    echo -e "${YELLOW}Usage:${NC}"
    echo -e "  ./run-local.sh [command]"
    echo
    echo -e "${YELLOW}Commands:${NC}"
    echo -e "  ${GREEN}run${NC}     - Run the application"
    echo -e "  ${GREEN}test${NC}    - Test database connections"
    echo
    echo -e "${YELLOW}Examples:${NC}"
    echo -e "  ./run-local.sh run     ${BLUE}# Run the application${NC}"
    echo -e "  ./run-local.sh test    ${BLUE}# Test database connections${NC}"
    ;;
esac 