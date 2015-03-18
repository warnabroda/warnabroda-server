ALTER TABLE warnabroda.messages ADD COLUMN Active TINYINT(1) NOT NULL DEFAULT 1 AFTER Lang_key;
ALTER TABLE warnabroda.subjects ADD COLUMN Active TINYINT(1) NOT NULL DEFAULT 1 AFTER Lang_key;
ALTER TABLE warnabroda.contact_types ADD COLUMN Active TINYINT(1) NOT NULL DEFAULT 1 AFTER Lang_key;

UPDATE warnabroda.messages SET Active='0' WHERE Id='78';

ALTER TABLE warnabroda.messages CHANGE COLUMN Id Id INT(11) NOT NULL ,
ADD COLUMN Last_modified_by INT(11) NULL AFTER Active,
ADD COLUMN Last_modified_date TIMESTAMP NULL AFTER Last_modified_by;

UPDATE warnabroda.messages SET Last_modified_by='2';
UPDATE warnabroda.messages SET Last_modified_date=now();

ALTER TABLE warnabroda.messages CHANGE COLUMN Id Id INT(11) NOT NULL AUTO_INCREMENT ;
