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

You can use the command `yomo wf dev workflow.yaml` to run [yomo-zipper](https://yomo.run/zipper) which will automatically emit the real noise data from CELLA office, or run `yomo wf run workflow.yaml` with the specific `yomo-source`. See [yomo-zipper](https://yomo.run/zipper#how-to-config-and-run-yomo-zipper) for details.
After running `yomo-zipper`, it will add the real-time data to `noise` collection in [https://fauna.com/](https://fauna.com/).

## How yomo-sink-faunadb work

![sink](./sink.png)
