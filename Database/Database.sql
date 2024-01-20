DROP DATABASE IF EXISTS IKEMON;

CREATE DATABASE IKEMON DEFAULT CHARACTER SET utf8 COLLATE utf8_hungarian_ci;

USE IKEMON;

CREATE TABLE TYPES (
    `TypeId` INT AUTO_INCREMENT,
    `TypeName` VARCHAR(255) NOT NULL,
    PRIMARY KEY(`TypeId`)
);

CREATE TABLE USERS (
    `UserId` VARCHAR(255) NOT NULL,
    `Username` VARCHAR(255) NOT NULL,
    `Password` VARCHAR(255) NOT NULL,
    `Email` VARCHAR(255) NOT NULL,
    `Money` DOUBLE,
    `Admin` BOOLEAN,
    `UserCreated` DATETIME DEFAULT NOW(),
    PRIMARY KEY (`UserId`)
);

CREATE TABLE CARDS (
    `CardId` VARCHAR(255) NOT NULL,
    `CardName` VARCHAR(255) NOT NULL,
    `TypeId` INT,
    `Hp` INT,
    `Attack` INT,
    `Defense` INT,
    `Price` DOUBLE,
    `Description` VARCHAR(255) NOT NULL,
    `Image` VARCHAR(255) NOT NULL,
    `OwnerId` VARCHAR(255) NOT NULL,
    `IsLocked` BOOLEAN DEFAULT FALSE,
    PRIMARY KEY(`CardId`)
);

CREATE TABLE SESSIONS (
    `SessionId` VARCHAR(255) NOT NULL,
    `UserId` VARCHAR(255) DEFAULT "",
    `Page` INT DEFAULT 0,
    `TypeId` VARCHAR(255) DEFAULT 0,
    `Created` DATETIME DEFAULT NOW(),
    PRIMARY KEY (`SessionId`)
);

INSERT INTO TYPES (`TypeName`) VALUES ('water');
INSERT INTO TYPES (`TypeName`) VALUES ('electric');
INSERT INTO TYPES (`TypeName`) VALUES ('grass');
INSERT INTO TYPES (`TypeName`) VALUES ('ice');
INSERT INTO TYPES (`TypeName`) VALUES ('fighting');
INSERT INTO TYPES (`TypeName`) VALUES ('poison');
INSERT INTO TYPES (`TypeName`) VALUES ('ground');
INSERT INTO TYPES (`TypeName`) VALUES ('psychic');
INSERT INTO TYPES (`TypeName`) VALUES ('bug');
INSERT INTO TYPES (`TypeName`) VALUES ('rock');
INSERT INTO TYPES (`TypeName`) VALUES ('ghost');
INSERT INTO TYPES (`TypeName`) VALUES ('dark');
INSERT INTO TYPES (`TypeName`) VALUES ('steel');

INSERT INTO CARDS VALUES ('card0', 'Pikachu', 2, 60, 20, 20, 160, 'Pikachu that can generate powerful electricity have cheek sacs that are extra soft and super stretchy.', 'https:\/\/assets.pokemon.com\/assets\/cms2\/img\/pokedex\/full\/025.png', "admin id placeholder", FALSE)