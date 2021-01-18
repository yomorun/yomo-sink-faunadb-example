# yomo-sink-faunadb

The example shows how to implement a `yomo-sink` to write data to FaunaDB.

## Using Fauna

1. **Register at:** [https://fauna.com/](https://fauna.com/)

2. **Create Database and Collection**

   In this example, the `noise` collection is created.

3. **Generate `secret`**

## How to run the example

```bash
FAUNA_SECRET=your-faunna-secret-here go run main.go
```

> BTW, `your-faunna-secret-here` is your `secret` in [https://fauna.com/](https://fauna.com/)

You will see the following message:

```shell
2020/12/31 18:47:20 Starting sink server...
2020/12/31 18:47:20 ✅ Listening on 0.0.0.0:4141
```

You can run [yomo-zipper](https://yomo.run/zipper) and [yomo-source-demo](https://github.com/yomorun/yomo-source-demo), then it will add the real-time data to `noise` collection in [https://fauna.com/](https://fauna.com/).

See [yomo-zipper](https://yomo.run/zipper#how-to-config-and-run-yomo-zipper) for details.

## How yomo-sink-faunadb work

![sink](./sink.png)
