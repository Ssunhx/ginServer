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
    