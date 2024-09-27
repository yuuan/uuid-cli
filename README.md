# uuid-cli

## Installation

```console
$ go install github.com/yuuan/uuid-cli
```

## Usage

### UUID version 4 (v4)

#### Generate UUIDv4

```console
$ uuid-cli v4
e111b0a2-fd11-4192-a9e5-9a37b920f5f6
```

#### Generate Multiple UUIDv4

```console
$ uuid-cli v4 3
62e8be5a-293d-4b8a-8cad-5e3f3496048c
bb0fbafe-29a0-474d-a7b6-09b5d090eadc
fcae52fb-272b-4563-945a-32a421d13c22
```

#### Generate UUIDv4 and Display Details

```console
$ uuid-cli v4 -d
{
    "uuid": "c7f6cf41-68d0-440a-89f5-f759d8ba2c2e",
    "version": 4
}
```

### UUID version 7 (v7)

#### Generate UUIDv7

```console
$ uuid-cli v7
01922f21-9096-715b-9c17-8712bd44b46c
```

#### Generate Multiple UUIDv7

```console
$ uuid-cli v7 3
019231b2-fe2d-7808-bbb6-bf251f061274
019231b2-fe2d-7809-9aa0-75e532434ec9
019231b2-fe2d-780a-b20b-e6ffba9e0d26
```

#### Generate UUIDv7 and Display Details

```console
$ uuid-cli v7 -t 2024-10-01T12:34:56+09:00 -d
{
    "uuid": "01924625-3300-7a4a-9916-d46d077dc8a1",
    "version": 7,
    "timestamp": "2024-10-01T12:34:56+09:00"
}
```

#### Generate UUIDv7 with a Specific Timestamp

```console
$ uuid-cli v7 -t 2024-10-01T12:34:56+09:00 -d
{
    "uuid": "01924625-3300-74af-ab7d-f0b39e0ae43b",
    "version": 7,
    "timestamp": "2024-10-01T12:34:56+09:00"
}
```

```console
$ uuid-cli v7 -t "2024/10/01 12:34:56 JST" -d
{
    "uuid": "01924625-3300-74d9-93a4-3064ec72aeca",
    "version": 7,
    "timestamp": "2024-10-01T12:34:56+09:00"
}
```

### Display Details (parse)

#### Display Details of a Specified UUID

```console
$ uuid-cli parse 01924813-9580-7edf-8331-f854b396ba2e
{
    "uuid": "01924813-9580-7edf-8331-f854b396ba2e",
    "version": 7,
    "timestamp": "2024-10-01T21:34:56+09:00"
}
```

#### Display Details of Multiple Specified UUIDs

```console
$ uuid-cli parse 1a6991ce-7383-4e7e-a41a-873395a86445 01924813-9580-7edf-8331-f854b396ba2e
{
    "uuid": "1a6991ce-7383-4e7e-a41a-873395a86445",
    "version": 4
}
{
    "uuid": "01924813-9580-7edf-8331-f854b396ba2e",
    "version": 7,
    "timestamp": "2024-10-01T21:34:56+09:00"
}
```
