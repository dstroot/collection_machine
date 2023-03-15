---------------------------------------------
-- Table created to contain unfunded records
-- for collection.
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='COLLECTION_ITEM')
BEGIN
  DROP TABLE dbo.COLLECTION_ITEM;
END;

CREATE TABLE dbo.COLLECTION_ITEM
(
  COLLECTION_ITEM_ID      BIGINT IDENTITY (1,1)  NOT NULL, -- SK/PK
  REFUND_TRANSACTION_ID   BIGINT             NOT NULL ,  -- RAL/Refund transaction ID
  PROVIDER_ORDER_ID       VARCHAR(56)        NOT NULL ,
  --INTUIT_ORDER_ID         VARCHAR(16)        NOT NULL ,  -- optional for Intuit (EFE#)
  FIRST_NAME              VARCHAR(56)        NOT NULL ,
  MIDDLE_INITIAL          VARCHAR(1)         NULL ,
  LAST_NAME               VARCHAR(56)        NOT NULL ,
  ADDRESS1                VARCHAR(56)        NOT NULL ,
  ADDRESS2                VARCHAR(56)        NULL ,
  CITY                    VARCHAR(56)        NOT NULL , -- City/locality
  STATE                   VARCHAR(2)         NOT NULL ,  -- State/province/territory
  ZIPCODE                 VARCHAR(5)         NOT NULL ,  -- More international
  ZIPCODE_PLUS            VARCHAR(4)         NULL ,
  COUNTRY                 VARCHAR(56)        NULL ,
  EMAIL                   VARCHAR(56)        NOT NULL ,
  MOBILE_PHONE            VARCHAR(15)        NULL ,
  AMOUNT_OWED             NUMERIC(8,2)       NOT NULL ,  -- Limited by ACH field size to 10
  PAYMENT_URL             VARCHAR(150)       NULL ,
  RTN                     VARCHAR(9)         NOT NULL ,  -- must be 9 digits
  DAN                     VARCHAR(17)        NOT NULL , -- up to 17 characters
  HASH                    VARCHAR(60)        NULL ,
  TOKEN_EXPIRES           BIGINT             NULL , -- was DATETIME, now LINUX format
  PAID_IND                BIT                NOT NULL DEFAULT 0 ,
  CREATION_DATE           DATETIME           NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE             DATETIME           NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  CREATED_BY              VARCHAR(56)        NULL DEFAULT SUSER_SNAME() ,
  MODIFIED_BY             VARCHAR(56)        NULL DEFAULT SUSER_SNAME() ,
  TAX_AMOUNT_OWED         NUMERIC(8,2)       NULL ,
  PROCESS_TYPE_ID         INT                NOT NULL, -- REFERENCES PROCESS,  // TODO IS THIS NEEDED?
  GIACT_CODE              INT                NULL
);

ALTER TABLE dbo.COLLECTION_ITEM
  ADD CONSTRAINT [PK_COLLECTION_ITEM] PRIMARY KEY  CLUSTERED ([COLLECTION_ITEM_ID] ASC);


CREATE NONCLUSTERED INDEX [IDX_REFUND_TRANSACTION_ID] ON [COLLECTION_ITEM]
(
  [REFUND_TRANSACTION_ID]         ASC
);


CREATE NONCLUSTERED INDEX [IDX_PROVIDER_ORDER_ID] ON [COLLECTION_ITEM]
(
  [PROVIDER_ORDER_ID]         ASC
);


CREATE NONCLUSTERED INDEX [IDX_RTN] ON [COLLECTION_ITEM]
(
  [RTN]         ASC
);


CREATE NONCLUSTERED INDEX [IDX_CREATION_DATE] ON [COLLECTION_ITEM]
(
  [CREATION_DATE]         ASC
);

INSERT INTO dbo.[COLLECTION_ITEM] (REFUND_TRANSACTION_ID,PROVIDER_ORDER_ID,FIRST_NAME,MIDDLE_INITIAL,LAST_NAME,ADDRESS1,CITY,STATE,ZIPCODE,EMAIL,MOBILE_PHONE,AMOUNT_OWED,RTN,DAN,TAX_AMOUNT_OWED,PROCESS_TYPE_ID) VALUES
  (1000,'EFE432XQW14JDR0J','Dan','J','stroot','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',49.95,'122105278','0000000001', NULL,10),
  (1001,'EFE48CV3C1GEAOG1','Daniel','X','Stroot','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',39.95,'122105278','0000000001', 2.85,10),
  (1002,'EFE432XQX14JDR59','Bill','M','Madison','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',39.95,'122105278','0000000006',0,10),
  (1003,'EFE432XQZ14JDOWF','Dusan','H','Franklin','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',39.95,'122105278','0000000004',NULL,10),
  (1004,'EFE432XR114JDP5L','Harold','G','Nixon','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',89.95,'122105278','8037675686',2.95,20),
  (1005,'EFE432S7314JDH13','Phil','L','Clinton','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',29.95,'122105278','8037675686',1.1,20),
  (1006,'53003220170140024141','Slammy','D','Clown','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',89.95,'122105278','0000000007',2.95,30),
  (1007,'53003220170140024142','Phil','R','Up','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',29.95,'122105278','8037675686',1.1,40),
  (1008,'EFE432S7314JDH13','Phil','R','Up','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',29.95,'122105278','0000000008',1.1,10),
  (1009,'EFE432S7314JDH13','Phil','R','Up','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',29.95,'122105278','0000000012',1.1,10),
  (10010,'EFE432S7314JDH13','Phil','R','Up','1223 Main Street','New Brunswick','CA','92101','dan.stroot@gmail.com','19494634044',29.95,'122105278','0000000013',1.1,10);


-- GIACT Test accounts
-- ---------------------------------------------------------------------------------
--   {"122105278", "0000000001", 19, "GN01", true},
--   {"122105278", "0000000002", 20, "GN05", false},
--   {"122105278", "0000000003", 5, "GP01", true},
--   {"122105278", "0000000004", 1, "GS01", false},
--   {"122105278", "0000000005", 2, "GS02", false},
--   {"122105278", "0000000006", 3, "GS03", true},
--   {"122105278", "0000000007", 4, "GS04", true},
--   {"122105278", "0000000008", 21, "ND00", true},
--   {"122105278", "0000000009", 22, "ND01", false},
--   {"122105278", "0000000010", 6, "RT00", false},
--   {"122105278", "0000000011", 7, "RT01", false},
--   {"122105278", "0000000012", 8, "RT02", false},
--   {"122105278", "0000000013", 9, "RT03", true},
--   {"122105278", "0000000014", 10, "RT04", false},
--   {"122105278", "0000000015", 11, "RT05", false},
--   {"122105278", "0000000016", 12, "1111", true},
--   {"122105278", "0000000017", 13, "2222", true},
--   {"122105278", "0000000018", 14, "3333", true},
--   {"122105278", "0000000019", 15, "5555", true},
--   {"122105278", "0000000020", 16, "7777", true},
--   {"122105278", "0000000021", 17, "8888", true},
--   {"122105278", "0000000022", 18, "9999", true},

---------------------------------------------
-- Table created to contain Intuit Turbo Tax
-- detail records for the Intuit Turbo Tax
-- collections process known as "Turbo Tax
-- Payments", "TTP", and "Autodebit".
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='INTUIT_COLLECTION_DETAIL')
BEGIN
  DROP TABLE dbo.INTUIT_COLLECTION_DETAIL;
END;

