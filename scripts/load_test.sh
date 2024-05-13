#!/bin/bash

# Prepare to send requests with random 'n' values between 20 and 40
echo "" > vegeta_targets.txt
for i in {1..100}; do
    echo "GET http://localhost:8080/fib?n=$((RANDOM % 21 + 20))" >> vegeta_targets.txt
done

# Perform the attack and generate a report
cat vegeta_targets.txt | vegeta attack -duration=10s -rate=200 | vegeta report

