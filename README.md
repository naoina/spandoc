# Spandoc

Spandoc is a tool to generate documentations from schema of [Cloud Spanner](https://cloud.google.com/spanner/).

## Installation

```bash
go get -u github.com/naoina/spandoc/cmd/spandoc
```

## Usage

```sql
/* This is User table. */
CREATE TABLE User (
  ID STRING(MAX) NOT NULL,      -- This is PK
  Age INT64,                    -- Age of a user
  CreatedAt TIMESTAMP NOT NULL, -- Created time
  -- Updated time
  UpdatedAt TIMESTAMP NOT NULL OPTIONS (
    allow_commit_timestamp = true
  ),
) PRIMARY KEY(ID);
```

Then run command below.

```bash
spandoc schema.sql
```

Output:

```text
# User
This is User table.

COLUMN | TYPE | NOT NULL | PRIMARY | OPTIONS | DESCRIPTION
------ | ---- | ---------| ------- | ------- | -----------
ID | STRING(MAX) | true | true |  | This is PK
Age | INT64 |  |  |  | Age of a user
CreatedAt | TIMESTAMP | true |  |  | Created time
UpdatedAt | TIMESTAMP | true |  | ALLOW_COMMIT_TIMESTAMP | Updated time
```

## LICENSE

MIT