CREATE TABLE INTUIT_COLLECTION_DETAIL
(
  INTUIT_COLLECTION_DETAIL_ID   BIGINT IDENTITY (1,1)  NOT NULL, -- SK/PK
  PROVIDER_ORDER_ID             VARCHAR(56)         NOT NULL ,
  --INTUIT_ORDER_ID               VARCHAR(16)         NOT NULL ,
  INTUIT_ORDER_DATE             DATETIME            NOT NULL ,
  INTUIT_ORDER_QTY              INT                 NOT NULL ,
  INTUIT_DOLLAR_AMOUNT          NUMERIC(8,2)        NOT NULL ,  -- Limited by ACH field size to 10
  INTUIT_ITEM_DESC              VARCHAR(256)        NOT NULL ,
  CREATION_DATE                 DATETIME            NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE                   DATETIME            NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CREATED_BY                    VARCHAR(56)         NULL DEFAULT SUSER_SNAME(),
  MODIFIED_BY                   VARCHAR(56)         NULL DEFAULT SUSER_SNAME()
);

ALTER TABLE dbo.INTUIT_COLLECTION_DETAIL
  ADD CONSTRAINT [PK_INTUIT_COLLECTION_DETAIL_ID] PRIMARY KEY  CLUSTERED ([INTUIT_COLLECTION_DETAIL_ID] ASC);

-- Creating a unique index instead in order to FK this column to collection_item_stg
/*
CREATE NONCLUSTERED INDEX [IDX_INTUIT_ORDER_ID] ON [INTUIT_COLLECTION_DETAIL]
(
  [INTUIT_ORDER_ID]         ASC
);
*/

INSERT INTO dbo.[INTUIT_COLLECTION_DETAIL] (PROVIDER_ORDER_ID, INTUIT_ORDER_DATE, INTUIT_ORDER_QTY, INTUIT_DOLLAR_AMOUNT, INTUIT_ITEM_DESC) VALUES
  ('EFE432XQW14JDR0J','1/21/2016 0:00',1,54.99,'TURBOTAX ONLINE DELUXE TY2014 PREP'),
  ('EFE432XQX14JDONH','1/21/2016 0:00',1,34.99,'TURBOTAX ONLINE TY2014 BUNDLE'),
  ('EFE48CV3C1GEAOG1','1/21/2016 0:00',1,34.99,'TURBOTAX ONLINE DELUXE TY2015 PREP'),
  ('EFE48CV3C1GEAOG1','1/21/2016 0:00',1,36.99,'TURBOTAX ONLINE STATE AL TY2015 PREP'),
  ('EFE48CV3C1GEAOG1','1/21/2016 0:00',1,36.99,'TURBOTAX ONLINE STATE GA TY2015 PREP'),
  ('EFE48CV3C1GEAOG1','1/21/2016 0:00',1,36.99,'TURBOTAX ONLINE STATE LA TY2015 PREP'),
  ('EFE48CV3C1GEAOG1','1/21/2016 0:00',1,36.99,'TURBOTAX ONLINE STATE NJ TY2015 PREP'),
  ('EFE48CV3C1GEAOG1','1/21/2016 0:00',1,36.99,'TURBOTAX ONLINE STATE SC TY2015 PREP'),
  ('EFE432XQX14JDR59','1/21/2016 0:00',1,44.99,'TURBOTAX MAX BUNDLE'),
  ('EFE432XQZ14JDOWF','1/21/2016 0:00',1,36.99,'TURBOTAX ONLINE STATE GA TY2015 PREP'),
  ('EFE432XQZ14JDOWF','1/21/2016 0:00',1,36.99,'TURBOTAX ONLINE STATE LA TY2015 PREP'),
  ('EFE432XR114JDP5L','1/21/2016 0:00',1,36.99,'TURBOTAX ONLINE STATE NJ TY2015 PREP'),
  ('EFE432S7314JDH13','1/21/2016 0:00',1,36.99,'TURBOTAX ONLINE STATE SC TY2015 PREP'),
  ('EFE432S7314JDH13','1/21/2016 0:00',1,44.99,'TURBOTAX MAX BUNDLE');
---------------------------------------------
-- Table created to store fees and their original values
-- before collections. After collections the fees
-- will be zeroed out.
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='COLLECTION_ITEM_FEE')
BEGIN
  DROP TABLE dbo.COLLECTION_ITEM_FEE;
END;


---------------------------------------------
-- Fee type table associated with COLLECTION_ITEM_FEE table
-- Will need to initially contain PSQL values until
-- SQL fee types are built out with perm architecture
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='COLLECTION_FEE_TYPE')
BEGIN
  DROP TABLE dbo.COLLECTION_FEE_TYPE;
END;

CREATE TABLE dbo.COLLECTION_FEE_TYPE
(
  COLLECTION_FEE_TYPE_ID    INT             NOT NULL, --SK/PK
  PERVASIVE_FEE_NAME        VARCHAR(56)     NOT NULL,
  ACTIVE_START_DATE         DATETIME        NULL ,
  ACTIVE_END_DATE           DATETIME        NULL ,
  CREATION_DATE             DATETIME        NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE               DATETIME        NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CREATED_BY                VARCHAR(56)     NULL DEFAULT SUSER_SNAME(),
  MODIFIED_BY               VARCHAR(56)     NULL DEFAULT SUSER_SNAME()
);

INSERT INTO dbo.COLLECTION_FEE_TYPE (COLLECTION_FEE_TYPE_ID,PERVASIVE_FEE_NAME, ACTIVE_START_DATE)
VALUES
('1','rsiteeffee',CURRENT_TIMESTAMP),
('2','rsiteprepfee',CURRENT_TIMESTAMP),
('3','rstatetaxonprod',CURRENT_TIMESTAMP),
('4','rstatetaxonpos',CURRENT_TIMESTAMP),
('5','raccthandlingfee',CURRENT_TIMESTAMP),
('6','rretailaccthandlingfee',CURRENT_TIMESTAMP),
('7','rtraninctfee',CURRENT_TIMESTAMP),
('8','raccthandlingfee',CURRENT_TIMESTAMP),
('9','rretailaccthandlingfee',CURRENT_TIMESTAMP);


-------------------------------------------
-- Create lookup table of retry types for collection process
------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='RETRY_TYPE')
BEGIN
  DROP TABLE dbo.RETRY_TYPE;
END;

CREATE TABLE dbo.RETRY_TYPE
(
  RETRY_TYPE_ID         INT           NOT NULL ,
  RETRY_TYPE_DESC       VARCHAR(256)  NULL ,
  ACTIVE_START_DATE     DATETIME      NULL ,
  ACTIVE_END_DATE       DATETIME      NULL ,
  CREATION_DATE         DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE           DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CREATED_BY            VARCHAR(56)   NULL DEFAULT SUSER_SNAME(),
  MODIFIED_BY           VARCHAR(56)   NULL DEFAULT SUSER_SNAME()
);

ALTER TABLE dbo.RETRY_TYPE
  ADD CONSTRAINT [PK_RETRY_TYPE_ID] PRIMARY KEY  CLUSTERED ([RETRY_TYPE_ID] ASC);

-- Insert values
INSERT INTO dbo.RETRY_TYPE (RETRY_TYPE_ID, RETRY_TYPE_DESC, ACTIVE_START_DATE) VALUES
(0,'Not a Retry Activity',CURRENT_TIMESTAMP),
(1,'Auto Debit Retry',CURRENT_TIMESTAMP),
(2,'Customer Initiated Payment Retry',CURRENT_TIMESTAMP);


---------------------------------------------
-- Table created to define a process. A process is
-- a sequence of activities, and a duration (in minutes)
-- between them. It also defines if an email or
-- sms is to be sent, and if so, what information
-- and templates to use. Finally it signals which
-- activities are the "end" of the process.
-- Processes can have multiple "end"s if the process
-- is restarted due to an ACH debit failure for example.
-- Some activities are immediately an "end" such
-- as the person making a payment or funding coming
-- in.
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='PROCESS_DEFINITION')
BEGIN
  DROP TABLE dbo.PROCESS_DEFINITION;
END;

