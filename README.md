# ginServer
a restful api server 

## 1、数据库搭建
* 1、拉取 mysql 5.7 的 docker 镜像
   
   `docker pull mysql:5.7`
   
* 2、 启动镜像
    
    `docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=qwer1234 -d mysql:5.7`

* 3、创建数据库
    
    创建数据库 gin_server
    
    字符集 utf8mb4

## 2、构建 docker 容器

####1、直接使用 golang 镜像构建容器
    
    这种方式构建的容器很大（docker ps -as 查看占用），因为 golang 官方镜像中包括编译和运行环境，外加一堆 GCC、build工具

#####1、编写 dockerfile
      
    FROM golang:latest
    
    WORKDIR $GOPATH/src/ginServer
    COPY . $GOPATH/src/ginServer
    RUN go build .
    
    EXPOSE 3000
    
    ENTRYPOINT ["./ginserver"]

#####2、构建镜像
    
    docker build -t ginserver .
    
#####3、启动容器

    docker run -p 3000:3000 --link mysql:mysql ginserver
    
    
####2、使用 Scratch 镜像

    Scratch 镜像，简洁、小巧、基本是个空镜像
    
#####1、修改 dockerfile

    FROM scratch
    
    WORKDIR $GOPATH/src/ginServer
    COPY . $GOPATH/src/ginServer
    
    EXPOSE 3000
    CMD ["./ginserver"]
    
#####2、编译可执行文件
    
    编译生成的可执行文件会依赖一些库，并且是动态链接。这里因为使用的是 Scratch 镜像，他是空镜像，因此需要将生成的可执行文件静态链接所
    依赖的库。
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goserver .
    
#####3、构建镜像
    
    同上
    
#####4、启动容器

    同上