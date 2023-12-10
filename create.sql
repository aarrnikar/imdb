CREATE TABLE name_basics (
    nconst VARCHAR(255) NOT NULL PRIMARY KEY,
    primaryName VARCHAR(255),
    birthYear INT,
    deathYear INT,
    primaryProfession VARCHAR(255),
    knownForTitles VARCHAR(255)
);

CREATE TABLE title_akas (
    titleId VARCHAR(255) NOT NULL,
    ordering INT NOT NULL,
    title VARCHAR(255),
    region VARCHAR(255),
    language VARCHAR(255),
    types VARCHAR(255),
    attributes VARCHAR(255),
    isOriginalTitle BOOLEAN,
    PRIMARY KEY(titleId, ordering)
);

CREATE TABLE title_basics (
    tconst VARCHAR(255) NOT NULL PRIMARY KEY,
    titleType VARCHAR(255),
    primaryTitle VARCHAR(255),
    originalTitle VARCHAR(255),
    isAdult BOOLEAN,
    startYear INT,
    endYear INT,
    runtimeMinutes INT,
    genres VARCHAR(255)
);

CREATE TABLE title_crew (
    tconst VARCHAR(255) NOT NULL PRIMARY KEY,
    directors VARCHAR(255),
    writers VARCHAR(255)
);

CREATE TABLE title_episode (
    tconst VARCHAR(255) NOT NULL PRIMARY KEY,
    parentTconst VARCHAR(255),
    seasonNumber INT,
    episodeNumber INT
);

CREATE TABLE title_principals (
    tconst VARCHAR(255) NOT NULL,
    ordering INT NOT NULL,
    nconst VARCHAR(255),
    category VARCHAR(255),
    job VARCHAR(255),
    characters VARCHAR(255),
    PRIMARY KEY(tconst, ordering)
);

CREATE TABLE title_ratings (
    tconst VARCHAR(255) NOT NULL PRIMARY KEY,
    averageRating DECIMAL(3, 1),
    numVotes INT
);