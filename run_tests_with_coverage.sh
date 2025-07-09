#!/bin/bash

# Parking Lot System - Test Coverage Runner
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

# Initialize variables for coverage calculation
TOTAL_STATEMENTS=0
TOTAL_COVERED=0
PACKAGE_COUNT=0

echo -e "\n${BLUE}Running tests with coverage for all packages...${NC}\n"

# Function to run tests for a package and extract coverage
run_package_tests() {
    local package_path="$1"
    local package_name=$(basename "$package_path")
    
    echo -e "${YELLOW}Testing package: $package_name${NC}"
    
    # Run tests with coverage
    if go test -coverprofile="$COVERAGE_DIR/${package_name}.out" -covermode=count "$package_path"; then
        echo -e "  ${GREEN}✓ Tests passed${NC}"
        
        # Extract coverage percentage
        local coverage=$(go tool cover -func="$COVERAGE_DIR/${package_name}.out" | grep total: | awk '{print $3}' | sed 's/%//')
        
        if [ -n "$coverage" ] && [ "$coverage" != "0.0" ]; then
            echo -e "  ${BLUE}Coverage: ${coverage}%${NC}"
            
            # Extract total statements and covered statements
            local total_func=$(go tool cover -func="$COVERAGE_DIR/${package_name}.out" | grep total: | awk '{print $2}' | sed 's/,//')
            local covered_func=$(go tool cover -func="$COVERAGE_DIR/${package_name}.out" | grep total: | awk '{print $1}' | sed 's/,//')
            
            if [ -n "$total_func" ] && [ -n "$covered_func" ]; then
                TOTAL_STATEMENTS=$((TOTAL_STATEMENTS + total_func))
                TOTAL_COVERED=$((TOTAL_COVERED + covered_func))
                PACKAGE_COUNT=$((PACKAGE_COUNT + 1))
            fi
        else
            echo -e "  ${YELLOW}No coverage data available${NC}"
        fi
    else
        echo -e "  ${RED}✗ Tests failed${NC}"
    fi
    
    echo ""
}

# Run tests for each package that has test files
echo "📦 Testing individual packages:"

# Test packages with known test files
run_package_tests "./pkg/lot"
run_package_tests "./pkg/attendant"
run_package_tests "./pkg/services"
run_package_tests "./pkg/stratergy"
run_package_tests "./pkg/vehicle"
run_package_tests "./pkg/utils"

# Run all tests together for overall coverage
echo -e "${BLUE}Running all tests together for overall coverage...${NC}\n"

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

# Calculate average coverage from individual packages
echo -e "\n${BLUE}Coverage Statistics:${NC}"
echo "======================"

if [ $PACKAGE_COUNT -gt 0 ]; then
    # Calculate average coverage
    AVERAGE_COVERAGE=$(echo "scale=2; $TOTAL_COVERED * 100 / $TOTAL_STATEMENTS" | bc -l 2>/dev/null || echo "0.00")
    
    echo -e "Total packages tested: ${PACKAGE_COUNT}"
    echo -e "Total statements: ${TOTAL_STATEMENTS}"
    echo -e "Total covered statements: ${TOTAL_COVERED}"
    echo -e "Average statement coverage: ${AVERAGE_COVERAGE}%"
    
    # Color code the average coverage
    if (( $(echo "$AVERAGE_COVERAGE >= 80" | bc -l) )); then
        echo -e "${GREEN}✓ Excellent coverage!${NC}"
    elif (( $(echo "$AVERAGE_COVERAGE >= 60" | bc -l) )); then
        echo -e "${YELLOW}⚠ Good coverage, but room for improvement${NC}"
    else
        echo -e "${RED}⚠ Low coverage - consider adding more tests${NC}"
    fi
else
    echo -e "${YELLOW}No coverage data available from individual packages${NC}"
fi

echo -e "\n${BLUE}Coverage Reports:${NC}"
echo "=================="
echo -e "HTML Report: ${GREEN}$COVERAGE_DIR/coverage.html${NC}"
echo -e "Individual package reports: ${GREEN}$COVERAGE_DIR/*.out${NC}"

# Show detailed coverage by function
echo -e "\n${BLUE}Detailed Coverage by Function:${NC}"
echo "================================"
go tool cover -func="$COVERAGE_DIR/all.out"

echo -e "\n${GREEN}🎉 Test coverage analysis completed!${NC}"
echo -e "${BLUE}Open $COVERAGE_DIR/coverage.html in your browser to view detailed coverage.${NC}" 