CREATE TABLE dbo.PROCESS_DEFINITION
(
  PROCESS_DEFINITION_ID   INT IDENTITY(1,1) NOT NULL, --SK/PK
  PROCESS_TYPE_ID         INT           NOT NULL , -- REFERENCES PROCESS,
  ACTIVITY_TYPE_ID        INT           NOT NULL , --REFERENCES ACTIVITY,
  DURATION_MINUTES        BIGINT        NOT NULL ,
  DEBIT_IND               BIT           NOT NULL ,
  DIRECT_MAIL_IND         BIT           NOT NULL ,
  DEBIT_RETRY_IND         BIT           NOT NULL ,
  CUSTOMER_INITIATED_RETRY_IND BIT      NOT NULL DEFAULT 0,
  PROCESS_END             BIT           NOT NULL ,
  EMAIL_FROM_NAME         VARCHAR(56)   NULL ,
  EMAIL_FROM_EMAIL        VARCHAR(56)   NULL ,
  EMAIL_SUBJECT           VARCHAR(100)  NULL ,
  EMAIL_TEMPLATE          VARCHAR(25)   NULL ,
  SMS_TEMPLATE            VARCHAR(25)   NULL ,
  ACTIVE_START_DATE       DATETIME      NULL ,
  ACTIVE_END_DATE         DATETIME      NULL ,
  CREATION_DATE           DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE             DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CREATED_BY              VARCHAR(56)   NULL DEFAULT SUSER_SNAME(),
  MODIFIED_BY             VARCHAR(56)   NULL DEFAULT SUSER_SNAME()
);

 ALTER TABLE dbo.PROCESS_DEFINITION
  ADD CONSTRAINT [PK_PROCESS_DEFINITION_ID] PRIMARY KEY  CLUSTERED ([PROCESS_DEFINITION_ID] ASC);

CREATE NONCLUSTERED INDEX [IDX_PROCESS_TYPE_ID] ON [PROCESS_DEFINITION]
(
  [PROCESS_TYPE_ID]         ASC
);

CREATE NONCLUSTERED INDEX [IDX_ACTIVITY_TYPE_ID] ON [PROCESS_DEFINITION]
(
  [ACTIVITY_TYPE_ID]         ASC
);

-- Insert values

