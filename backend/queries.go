package main

var (
	topTitlesQuery = `SELECT *
FROM
    title_basics
JOIN
    title_ratings ON title_basics.tconst = title_ratings.tconst
WHERE
    title_basics.titleType = '%s' AND
    title_ratings.numVotes >= %d
ORDER BY
    title_ratings.averageRating DESC
LIMIT %d;`
)
