package storedb

import (
  "database/sql"
)

const (
  EnableForeignKeysSql = "PRAGMA foreign_keys = ON;"

  RulesSql =
    `create table if not exists Rules (
      id integer primary key,
      name text,
      activeTime integer, -- When this rule takes effect (In unix time, aka epoch time)
      timestamp integer   -- Time this record is created (In unix time, aka epoch time)
     );
    `

  RulesIndexSql = "create index if not exists Rules_Index on Rules(name);"

  ActiveRulesSql =
    `create table if not exists ActiveRules (
      id integer primary key,
      ruleId integer references RulesSql(id)
     );
    `

  RulePropertiesSql =
    `create table if not exists RuleProperties (
      id integer primary key,
      ruleId integer references RulesSql(id),
      key text unique,
      value numeric
     );
    `
)

func ConfigureSqlite3(db *sql.DB) error {
  _, err := db.Exec(EnableForeignKeysSql)
  return err
}

func CreateAllTables(db *sql.DB) error {
  sqlStatements := []string{RulesSql, RulesIndexSql, ActiveRulesSql, RulePropertiesSql}

  for _, statement := range sqlStatements {
    _, err := db.Exec(statement)

    if err != nil {
      return err
    }
  }

  return nil
}

func DropTables(db *sql.DB, tableNames []string) error {
  for _, tableName := range tableNames {
    _, err := db.Exec("drop table ?", tableName)

    if err != nil {
      return err
    }
  }

  return nil
}

func DropAllTables(db *sql.DB) error {
  tableNames := []string{"Rules", "ActiveRules"}
  return DropTables(db, tableNames)
}
