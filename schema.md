# Mitsuami DataBase Schema

## items

|    カラム名     |    型     | 初期値 |    制約    |      説明       |
|:-----------:|:--------:|:---:|:--------:|:-------------:|
|     ID      | varchar  | なし  |    PK    |      ID       |
|    Type     |   int    |  0  | Non Null | 本: 0, それ以外: 1 |
|    Name     | varchar  | なし  | Non Null |      物品名      |
|    Code     | varchar  | なし  | Non Null |  ISBN/JANなど   |
| Description | varchar  | なし  | Non Null |      説明       |
|    Image    | varchar  | なし  | Non Null |     画像URL     |
|  CreatedAt  | datetime | なし  | Non Null |               |
|  UpdatedAt  | datetime | なし  |          |               |
|  DeletedAt  | datetime | なし  |          |               |

## users

|    カラム名     |    型     | 初期値 |    制約    | 説明  |
|:-----------:|:--------:|:---:|:--------:|:---:|
|     ID      | varchar  | なし  |    PK    | ID  |
|    Name     | varchar  | なし  |   UNIQ   | 名前  |
| DisplayName | varchar  | なし  | Non Null | 表示名 |
|  CreatedAt  | datetime | なし  | Non Null |     |
|  UpdatedAt  | datetime | なし  |          |     |
|  DeletedAt  | datetime | なし  |          |     |

## log

|   カラム名    |    型     | 初期値 |    制約    |       説明        |
|:---------:|:--------:|:---:|:--------:|:---------------:|
|    ID     | varchar  | なし  |    PK    |       ID        |
|  UserID   | varchar  | なし  |    FK    |     ユーザーID      |
|  ItemID   | varchar  | なし  |    FK    |      備品ID       |
|  DueDate  | datetime | なし  | Non Null |      返却期限       |
|  Status   |   int    |  0  | Non Null | 0: 貸出中, 1: 返却済み |
| CreatedAt | datetime | なし  | Non Null |                 |
| UpdatedAt | datetime | なし  |          |                 |
