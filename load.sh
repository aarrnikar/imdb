#!/bin/bash

for file in dataset/*.tsv
do
    table_name=$(basename "$file" .tsv)
    echo "LOAD DATA LOCAL INFILE '$PWD/$file' INTO TABLE $table_name"
    echo "FIELDS TERMINATED BY '\t'"
    echo "LINES TERMINATED BY '\n'"
    echo "IGNORE 1 ROWS;"
    echo ""
done