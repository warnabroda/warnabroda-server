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

###Desenv da area de resposta 22-mar-2015

CREATE TABLE warning_resp (
  id bigint(20) NOT NULL AUTO_INCREMENT,
  id_warning bigint(20) DEFAULT NULL,
  resp_hash varchar(255) DEFAULT NULL,
  message text,
  reply_to varchar(45) DEFAULT NULL,
  ip varchar(255) DEFAULT NULL,
  browser varchar(255) DEFAULT NULL,
  operating_system varchar(255) DEFAULT NULL,
  device varchar(255) DEFAULT NULL,
  raw varchar(255) DEFAULT NULL,
  created_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  reply_date timestamp NULL DEFAULT NULL,
  response_read timestamp NULL DEFAULT NULL,
  lang_key varchar(10) DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY Id_UNIQUE (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
