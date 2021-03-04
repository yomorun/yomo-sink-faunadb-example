# yomo-sink-faunadb

FaunaDB 🙌 YoMo. This example demonstrates how to implement a [yomo-sink](https://yomo.run/sink) to bulk write data in FaunaDB.

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

### Run `yomo-zipper`

Configure [YoMo-Zipper](https://yomo.run/zipper):

```yaml
name: YoMoZipper 
host: localhost
port: 9000
sinks:
  - name: FaunaDB
    host: localhost
    port: 4141
```

Start this zipper will listen on `9000` port, send data streams directly to `4141`:

```bash
cd ./zipper && yomo wf run

2020/12/31 19:45:15 Found 0 flows in zipper config
2020/12/31 19:45:15 Found 1 sinks in zipper config
2020/12/31 19:45:15 Sink 1: FaunaDB on localhost:4141
2020/12/31 19:45:15 Running YoMo workflow...
2020/12/31 19:45:15 ✅ Listening on 0.0.0.0:9000
2020/12/31 19:45:32 ✅ Connect to FaunaDB (localhost:4141) successfully.
```

### Emulate a data source for testing

```bash
cd source && go run main.go

2020/12/31 19:45:40 ✅ Connected to yomo-zipper localhost:9000
2020/12/31 19:45:40 ✅ Emit 143.58589 to yomo-zipper
2020/12/31 19:45:40 ✅ Emit 101.46548 to yomo-zipper
```

This will start a [YoMo-Source](https://yomo.run/source), demonstrate a random float every 100ms to [YoMo-Zipper](https://yomo.run/zipper).

## How yomo-sink-faunadb work

![sink](./sink.png)
