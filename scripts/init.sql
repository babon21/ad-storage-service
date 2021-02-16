--
-- Table structure for table `ad`
--

DROP TABLE IF EXISTS ad;

CREATE TABLE ad
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(200)  NOT NULL,
    description VARCHAR(1000) NOT NULL,
    price       INTEGER       NOT NULL,
    main_photo  TEXT          NOT NULL,
    photos      TEXT[] NOT NULL,
    date_added  TIMESTAMP     NOT NULL
);
