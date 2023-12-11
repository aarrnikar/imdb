package main

import "database/sql"

type NameBasics struct {
	Nconst            string `db:"nconst" json:",omitempty"`
	PrimaryName       string `db:"primaryName" json:",omitempty"`
	BirthYear         int    `db:"birthYear" json:",omitempty" json:",omitempty"`
	DeathYear         int    `db:"deathYear" json:",omitempty"`
	PrimaryProfession string `db:"primaryProfession" json:",omitempty"`
	KnownForTitles    string `db:"knownForTitles" json:",omitempty"`
}

type TitleAkas struct {
	TitleID         string `db:"titleId" json:",omitempty"`
	Ordering        int    `db:"ordering" json:",omitempty"`
	Title           string `db:"title" json:",omitempty"`
	Region          string `db:"region" json:",omitempty"`
	Language        string `db:"language" json:",omitempty"`
	Types           string `db:"types" json:",omitempty"`
	Attributes      string `db:"attributes" json:",omitempty"`
	IsOriginalTitle bool   `db:"isOriginalTitle" json:",omitempty"`
}

type TitleBasics struct {
	Tconst         sql.NullString `db:"tconst" json:",omitempty"`
	TitleType      sql.NullString `db:"titleType" json:",omitempty"`
	PrimaryTitle   sql.NullString `db:"primaryTitle" json:",omitempty"`
	OriginalTitle  sql.NullString `db:"originalTitle" json:",omitempty"`
	IsAdult        sql.NullBool   `db:"isAdult" json:",omitempty"`
	StartYear      sql.NullInt32  `db:"startYear" json:",omitempty"`
	EndYear        sql.NullInt32  `db:"endYear" json:",omitempty"`
	RuntimeMinutes sql.NullInt32  `db:"runtimeMinutes" json:",omitempty"`
	Genres         sql.NullString `db:"genres" json:",omitempty"`
}

type TitleCrew struct {
	Tconst    string `db:"tconst" json:",omitempty"`
	Directors string `db:"directors" json:",omitempty"`
	Writers   string `db:"writers" json:",omitempty"`
}

type TitleEpisode struct {
	Tconst        string `db:"tconst" json:",omitempty"`
	ParentTconst  string `db:"parentTconst" json:",omitempty"`
	SeasonNumber  int    `db:"seasonNumber" json:",omitempty"`
	EpisodeNumber int    `db:"episodeNumber" json:",omitempty"`
}

type TitlePrincipals struct {
	Tconst     string `db:"tconst" json:",omitempty"`
	Ordering   int    `db:"ordering" json:",omitempty"`
	Nconst     string `db:"nconst" json:",omitempty"`
	Category   string `db:"category" json:",omitempty"`
	Job        string `db:"job" json:",omitempty"`
	Characters string `db:"characters" json:",omitempty"`
}

type TitleRatings struct {
	Tconst        sql.NullString  `db:"tconst" json:",omitempty"`
	AverageRating sql.NullFloat64 `db:"averageRating" json:",omitempty"`
	NumVotes      sql.NullInt32   `db:"numVotes" json:",omitempty"`
}

type Title struct {
	TitleBasics
	TitleRatings
}
