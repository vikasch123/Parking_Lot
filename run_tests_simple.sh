#!/bin/bash

# Parking Lot System - Simple Test Coverage Runner
# This script runs all tests and provides statement coverage statistics

set -e  # Exit on any error

echo "🚗 Parking Lot System - Test Coverage Runner"
echo "=============================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Create coverage directory
COVERAGE_DIR="coverage"
mkdir -p "$COVERAGE_DIR"

echo -e "\n${BLUE}Running all tests with coverage...${NC}\n"

# Run all tests with coverage
if go test -coverprofile="$COVERAGE_DIR/all.out" -covermode=count ./...; then
    echo -e "${GREEN}✓ All tests passed${NC}"
    
    # Get overall coverage
    OVERALL_COVERAGE=$(go tool cover -func="$COVERAGE_DIR/all.out" | grep total: | awk '{print $3}' | sed 's/%//')
    echo -e "${BLUE}Overall coverage: ${OVERALL_COVERAGE}%${NC}"
    
    # Generate HTML coverage report
    go tool cover -html="$COVERAGE_DIR/all.out" -o "$COVERAGE_DIR/coverage.html"
    echo -e "${GREEN}✓ HTML coverage report generated: $COVERAGE_DIR/coverage.html${NC}"
else
    echo -e "${RED}✗ Some tests failed${NC}"
    exit 1
fi

# Calculate average coverage using awk
echo -e "\n${BLUE}Coverage Statistics:${NC}"
echo "======================"

# Extract coverage data and calculate average
COVERAGE_DATA=$(go tool cover -func="$COVERAGE_DIR/all.out" | grep total: | awk '{
    total_statements += $2
    total_covered += $1
    count++
}
END {
    if (count > 0) {
        avg_coverage = (total_covered * 100) / total_statements
        printf "Total packages: %d\n", count
        printf "Total statements: %d\n", total_statements
        printf "Total covered statements: %d\n", total_covered
        printf "Average coverage: %.2f%%\n", avg_coverage
    }
}')

echo "$COVERAGE_DATA"

# Show detailed coverage by function
echo -e "\n${BLUE}Detailed Coverage by Function:${NC}"
echo "================================"
go tool cover -func="$COVERAGE_DIR/all.out"

echo -e "\n${BLUE}Coverage Reports:${NC}"
echo "=================="
echo -e "HTML Report: ${GREEN}$COVERAGE_DIR/coverage.html${NC}"
echo -e "Raw coverage data: ${GREEN}$COVERAGE_DIR/all.out${NC}"

echo -e "\n${GREEN}🎉 Test coverage analysis completed!${NC}"
echo -e "${BLUE}Open $COVERAGE_DIR/coverage.html in your browser to view detailed coverage.${NC}" 