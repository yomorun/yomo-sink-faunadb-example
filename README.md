# yomo-sink-faunadb
The example shows how to implement a sink to write data to FaunaDB.



## Using Fauna

1. **Register at:** [https://fauna.com/](https://fauna.com/)

2. **Create Database and Collection**

   In this example, the `noise` Collection is created

3. **Generate `secret`**



## How to run the example

```shell script
FAUNA_SECRET=your-faunna-secret-here go run main.go
```
> BTW, `your-faunna-secret-here` is your `secret` in [https://fauna.com/](https://fauna.com/)

You will see the following message:

```shell script
2020/12/31 18:47:20 Starting sink server...
2020/12/31 18:47:20 ✅ Listening on 0.0.0.0:4141
```

You can config the address of yomo-sink-faunadb `localhost:4141` in [workflow.yaml](https://github.com/yomorun/yomo/blob/master/example/workflow.yaml), run [yomo-zipper](https://github.com/yomorun/yomo) and [yomo-source-demo](https://github.com/yomorun/yomo-source-demo), visit `noise` Collection in [https://fauna.com/](https://fauna.com/), then it will show the data in realtime.



## How yomo-sink-faunadb work


![sink](./sink.png)

