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

INSERT INTO TYPES (`TypeName`) VALUES ('normal');
INSERT INTO TYPES (`TypeName`) VALUES ('fire');
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

INSERT INTO CARDS VALUES ('card0', 'Pikachu', 4, 60, 20, 20, 160, 'Pikachu that can generate powerful electricity have cheek sacs that are extra soft and super stretchy.', 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/025.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('card1', 'Charizard', 2, 78, 84, 78, 534, 'It spits fire that is hot enough to melt boulders. It may cause forest fires by blowing flames.', 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/006.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('card2', 'Bulbasaur', 5, 45, 49, 49, 318, 'A strange seed was planted on its back at birth. The plant sprouts and grows with this POKéMON.', 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/001.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('card3', 'Squirtle', 3, 44, 48, 65, 314, 'After birth, its back swells and hardens into a shell. Powerfully sprays foam from its mouth.', 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/007.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('card4', 'Caterpie', 11, 45, 30, 35, 195, 'Its short feet are tipped with suction pads that enable it to tirelessly climb slopes and walls.', 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/010.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('card5', 'Weedle', 11, 40, 35, 30, 195, 'Often found in forests, eating leaves. It has a sharp venomous stinger on its head.', 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/013.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('card6', 'Pidgey', 1, 40, 45, 40, 251, 'A common sight in forests and woods. It flaps its wings at ground level to kick up blinding sand.', 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/016.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('card7', 'Rattata', 1, 30, 56, 35, 253, 'Bites anything when it attacks. Small and very quick, it is a common sight in many places.', 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/019.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('card8', 'Spearow', 1, 40, 60, 30, 262, 'Eats bugs in grassy areas. It has to flap its short wings at high speed to stay airborne.', 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/021.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('card9', 'Ekans', 8, 35, 60, 44, 288, 'Moves silently and stealthily. Eats the eggs of birds, such as PIDGEY and SPEAROW, whole.', 'https://assets.pokemon.com/assets/cms2/img/pokedex/full/023.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('659ae5daad418', 'Captcha kód', 5, 100, 100, 100, 1000, 'The second fastest thing in the universe.', 'img/c.png', 'admin', FALSE);
INSERT INTO CARDS VALUES ('659aebd5c06b1', 'Segmentation Fault', 14, 200, 200, 200, 1500, 'I fear no man, but that thing. It scares me.', 'img/segmentation.png', 'admin', FALSE);