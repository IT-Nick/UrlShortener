CREATE TABLE shortener
(
    id       serial,
    short    VARCHAR(100),
    original VARCHAR(200)
);

CREATE INDEX idx_shortener_short_original ON shortener(short, original); 
CREATE INDEX idx_shortener_original ON shortener(original); 
