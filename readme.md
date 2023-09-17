# Go WebSocket 示例

这是一个简单的 Go WebSocket 示例，包括一个服务器（server）和一个客户端（client），用于发送学生信息。

## 运行服务器

1. 进入 `server` 文件夹：

    ```shell
    cd server
    ```

2. 使用以下命令编译服务器代码：

    ```shell
    go build -o server main.go
    ```

3. 运行服务器：

    ```shell
    ./server
    ```

服务器将在 `localhost:8080` 上启动，并等待来自客户端的WebSocket连接。

## 运行客户端

1. 进入 `client` 文件夹：

    ```shell
    cd client
    ```

2. 使用以下命令编译客户端代码：

    ```shell
    go build -o client main.go
    ```

3. 运行客户端：

    ```shell
    ./client
    ```

客户端将建立WebSocket连接到服务器（默认连接到 `ws://localhost:8080/ws`），并发送学生信息到服务器。

## 注意事项

- 请确保服务器在运行时客户端也在运行，以便建立WebSocket连接并发送数据。
- 在运行服务器和客户端之前，确保你已经正确安装了Go并配置了工作环境。
- 你可以根据需要修改客户端和服务器代码，以适应你的项目要求。

祝贺你成功运行了Go WebSocket示例！
