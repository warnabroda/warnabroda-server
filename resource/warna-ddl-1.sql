ALTER TABLE warnabroda.messages ADD COLUMN Active TINYINT(1) NOT NULL DEFAULT 1 AFTER Lang_key;
ALTER TABLE warnabroda.subjects ADD COLUMN Active TINYINT(1) NOT NULL DEFAULT 1 AFTER Lang_key;
ALTER TABLE warnabroda.contact_types ADD COLUMN Active TINYINT(1) NOT NULL DEFAULT 1 AFTER Lang_key;

UPDATE warnabroda.messages SET Active='0' WHERE Id='78';