-- sqlc에서는 버저닝을 할 수 없기 때문에 migration 파일은 따로 관리해야한다.
CREATE TABLE "user" (
    "createdAt"   timestamptz        NOT NULL,
    "updatedAt"   timestamptz        NOT NULL,
    "deletedAt"   timestamptz,
    "id"          UUID             PRIMARY KEY,
    "email"       VARCHAR(50)       NOT NULL,
    "nickname"    VARCHAR(100)     NOT NULL,
    "providers"   text[]             NOT NULL,
    CONSTRAINT "uniq_email" UNIQUE ("email")
);