# API Document

## login

## logout

## add equipment `POST /items`

```jsonc
{
    // type 0: book/1: other, required
    "type": 0,
    // isbn code
    "code": "9784873119045",
    "name": "プログラミング TypeScript",
    "description": "プログラミング言語TypeScriptの解説書。TypeScriptの型に関する基礎的な内容からその応用、エラー処理の手法、非同期プログラミング、各種フレームワークの利用法、既存のJavaScriptプロジェクトのTypeScript移行の方法まで、言語全般を総合的に解説します。本書全体を通じて、TypeScriptの洗練された型システムを最大限活用するために、コードをどのように記述すべきか、なぜそうすべきかを学べます。",
    // image
    "image": "https://cover.openbd.jp/9784873119045.jpg"
}
```

## get item data `GET /items/{:id}`

```jsonc
{
    "id": "w0fi88c",
    // type 0: book/1: other, required
    "type": 0,
    // isbn code
    "code": "9784873119045",
    "name": "プログラミング TypeScript",
    "description": "プログラミング言語TypeScriptの解説書。TypeScriptの型に関する基礎的な内容からその応用、エラー処理の手法、非同期プログラミング、各種フレームワークの利用法、既存のJavaScriptプロジェクトのTypeScript移行の方法まで、言語全般を総合的に解説します。本書全体を通じて、TypeScriptの洗練された型システムを最大限活用するために、コードをどのように記述すべきか、なぜそうすべきかを学べます。",
    // image
    "image": "https://cover.openbd.jp/9784873119045.jpg",
    "status": {
        // status: 0 lending/1 available/2 reserved, required
        "status": 1,
        "user": "John Doe",
        "dueDate": "2022-05-30"
    }
}
```

## start using (borrow item) `POST /items/{:id}/using`

response:

```jsonc
{
    "status": "ok"
}
```

## end using (return) `DELETE /items/{:id}/using`

response:

```jsonc
{
    "status": "ok"
}
```

## get item list `GET /items`

```jsonc
[
    {
        "id": "w0fi88c",
        // type 0: book/1: other, required
        "type": 0,
        // isbn code
        "code": "9784873119045",
        "name": "プログラミング TypeScript",
        "description": "プログラミング言語TypeScriptの解説書。TypeScriptの型に関する基礎的な内容からその応用、エラー処理の手法、非同期プログラミング、各種フレームワークの利用法、既存のJavaScriptプロジェクトのTypeScript移行の方法まで、言語全般を総合的に解説します。本書全体を通じて、TypeScriptの洗練された型システムを最大限活用するために、コードをどのように記述すべきか、なぜそうすべきかを学べます。",
        // image
        "image": "https://cover.openbd.jp/9784873119045.jpg",
        "status": {
            // status: 0 lending/1 available/2 reserved, required
            "status": 1,
            "user": "John Doe",
            "dueDate": "2022-05-30"
        }
    }
]
```