INSERT INTO dbo.PROCESS_DEFINITION (PROCESS_TYPE_ID, ACTIVITY_TYPE_ID, DURATION_MINUTES, DEBIT_IND, DIRECT_MAIL_IND,
DEBIT_RETRY_IND, CUSTOMER_INITIATED_RETRY_IND, PROCESS_END, EMAIL_FROM_NAME, EMAIL_FROM_EMAIL, EMAIL_SUBJECT, EMAIL_TEMPLATE, SMS_TEMPLATE, ACTIVE_START_DATE)
VALUES
  (10,0,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),    -- Start Collections process
  (10,1,3,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Your TurboTax Payment is Due','1',NULL,CURRENT_TIMESTAMP),
  (10,2,3,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Reminder: Your TurboTax Payment is Due','2',NULL,CURRENT_TIMESTAMP),
  (10,3,3,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Important Notice of Action: Bank Account Debit','3','sms',CURRENT_TIMESTAMP),
  (10,4,1,1,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),    -- Autodebit step
  (10,5,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),    -- Mark autodebit failue PERMANENT
  (10,6,0,0,0,0,0,1,'Intuit','dan.stroot@sbtpg.com','Important Request for Action: Balance Past Due','5',NULL,CURRENT_TIMESTAMP),
  (10,7,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),    -- Mark autodebit failue RETRY
  (10,8,0,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Important Notice of Action: Bank Account Retry Pending','4',NULL,CURRENT_TIMESTAMP),
  (10,9,0,0,0,1,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),    -- Perform autodebit failue RETRY
  (10,10,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- mark second attempt failed
  (10,11,0,0,0,0,0,1,'Intuit','dan.stroot@sbtpg.com','Important Request for Action: Balance Past Due','5',NULL,CURRENT_TIMESTAMP),
  (10,21,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- CC Payment
  (10,22,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Auth Bank Payment
  (10,23,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Mark Bank Payment failue PERMANENT
  (10,24,0,0,0,0,0,1,'Intuit','dan.stroot@sbtpg.com','Important Request for Action: Balance Past Due','5',NULL,CURRENT_TIMESTAMP),
  (10,25,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Mark Bank Payment failue RETRY
  (10,26,0,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Important Notice of Action: Bank Account Retry Pending','4',NULL,CURRENT_TIMESTAMP),
  (10,27,0,0,0,0,1,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Perform Bank Payment failue RETRY
  (10,28,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- mark second attempt failed
  (10,29,0,0,0,0,0,1,'Intuit','dan.stroot@sbtpg.com','Important Request for Action: Balance Past Due','5',NULL,CURRENT_TIMESTAMP),

  -- Deferral for email 1
  (10,30,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),  -- Autodebit deferred (email 1)
  (10,31,4,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Autodebit Deferral Completed','8',NULL,CURRENT_TIMESTAMP),
  (10,32,1,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Autodebit Deferral Completed','9',NULL,CURRENT_TIMESTAMP),
  (10,33,0,1,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Autodebit step

  -- Deferral for email 2
  (10,35,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),  -- Autodebit deferred (email 2)
  (10,36,4,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Autodebit Deferral Completed','8',NULL,CURRENT_TIMESTAMP),
  (10,37,1,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Autodebit Deferral Completed','9',NULL,CURRENT_TIMESTAMP),
  (10,38,0,1,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Autodebit step

  -- Deferral for email 3
  (10,40,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),  -- Autodebit deferred (email 3)
  (10,41,4,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Autodebit Deferral Completed','8',NULL,CURRENT_TIMESTAMP),
  (10,42,1,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Autodebit Deferral Completed','9',NULL,CURRENT_TIMESTAMP),
  (10,43,0,1,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Autodebit step

  (10,70,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Stopped Autodebit (agent)
  (10,71,0,0,0,0,0,1,'Intuit','dan.stroot@sbtpg.com','Important Request for Action: Balance Past Due','5',NULL,CURRENT_TIMESTAMP),

  (10,75,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Stopped due to bad account (GIACT)
  (10,76,0,0,0,0,0,1,'Intuit','dan.stroot@sbtpg.com','Important Request for Action: Balance Past Due','5',NULL,CURRENT_TIMESTAMP),

  (10,79,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees Refunded
  (10,89,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees suppressed (Melissa)
  (10,90,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees waived (Agent)

  (10,99,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Refund Funded

  (20,0,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),    -- Start Collections process
  (20,1,3,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Your TurboTax Payment is Due','1',NULL,CURRENT_TIMESTAMP),
  (20,2,0,0,0,0,0,0,'Intuit','dan.stroot@sbtpg.com','Reminder: Your TurboTax Payment is Due','2',NULL,CURRENT_TIMESTAMP),
  (20,15,0,0,1,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Direct mail step
  (20,21,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- CC Payment
  (20,79,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees Refunded
  (20,89,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees suppressed
  (20,90,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees waived
  (20,99,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Refund Funded

  (30,100,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),    -- Start Collections process
  (30,101,2,0,0,0,0,0,'Taxslayer','support@e.taxslayer.com','Information about your tax refund','TAX_SLAYER_1',NULL,CURRENT_TIMESTAMP),
  (30,102,2,0,0,0,0,0,'Taxslayer','support@e.taxslayer.com','Information about your tax refund','TAX_SLAYER_2',NULL,CURRENT_TIMESTAMP),
  (30,103,2,0,0,0,0,0,'Taxslayer','support@e.taxslayer.com','Information about your tax refund','TAX_SLAYER_3',NULL,CURRENT_TIMESTAMP),
  (30,104,2,0,0,0,0,0,'Taxslayer','support@e.taxslayer.com','Information about your tax refund','TAX_SLAYER_4',NULL,CURRENT_TIMESTAMP),
  (30,105,1,1,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Autodebit step
  (30,110,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Mark autodebit failue PERMANENT
  (30,121,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- CC Payment
  (30,122,2,0,0,0,0,1,'Taxslayer','support@e.taxslayer.com','Thank You!','TAX_SLAYER_5',NULL,CURRENT_TIMESTAMP),
  (30,123,0,1,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Auth Bank Payment
  (30,124,2,0,0,0,0,1,'Taxslayer','support@e.taxslayer.com','Thank You!','TAX_SLAYER_5',NULL,CURRENT_TIMESTAMP),
  (30,125,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Mark Bank Payment failue PERMANENT
  (30,170,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Agent Stopped
  (30,175,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- GIACT Stop
  (30,179,0,1,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees Refunded
  (30,189,0,1,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees supressed
  (30,190,0,1,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees waived
  (30,199,0,1,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Refund Funded

  (40,100,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),    -- Start Collections process
  (40,101,2,0,0,0,0,0,'Taxslayer','support@e.taxslayer.com','Information about your tax refund','TAX_SLAYER_1',NULL,CURRENT_TIMESTAMP),
  (40,102,2,0,0,0,0,0,'Taxslayer','support@e.taxslayer.com','Information about your tax refund','TAX_SLAYER_2',NULL,CURRENT_TIMESTAMP),
  (40,103,2,0,0,0,0,0,'Taxslayer','support@e.taxslayer.com','Information about your tax refund','TAX_SLAYER_3',NULL,CURRENT_TIMESTAMP),
  (40,104,2,0,0,0,0,1,'Taxslayer','support@e.taxslayer.com','Information about your tax refund','TAX_SLAYER_4',NULL,CURRENT_TIMESTAMP),
  (40,121,0,0,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- CC Payment
  (40,122,2,0,0,0,0,1,'Taxslayer','support@e.taxslayer.com','Thank You!','TAX_SLAYER_5',NULL,CURRENT_TIMESTAMP),
  (40,123,0,1,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Auth Bank Payment
  (40,124,2,0,0,0,0,1,'Taxslayer','support@e.taxslayer.com','Thank You!','TAX_SLAYER_5',NULL,CURRENT_TIMESTAMP),
  (40,125,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Mark Bank Payment failue PERMANENT
  (40,170,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Agent Stopped
  (40,175,0,0,0,0,0,1,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- GIACT Stop
  (40,179,0,1,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees Refunded
  (40,189,0,1,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees supressed
  (40,190,0,1,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP),   -- Fees waived
  (40,199,0,1,0,0,0,0,NULL,NULL,NULL,NULL,NULL,CURRENT_TIMESTAMP);   -- Refund Funded

-------------------------------------------
-- Create lookup table of payment types for collection process
------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='PAYMENT_TYPE')  -- [DJS] Modified
BEGIN
  DROP TABLE dbo.PAYMENT_TYPE;
END;

CREATE TABLE dbo.PAYMENT_TYPE
(
  PAYMENT_TYPE_ID         INT           NOT NULL ,
  PAYMENT_TYPE_NAME       VARCHAR(56)   NULL ,
  PAYMENT_TYPE_DESC       VARCHAR(256)  NULL ,
  ACTIVE_START_DATE       DATETIME      NULL ,
  ACTIVE_END_DATE         DATETIME      NULL ,
  CREATION_DATE           DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE             DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CREATED_BY              VARCHAR(56)   NULL DEFAULT SUSER_SNAME(),
  MODIFIED_BY             VARCHAR(56)   NULL DEFAULT SUSER_SNAME()
);

ALTER TABLE dbo.PAYMENT_TYPE
  ADD CONSTRAINT [PK_PAYMENT_TYPE_ID] PRIMARY KEY  CLUSTERED ([PAYMENT_TYPE_ID] ASC);

-- Insert values

INSERT INTO dbo.PAYMENT_TYPE (PAYMENT_TYPE_ID, PAYMENT_TYPE_NAME, ACTIVE_START_DATE) VALUES
(1,'Bank Ach',CURRENT_TIMESTAMP),
(2,'Credit Card',CURRENT_TIMESTAMP);

---------------------------------------------
-- Table created to store payments
-- for each collection item.
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='COLLECTION_PAYMENT')
BEGIN
  DROP TABLE dbo.COLLECTION_PAYMENT;
END;

CREATE TABLE dbo.COLLECTION_PAYMENT
(
  COLLECTION_PAYMENT_ID   BIGINT IDENTITY (1,1) NOT NULL ,
  REFUND_TRANSACTION_ID   BIGINT          NOT NULL,
  PAYMENT_AMOUNT          NUMERIC(8,2)    NOT NULL ,  -- Limited by ACH field size to 10
  PAYMENT_DATE            DATETIME        NOT NULL ,
  PAYMENT_TYPE_ID         INT             NOT NULL ,   -- (1) Bank ACH, (2) Credit Card See payment type table
  PAYMENT_CC_JSON         VARCHAR(MAX)    NULL ,  -- JSON response from stripe
  CUSTOMER_INITIATED_IND  BIT             NOT NULL DEFAULT 0,  -- t/f, cust approved bank ACH?
  ACH_SENT_IND            BIT             NULL ,  -- t/f ACH processed
  CREATION_DATE           DATETIME        NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE             DATETIME        NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CREATED_BY              VARCHAR(56)     NULL DEFAULT SUSER_SNAME(),
  MODIFIED_BY             VARCHAR(56)     NULL DEFAULT SUSER_SNAME(),
  ACH_REJECTED_IND        BIT             NOT NULL DEFAULT 0,
  ACH_RETRY_IND           BIT             NOT NULL DEFAULT 0
);

ALTER TABLE dbo.COLLECTION_PAYMENT
  ADD CONSTRAINT [PK_COLLECTION_PAYMENT_ID] PRIMARY KEY  CLUSTERED ([COLLECTION_PAYMENT_ID] ASC);

CREATE NONCLUSTERED INDEX [IDX_REFUND_TRANSACTION_ID] ON [COLLECTION_PAYMENT]
(
  [REFUND_TRANSACTION_ID]         ASC
);


---------------------------------------------
-- Table created to reversals to payments initiated
-- for each collection item.
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='COLLECTION_REFUND')
BEGIN
  DROP TABLE dbo.COLLECTION_REFUND;
END;

CREATE TABLE dbo.COLLECTION_REFUND
(
  COLLECTION_REFUND_ID    BIGINT        IDENTITY (1,1) NOT NULL ,
  COLLECTION_PAYMENT_ID   BIGINT        NOT NULL ,
  REFUND_TRANSACTION_ID   BIGINT        NOT NULL,
  REFUND_AMOUNT           NUMERIC(8,2)  NOT NULL ,  -- Limited by ACH field size to 10
  REFUND_DATE             DATETIME      NOT NULL ,
  PROCESSED_IND           BIT           NOT NULL DEFAULT 0 , -- t/f ACH or credit card processed
  CREATION_DATE           DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE             DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CREATED_BY              VARCHAR(56)   NULL DEFAULT SUSER_SNAME(),
  MODIFIED_BY             VARCHAR(56)   NULL DEFAULT SUSER_SNAME()
);


-- ALTER TABLE dbo.COLLECTION_REFUND
--   ADD CONSTRAINT [PK_COLLECTION_REFUND_ID] PRIMARY KEY  CLUSTERED ([COLLECTION_REFUND_ID] ASC);
--
-- CREATE NONCLUSTERED INDEX [IDX_COLLECTION_PAYMENT_ID] ON [COLLECTION_REFUND]
-- (
--   [COLLECTION_PAYMENT_ID]         ASC
-- );
--
-- CREATE NONCLUSTERED INDEX [IDX_REFUND_TRANSACTION_ID] ON [COLLECTION_REFUND]
-- (
--   [REFUND_TRANSACTION_ID]         ASC
-- );



---------------------------------------------
-- Table created to store emails for delivery
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='EMAIL_QUEUE')
BEGIN
  DROP TABLE dbo.EMAIL_QUEUE;
END;

CREATE TABLE dbo.EMAIL_QUEUE
(
  EMAIL_QUEUE_ID          BIGINT IDENTITY (1,1) NOT NULL ,
  REFUND_TRANSACTION_ID   BIGINT        NOT NULL ,
  COLLECTION_ACTIVITY_ID  BIGINT        NULL,
  FIRST_NAME              VARCHAR(56)   NOT NULL ,
  MIDDLE_INITIAL          VARCHAR(1)    NULL ,
  LAST_NAME               VARCHAR(56)   NOT NULL ,
  EMAIL_ADDRESS           VARCHAR(56)   NOT NULL ,
  SUBJECT                 VARCHAR(56)   NULL ,
  AMOUNT                  NUMERIC(10,2) NULL ,
  URL                     VARCHAR(150)  NULL ,
  FROM_NAME               VARCHAR(56)   NOT NULL ,  -- Send from various names (support, sales, marketing, etc.)
  FROM_EMAIL_ADDRESS      VARCHAR(56)   NOT NULL ,  -- Send from various emails
  TEMPLATE                VARCHAR(56)   NOT NULL , -- Template to use
  SENT_IND                BIT           NOT NULL DEFAULT 0,  -- 1 (true) if sent/ 0 if not
  CREATION_DATE           DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE             DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CREATED_BY              VARCHAR(56)   NULL DEFAULT SUSER_SNAME(),
  MODIFIED_BY             VARCHAR(56)   NULL DEFAULT SUSER_SNAME(),
  QUEUE_VENDOR_ID         INT NULL,
  RESPONSE_MESSAGE        VARCHAR(MAX) NULL,
  ERROR_CODE              INT NULL,
  ERROR_CODE_DESC         VARCHAR(125) NULL
);

ALTER TABLE dbo.EMAIL_QUEUE
  ADD CONSTRAINT [PK_EMAIL_QUEUE_ID] PRIMARY KEY CLUSTERED ([EMAIL_QUEUE_ID] ASC);

CREATE NONCLUSTERED INDEX [IDX_REFUND_TRANSACTION_ID] ON [EMAIL_QUEUE]
(
  [REFUND_TRANSACTION_ID]         ASC
);

CREATE NONCLUSTERED INDEX [IDX_COLLECTION_ACTIVITY_ID] ON [EMAIL_QUEUE]
(
  [COLLECTION_ACTIVITY_ID]         ASC
);

---------------------------------------------
-- Table created to store SMS messages for delivery
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='SMS_QUEUE')
BEGIN
  DROP TABLE SMS_QUEUE;
END;

CREATE TABLE SMS_QUEUE
(
  SMS_QUEUE_ID            BIGINT IDENTITY (1,1) NOT NULL ,
  REFUND_TRANSACTION_ID   BIGINT      NOT NULL,
  COLLECTION_ACTIVITY_ID  BIGINT      NULL,
  FIRST_NAME              VARCHAR(56) NOT NULL ,
  MIDDLE_INITIAL          VARCHAR(1)  NULL ,
  LAST_NAME               VARCHAR(56) NOT NULL ,
  SMS_PHONE               VARCHAR(56) NOT NULL ,
  FROM_NAME               VARCHAR(56) NOT NULL ,  -- send from various names (support, sales, marketing, etc.)
  TEMPLATE                VARCHAR(56) NOT NULL , -- Template to use
  SENT_IND                BIT         NOT NULL DEFAULT 0,  -- 1 (true) if sent/ 0 if not
  MESSAGE_SID             VARCHAR(34) NULL, -- Twilio SID
  CREATION_DATE           DATETIME    NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE             DATETIME    NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  CREATED_BY              VARCHAR(56) NULL DEFAULT SUSER_SNAME(),
  MODIFIED_BY             VARCHAR(56) NULL DEFAULT SUSER_SNAME()
);

ALTER TABLE dbo.SMS_QUEUE
  ADD CONSTRAINT [PK_SMS_QUEUE_ID] PRIMARY KEY CLUSTERED ([SMS_QUEUE_ID] ASC);

CREATE NONCLUSTERED INDEX [IDX_REFUND_TRANSACTION_ID] ON [SMS_QUEUE]
(
  [REFUND_TRANSACTION_ID]         ASC
);

---------------------------------------------
-- Table created to store direct mail messages for delivery
-- Join to COLLECTION_ITEM for name/address/etc.
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='DIRECT_MAIL_QUEUE')
BEGIN
  DROP TABLE dbo.DIRECT_MAIL_QUEUE;
END;

CREATE TABLE dbo.DIRECT_MAIL_QUEUE (
  DIRECT_MAIL_QUEUE_ID    BIGINT IDENTITY (1,1) NOT NULL ,
  REFUND_TRANSACTION_ID   BIGINT      NOT NULL ,
  DM_PROCESSED_IND        BIT         NULL DEFAULT 0 , -- t/f Processed to file
  CREATION_DATE           DATETIME    NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE             DATETIME    NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  CREATED_BY              VARCHAR(56) NULL DEFAULT SUSER_SNAME() ,
  MODIFIED_BY             VARCHAR(56) NULL DEFAULT SUSER_SNAME()
);

ALTER TABLE dbo.DIRECT_MAIL_QUEUE
  ADD CONSTRAINT [PK_DIRECT_MAIL_QUEUE_ID] PRIMARY KEY  CLUSTERED ([DIRECT_MAIL_QUEUE_ID] ASC);

CREATE NONCLUSTERED INDEX [IDX_REFUND_TRANSACTION_ID] ON [DIRECT_MAIL_QUEUE]
(
  [REFUND_TRANSACTION_ID]         ASC
);

---------------------------------------------
-- Table created to contain collections
-- activity for each collection item. Each
-- step will be tracked, logged, etc. in
-- this table. For example: first email sent,
-- second email sent, final email, etc.
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='COLLECTION_ACTIVITY')
BEGIN
  DROP TABLE dbo.COLLECTION_ACTIVITY;
END;

CREATE TABLE dbo.COLLECTION_ACTIVITY
(
  COLLECTION_ACTIVITY_ID    BIGINT IDENTITY(1,1) NOT NULL ,
  REFUND_TRANSACTION_ID     BIGINT NOT NULL ,
  PROCESS_TYPE_ID           INT NOT NULL , --REFERENCES ACTIVITY,
  ACTIVITY_TYPE_ID          INT NOT NULL , --REFERENCES ACTIVITY,
  TIMESTAMP_UTC             DATETIME  NOT NULL DEFAULT (SYSUTCDATETIME()) ,  -- SYSUTCDATETIME() needed by software to match servers
  CREATION_DATE             DATETIME  NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  CREATED_BY                VARCHAR(56) NULL DEFAULT SUSER_SNAME() ,
);

ALTER TABLE dbo.COLLECTION_ACTIVITY
  ADD CONSTRAINT [PK_COLLECTION_ACTIVITY_ID] PRIMARY KEY  CLUSTERED ([COLLECTION_ACTIVITY_ID] ASC);

CREATE NONCLUSTERED INDEX [IDX_REFUND_TRANSACTION_ID] ON [COLLECTION_ACTIVITY]
(
  [REFUND_TRANSACTION_ID] ASC
);

CREATE NONCLUSTERED INDEX [IDX_PROCESS_TYPE_ID] ON [COLLECTION_ACTIVITY]
(
  [PROCESS_TYPE_ID] ASC
);

CREATE NONCLUSTERED INDEX [IDX_ACTIVITY_TYPE_ID] ON [COLLECTION_ACTIVITY]
(
  [ACTIVITY_TYPE_ID] ASC
);

-- test data

INSERT INTO COLLECTION_ACTIVITY (REFUND_TRANSACTION_ID,PROCESS_TYPE_ID,ACTIVITY_TYPE_ID) VALUES
  (1000,10,0),
  (1001,10,0),
  (1002,10,0),
  (1003,10,0),
  (1004,20,0),
  (1005,20,0),
  (1006,30,100),  -- TAXSLAYER
  (1007,40,100);  -- TAXSLAYER
  -- (1008,10,0),  -- Bad GIACT
  -- (1009,10,0),  -- Bad GIACT
  -- (10010,10,0);  -- Bad GIACT
  -- (1003,10,2),
  -- (1003,10,3),
  -- (1004,10,0),
  -- (1004,10,1),
  -- (1004,10,2),
  -- (1004,10,3),
  -- (1004,10,4),
  -- (1005,10,0),
  -- (1005,10,1),
  -- (1005,10,2),
  -- (1005,10,3),
  -- (1005,10,4);

---------------------------------------------
-- Table created to contain process
-- descriptions.
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='PROCESS_TYPE')
BEGIN
  DROP TABLE dbo.PROCESS_TYPE;
END;

CREATE TABLE dbo.PROCESS_TYPE
(
  PROCESS_TYPE_ID       INT           NOT NULL ,  /* not IDENTITY since IDs will be chosen and consistent between environments */
  PROCESS_NAME          VARCHAR(56)   NULL ,
  PROCESS_DESC          VARCHAR(256)  NULL ,
  ACTIVE_START_DATE     DATETIME      NULL ,
  ACTIVE_END_DATE       DATETIME      NULL ,
  CREATION_DATE         DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE           DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  CREATED_BY            VARCHAR(56)   NULL DEFAULT SUSER_SNAME() ,
  MODIFIED_BY           VARCHAR(56)   NULL DEFAULT SUSER_SNAME()
 );

ALTER TABLE dbo.PROCESS_TYPE
  ADD CONSTRAINT [PK_PROCESS_TYPE_ID] PRIMARY KEY  CLUSTERED ([PROCESS_TYPE_ID] ASC);

-- Insert process types
INSERT INTO dbo.PROCESS_TYPE (PROCESS_TYPE_ID, PROCESS_NAME, PROCESS_DESC, ACTIVE_START_DATE, ACTIVE_END_DATE) VALUES
  (10,'INTUIT COLLECTIONS PROCESS','This process is used to recover fees due to unfunded RTs',CURRENT_TIMESTAMP, NULL),
  (20,'INTUIT BAD RTN COLLECTIONS PROCESS','This process is used to recover fees due to unfunded RTs where the routing number has a high number of unfunded transactions',CURRENT_TIMESTAMP, NULL),
  (30,'TAXSLAYER COLLECTIONS PROCESS','This process is used to recover fees due to unfunded RTs',CURRENT_TIMESTAMP, NULL),
  (40,'TAXSLAYER BAD RTN COLLECTIONS PROCESS','This process is used to recover fees due to unfunded RTs where the routing number has a high number of unfunded transactions',CURRENT_TIMESTAMP, NULL);

---------------------------------------------
-- Table created to contain activity
-- decriptions for processes.
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='ACTIVITY_TYPE')
BEGIN
  DROP TABLE dbo.ACTIVITY_TYPE;
END;

CREATE TABLE dbo.ACTIVITY_TYPE
(
  ACTIVITY_TYPE_ID    INT           NOT NULL ,   /* not IDENTITY since IDs will be chosen and consistent between environments */
  ACTIVITY_NAME       VARCHAR(56)   NOT NULL ,
  ACTIVITY_DESC       VARCHAR(256)  NULL ,
  ACTIVE_START_DATE   DATETIME      NULL ,
  ACTIVE_END_DATE     DATETIME      NULL ,
  CREATION_DATE       DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE         DATETIME      NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  CREATED_BY          VARCHAR(56)   NULL DEFAULT SUSER_SNAME() ,
  MODIFIED_BY         VARCHAR(56)   NULL DEFAULT SUSER_SNAME()
 );

ALTER TABLE dbo.ACTIVITY_TYPE
  ADD CONSTRAINT [PK_ACTIVITY_TYPE_ID] PRIMARY KEY CLUSTERED ([ACTIVITY_TYPE_ID] ASC);

-- values needed
INSERT INTO OLTP_SYS.dbo.ACTIVITY_TYPE (ACTIVITY_TYPE_ID, ACTIVITY_NAME, ACTIVITY_DESC, ACTIVE_START_DATE, ACTIVE_END_DATE) VALUES
  (0,'Collections Process Began','Start of collection processing',CURRENT_TIMESTAMP,NULL),
  (1,'First Email Sent','First email in process sent to the customer',CURRENT_TIMESTAMP,NULL),
  (2,'Second Email Sent','Second email in process sent to the customer',CURRENT_TIMESTAMP,NULL),
  (3,'Third Email Sent','Third email in process sent to the customer',CURRENT_TIMESTAMP,NULL),
  (4,'Account Autodebited','ACH forced on customer account for payment',CURRENT_TIMESTAMP,NULL),
  (5,'Autodebit Failed - Permanent','Attempted to recover fees from customer but the ach failed. Will not retry to collect through ACH.',CURRENT_TIMESTAMP,NULL),
  (6,'Final Email Sent - After Perm Failure','Final email sent to customers after the permanent autodebit fails',CURRENT_TIMESTAMP,NULL),
  (7,'Autodebit Failed - Retry','Attempted to recover fees from customer but the ach failed. Will retry to collect through ACH.',CURRENT_TIMESTAMP,NULL),
  (8,'Autodebit Email Retry Sent','Email associated with autodebit sent to customer.',CURRENT_TIMESTAMP,NULL),
  (9,'Autodebit Account Retried','ACH forced on customer account for payment. Not the first attempt.',CURRENT_TIMESTAMP,NULL),
  (10,'Autodebit Retry Failed','Second or later attempt at collecting payment failed.',CURRENT_TIMESTAMP,NULL),
  (11,'Final Email Sent','Final email sent to customer before TPG will attempt further collections like autodebit.',CURRENT_TIMESTAMP,NULL),
  (15,'Direct Mail Sent','Direct mail sent to customer. For Intuit this is done for customers on the bad RTN list.',CURRENT_TIMESTAMP,NULL),
  (21,'Credit Card Payment','Customer paid by credit card',CURRENT_TIMESTAMP,NULL),
  (22,'Bank Authorization Payment','Customer paid by ACH',CURRENT_TIMESTAMP,NULL),
  (23,'Bank Authorization Payment - ACH Failed Permanent','Customer-initiated ACH that failed and will not be retried',CURRENT_TIMESTAMP,NULL),
  (24,'Final Email Sent - ACH Failed Permanent','Email sent to customer after a Customer-initiated ACH that failed and will not be retried',CURRENT_TIMESTAMP,NULL),
  (25,'Bank Authorization Payment - Retry','Activity noting the ACH needed to be resubmitted for a failed custmer-initiated ACH',CURRENT_TIMESTAMP,NULL),
  (26,'Bank Authorization Payment Email Retry Sent','Email sent to customer when a custmer-initiated ACH is retried',CURRENT_TIMESTAMP,NULL),
  (27,'Bank Authorization Payment - Retried','Customer paid by ACH. The payment failed. Retried ACH.',CURRENT_TIMESTAMP,NULL),
  (28,'Bank Authorization Payment Retry Failed','Customer paid by ACH',CURRENT_TIMESTAMP,NULL),
  (29,'Final Email Sent After Second ACH Attempt','Email sent when the retried ach fails a second time',CURRENT_TIMESTAMP,NULL),

  (30,'Initiate Deferral Wait Period - Email 1','Start deferral period 1',CURRENT_TIMESTAMP,NULL),
  (31,'Send Deferral Email','Email sent to customer noting deferral',CURRENT_TIMESTAMP,NULL),
  (32,'Send Deferral Period End Email','Email sent to customer noting end of deferral',CURRENT_TIMESTAMP,NULL),
  (33,'Account Autodebited','ACH sent to customer account for payment',CURRENT_TIMESTAMP,NULL),

  (35,'Initiate Deferral Wait Period - Email 2','Start deferral period 2',CURRENT_TIMESTAMP,NULL),
  (36,'Send Deferral Email','Email sent to customer noting deferral',CURRENT_TIMESTAMP,NULL),
  (37,'Send Deferral Period End Email','Email sent to customer noting end of deferral',CURRENT_TIMESTAMP,NULL),
  (38,'Account Autodebited','ACH sent to customer account for payment',CURRENT_TIMESTAMP,NULL),

  (40,'Initiate Deferral Wait Period - Email 3','Start deferral period 3',CURRENT_TIMESTAMP,NULL),
  (41,'Send Deferral Email','Email sent to customer noting deferral',CURRENT_TIMESTAMP,NULL),
  (42,'Send Deferral Period End Email','Email sent to customer noting end of deferral',CURRENT_TIMESTAMP,NULL),
  (43,'Account Autodebited','ACH sent to customer account for payment',CURRENT_TIMESTAMP,NULL),

  (70,'Agent stopped','Autodebit canceled by Agent',CURRENT_TIMESTAMP,NULL),
  (71,'Final Email Sent','Final email sent to customers after autodebit fails or is canceled',CURRENT_TIMESTAMP,NULL),

  (75,'GIACT stopped','Autodebit canceled due to bad RTN/DAN',CURRENT_TIMESTAMP,NULL),
  (76,'Final Email Sent','Final email sent to customers after autodebit fails or is canceled',CURRENT_TIMESTAMP,NULL),

  (79,'Refunded','Customer was granted a refund of previously paid fees. This is typically initiated through customer service.',CURRENT_TIMESTAMP,NULL),

  (89,'Fees Suppressed','Customer was granted a waiver of fees and does not have to pay them.',CURRENT_TIMESTAMP,NULL),
  (90,'Fees Waived','Customer was granted a waiver of fees and does not have to pay them. This is initiated through customer service.',CURRENT_TIMESTAMP,NULL),

  (99,'Refund Funded - Stop Collection','Payment has been made by the customer. No further collections are required.',CURRENT_TIMESTAMP,NULL),

  (100, 'Start Collections Process', 'Start of collection processing.', CURRENT_TIMESTAMP, NULL),
  (101, 'First Email Sent', 'First email in process sent to the customer.', CURRENT_TIMESTAMP, NULL),
  (102, 'Second Email Sent', 'Second email in process sent to the customer.', CURRENT_TIMESTAMP, NULL),
  (103, 'Third Email Sent', 'Third email in process sent to the customer.', CURRENT_TIMESTAMP, NULL),
  (104, 'Fourth Email Sent', 'Fourth email in  process sent to the customer.', CURRENT_TIMESTAMP, NULL),
  (105, 'Account Autodebited', 'ACH sent to customer account for payment.', CURRENT_TIMESTAMP, NULL),
  (110, 'Autodebit Failed - Permanent', 'Attempted to recover fees from customer but the ACH failed.', CURRENT_TIMESTAMP, NULL),
  (121, 'Credit Card Payment', 'Customer paid by credit card', CURRENT_TIMESTAMP, NULL),
  (122, 'Thank you email', 'Thank you email sent to customer', CURRENT_TIMESTAMP, NULL),
  (123, 'Bank Authorization Payment', 'Customer paid by bank account.', CURRENT_TIMESTAMP, NULL),
  (124, 'Thank you email', 'Thank you email sent to customer', CURRENT_TIMESTAMP, NULL),
  (125, 'Bank Authorization Payment - ACH Failed Permanent', 'Customer-initiated ACH failed.', CURRENT_TIMESTAMP, NULL),
  (175, 'Stopped due to bad account (GIACT)', 'Autodebit canceled due to bad RTN/DAN', CURRENT_TIMESTAMP, NULL),
  (179, 'Refunded - CSR Initiated', 'Customer was granted a refund of previously paid fees.', CURRENT_TIMESTAMP, NULL),
  (190, 'Fees Waived', 'Customer was granted a waiver of fees through customer service.',CURRENT_TIMESTAMP,NULL),
  (199, 'Refund Funded - Stop Collection', 'Payment has been made by the customer. No further collections are required.', CURRENT_TIMESTAMP, NULL);


---------------------------------------------
-- Table created to hold bad RTN numbers.  This is used
-- for accounts that will recieve a direct mail instead of
-- an automatic debit. The reject rate on these is too
-- high to run automatic debit against.
-- NOTE: RTN should be indexed!
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='BAD_RTN')
BEGIN
  DROP TABLE dbo.BAD_RTN;
END;

CREATE TABLE dbo.BAD_RTN
(
  BAD_RTN_ID          BIGINT IDENTITY(1,1)  NOT NULL,
  RTN                 VARCHAR(9)      NOT NULL ,  -- Can start with 0: varchar
  APPLIED_PRODUCTS    BIGINT          NOT NULL,
  FUNDED_PRODUCTS     BIGINT          NOT NULL,
  FUNDING_PERCENTAGE  NUMERIC (4,2)   NOT NULL,
  BANK_NAME           VARCHAR(56)     NOT NULL,
  LOW_FUNDING_IND     BIT             NULL DEFAULT 0,  -- we track low funding rate RTNs
  PREPAID_IND         BIT             NULL DEFAULT 0,  -- we track "prepaid" RTNs
  ACTIVE_START_DATE   DATETIME        NULL ,
  ACTIVE_END_DATE     DATETIME        NULL ,
  CREATION_DATE       DATETIME        NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  MODIFY_DATE         DATETIME        NOT NULL DEFAULT (CURRENT_TIMESTAMP) ,
  CREATED_BY          VARCHAR(56)     NULL DEFAULT SUSER_SNAME(),
  MODIFIED_BY         VARCHAR(56)     NULL DEFAULT SUSER_SNAME()
);

ALTER TABLE dbo.BAD_RTN
  ADD CONSTRAINT [PK_dbo.BAD_RTN_ID] PRIMARY KEY CLUSTERED (BAD_RTN_ID ASC);

ALTER TABLE dbo.BAD_RTN
ADD  CONSTRAINT [RTN_UNIQUE] UNIQUE NONCLUSTERED
(
  [RTN] ASC
)
WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = ON, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY];

-- Bad RTN list
INSERT INTO dbo.BAD_RTN (RTN, APPLIED_PRODUCTS, FUNDED_PRODUCTS, FUNDING_PERCENTAGE, BANK_NAME) VALUES
  ('         ',0,0,0.00,'NO RTN AVAILABLE'),
  ('231374945',234,22,9.00,'SANTANDER BANK'),
  ('271970066',174,17,10.00,'CITIBANK F S B'),
  ('026012894',1012,109,11.00,'TD BANK NA'),
  ('021207167',477,56,12.00,'TD BANK NA'),
  ('211174123',347,45,13.00,'TD BANK N.A.'),
  ('121141822',153,23,15.00,'BANK OF AMERICA NA'),
  ('063112809',1131,211,19.00,'TD BANK N.A.'),
  ('053274074',1113,224,20.00,'PREPAID SUNTRUST'),
  ('028001489',203,41,20.00,'CITIBANK N.A'),
  ('041215663',225,50,22.00,'SUTTON BANK'),
  ('053201885',102,23,23.00,'TD BANK NA'),
  ('101089742',33262,10371,31.00,'H & R BLOCK BANK'),
  ('124085260',108,36,33.00,'GE CAPITAL BANK'),
  ('063192971',161,57,35.00,'REPUBLIC BANK & TRUST CO'),
  ('051000101',816,292,36.00,'BANK OF AMERICA, N.A'),
  ('063115194',4119,1481,36.00,'REGIONS BANK'),
  ('123306160',4428,1767,40.00,'US BANK NA'),
  ('021001486',638,261,41.00,'CITIBANK N.A.'),
  ('124303065',2683,1121,42.00,'BANK OF AMERICA NA'),
  ('031101169',134495,59697,44.00,'THE BANCORP.COM'),
  ('124303120',98012,45775,47.00,'GREEN DOT BANK'),
  ('122287675',2482,1182,48.00,'BOFI FEDERAL BANK'),
  ('121139313',3960,1888,48.00,'US BANK NA (LAMORINDA NB)'),
  ('021272626',227,112,49.00,'BANCO POPULAR, NORTH AMERICA'),
  ('071925130',448,230,51.00,'FIFTH THIRD BANK'),
  ('067004764',295,159,54.00,'CITIBANK /FLORIDA/BR'),
  ('101205681',8618,4788,56.00,'UMB, NA'),
  ('114994196',949,541,57.00,'FIRST BANK & TRUST'),
  ('322283699',482,299,62.00,'KINECTA FEDERAL CREDIT UNION'),
  ('062202985',419,269,64.00,'METROPOLITAN COMMERCIAL BANK'),
  ('124071889',19157,12686,66.00,'AMERICAN EXPRESS CENTURION BANK'),
  ('067014822',7966,5290,66.00,'TD BANK, NA'),
  ('322270288',516,348,67.00,'ONEWEST BANK N.A'),
  ('061120000',1437,976,68.00,'GREEN DOT BANK'),
  ('053111632',1880,1291,69.00,'BRANCH BANKING AND TRUST COMPANY'),
  ('071214579',335,233,70.00,'BANK OF AMERICA'),
  ('124085066',312,220,71.00,'AMERICAN EXPRESS BANK FSB'),
  ('073972181',133799,95453,71.00,'METABANK'),
  ('121000248',2731,1982,73.00,'WELLS FARGO BANK'),
  ('124303162',576,422,73.00,'GOBANK, A DIVISION OF GREEN DOT BANK'),
  ('021502011',690,507,73.00,'BANCO POPULAR'),
  ('071925046',458,338,74.00,'AMERICAN CHARTERED BANK'),
  ('026014928',276,204,74.00,'METROPOLITAN COMMERCIAL BANK'),
  ('122239869',257,193,75.00,'BANC OF CALIFORNIA'),
  ('124385119',331,251,76.00,'SALLIE MAE BANK'),
  ('266086554',1802,1373,76.00,'CITIBANK'),
  ('044204370',526,403,77.00,'OHIO VALLEY BANK CO'),
  ('021409169',47196,36436,77.00,'JP MORGAN CHASE BANK'),
  ('021912915',330,256,78.00,'TD BANK USA, N.A.'),
  ('021201503',715,557,78.00,'TD BANK NA'),
  ('114924742',617743,481646,78.00,'THE BANCORP BANK'),
  ('112201289',270,211,78.00,'CITIZENS BANK OF LAS CRUCES'),
  ('026009593',6141,4805,78.00,'BANK OF AMERICA N.A.'),
  ('111916656',314,248,79.00,'PEOPLES NATIONAL BANK'),
  ('031902766',2525,2024,80.00,'PNC BANK'),
  ('111017694',401,326,81.00,'BRANCH BANKING AND TRUST CO'),
  ('071025797',1514,1231,81.00,'REPUBLIC BANK OF CHICAGO'),
  ('065402892',1194,982,82.00,'REGIONS BANK'),
  ('122244184',3660,3019,82.00,'M B FINANCIAL BANK'),
  ('071923909',7736,6465,84.00,'FIFTH THIRD BANK'),
  ('054001725',3125,2613,84.00,'TD BANK, NA'),
  ('122106293',823,689,84.00,'SUNBANK NA'),
  ('271984861',290,243,84.00,'COMMUNITY TRUST CREDIT UNION'),
  ('263190812',882,740,84.00,'FIFTH THIRD BANK'),
  ('063192874',2144,1799,84.00,'URBAN TRUST BANK'),
  ('031100209',407,342,84.00,'CITIBANK NA'),
  ('263184815',570,479,84.00,'AXIOM BANK'),
  ('122106183',534,449,84.00,'TCF NATIONAL BANK'),
  ('071906641',598,505,84.00,'MB FINANCIAL BANK'),
  ('267084199',5521,4665,84.00,'UNKNOWN'),
  ('113024588',15255,12929,85.00,'WEX BANK'),
  ('052002166',1152,977,85.00,'CITIBANK MARYLAND, N.A'),
  ('122242597',2564,2177,85.00,'FIRST CITIZENS BANK'),
  ('101114659',292,248,85.00,'FIRST BANK KANSAS'),
  ('124085024',33420,28621,86.00,'GREEN DOT BANK');

---------------------------------------------
-- Table created to hold GIACT Codes
---------------------------------------------

IF EXISTS (SELECT * FROM SYS.TABLES WHERE NAME='GiactCode')
BEGIN
  DROP TABLE dbo.GiactCode;
END;

CREATE TABLE [dbo].[GiactCode](
	[GiactCodeId]      [int] NOT NULL,
	[GiactCodeName]    [varchar](50) NOT NULL,
	[GiactCodeDesc]    [varchar](255) NOT NULL,
	[DebitFlag]        [bit] NOT NULL,
	[ActiveStartDate]  [datetime] NOT NULL DEFAULT getdate(),
	[ActiveEndDate]    [datetime] NULL,
	[CreationDate]     [datetime] NULL DEFAULT getdate(),
	[CreatedBy]        [varchar](255) NULL DEFAULT suser_name(),
	[ModifyDate]       [datetime] NULL DEFAULT getdate(),
	[ModifyBy]         [varchar](255) NULL DEFAULT suser_name(),
 CONSTRAINT [pk_GiactCodeId] PRIMARY KEY CLUSTERED
(
	[GiactCodeId] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]

INSERT INTO dbo.GiactCode (GiactCodeId, GiactCodeName, GiactCodeDesc, DebitFlag) VALUES
  -- (0, 'Null', 'There is no AccountResponseCode value for this result.', 0),
  (1, 'GS01', 'Invalid Routing Number - The routing number supplied fails the validation test.', 0),
  (2, 'GS02', 'Invalid Account Number - The account number supplied fails the validation test.', 0),
  (3, 'GS03', 'Invalid Check Number - The check number supplied fails the validation test.', 1),
  (4, 'GS04', 'Invalid Amount - The amount supplied fails the validation test.', 1),
  (5, 'GP01', 'The account was found as active in your Private Bad Checks List.', 1),
  (6, 'RT00', 'The routing number belongs to a reporting bank; however, no positive nor negative information has been reported.', 0),
  (7, 'RT01', 'This account should be declined based on the risk factor being reported.', 0),
  (8, 'RT02', 'This item should be rejected based on the risk factor being reported.', 0),
  (9, 'RT03', 'Current negative data exists on this account. Accept transaction with risk.', 1),
  (10, 'RT04', 'Non-Demand Deposit Account, Credit Card Check, Line of Credit, Home Equity, or a Brokerage check.', 0),
  (11, 'RT05', 'N/A.', 0),
  (12, '1111', 'Account Verified – The account was found to be an open and valid checking account.', 1),
  (13, '2222', 'AMEX – The account was found to be an American Express Travelers Cheque account.', 1),
  (14, '3333', 'Non-Participant Provider – This account was reported with acceptable, positive data.', 1),
  (15, '5555', 'Savings Account Verified – The account was found to be an open and valid savings account.', 1),
  (16, '7777', 'The checking account was found to have a positive history.', 1),
  (17, '8888', 'The savings account was found to have a positive history.', 1),
  (18, '9999', 'This account was reported with acceptable, positive data found in recent transactions.', 1),
  (19, 'GN01', 'Negative information was found in this accounts history.', 1),
  (20, 'GN05', 'The routing number is reported as not currently assigned to a financial institution.', 0),
  (21, 'ND00', 'No positive or negative information has been reported on the account.', 1),
  (22, 'ND01', 'This routing number can only be valid for US Government financial institutions.', 0);
