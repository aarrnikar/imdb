#!/bin/bash

urls=(
    "https://datasets.imdbws.com/name.basics.tsv.gz"
    "https://datasets.imdbws.com/title.akas.tsv.gz"
    "https://datasets.imdbws.com/title.basics.tsv.gz"
    "https://datasets.imdbws.com/title.crew.tsv.gz"
    "https://datasets.imdbws.com/title.episode.tsv.gz"
    "https://datasets.imdbws.com/title.principals.tsv.gz"
    "https://datasets.imdbws.com/title.ratings.tsv.gz"
)

dataset_dir="dataset"
mkdir -p "$dataset_dir"

for url in "${urls[@]}"
do
    filename=$(basename "$url")
    tsv_filename="${filename%.gz}"

    # Download the file
    curl -O "$url"

    # Move the downloaded file to the dataset directory
    mv "$filename" "$dataset_dir"

    # Decompress the file in the dataset directory
    gunzip "$dataset_dir/$filename"
done