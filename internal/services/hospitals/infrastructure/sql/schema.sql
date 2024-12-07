-- sqlc에서는 버저닝을 할 수 없기 때문에 migration 파일은 따로 관리해야한다.
CREATE TABLE "hospital" (
    "createdAt"   timestamptz        NOT NULL,
    "updatedAt"   timestamptz        NOT NULL,
    "deletedAt"   timestamptz,
    "id"          UUID             PRIMARY KEY,
    "name"        VARCHAR(100)       NOT NULL,
    "phone"       VARCHAR(22)        NOT NULL,
    "city"        VARCHAR(50)        NOT NULL,
    "roadAddress" VARCHAR(255)      NOT NULL,
    "latitude"    FLOAT              NOT NULL,
    "longitude"   FLOAT              NOT NULL,
    "zipCode"     VARCHAR(20)        NOT NULL
);