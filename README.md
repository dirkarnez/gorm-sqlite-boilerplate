gorm-sqlite-boilerplate
=======================
This sample shows how to reset hardcoded rows when doing code-first migration - this program will drop the table and recreate table as defined in golang's struct.

### Browser
- [SQLiteDatabaseBrowserPortable_3.12.2_English.paf.exe](https://github.com/sqlitebrowser/sqlitebrowser/releases/download/v3.12.2/SQLiteDatabaseBrowserPortable_3.12.2_English.paf.exe)

### TODOs
- [ ] [dirkarnez/shopspring-decimal-playground](https://github.com/dirkarnez/shopspring-decimal-playground)
- [x] gorm gen

### Proofs
- [x] 1000 * 0.01 = 10.0 (0.01 is just divide by 100)
- [x] 10 - 10/3 - 10/3 = 3.34 (instead of 3.333333...)
- [x] 3.34 + 0.005 = 3.35 (0.005 is considered 0.01, which is true for money calculation)

### Notes
- In init stage, consider drop table if the table structure needs to be updated.
