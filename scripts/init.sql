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

CREATE INDEX ad_price_index ON ad (price);
CREATE INDEX ad_price_desc_index ON ad (price DESC);
CREATE INDEX ad_date_added_index ON ad (date_added);
CREATE INDEX ad_date_added_desc_index ON ad (date_added DESC);
