# meta-go

> IDE: VSCode https://code.visualstudio.com

> 安装 Go
```
查看系统信息
# uname -a

下载安装包
# wget https://golang.google.cn/dl/go1.17.5.linux-amd64.tar.gz

解压
# tar -C /opt/ -zvxf go1.17.5.linux-amd64.tar

设置环境变量
# vim /etc/profile.d/x_go.sh 
写入 export PATH=$PATH:/opt/go/bin

查看版本
# go version
```

> Go 模块代理 https://goproxy.cn
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

> 安装 Git
```
# yum install -y git
```

> 安装 zsh
```
# yum install -y zsh
```

> 安装 oh-my-zsh
```
sh -c "$(wget https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh -O -)"
```

> 安装 MySQL
```
查看是否安装过 MySQL
# rpm -qa | grep -i mysql

删除 MySQL
# yum -y remove mysql80-community-release-el7-1.noarch

重新配置安装源（https://dev.mysql.com/downloads/repo/yum/）
# rpm -Uvh https://dev.mysql.com/get/mysql80-community-release-el7-3.noarch.rpm

安装
# yum --enablerepo=mysql80-community install mysql-community-server

- 如果有提示
All matches were filtered out by modular filtering for argument: mysql-community-server
Error: Unable to find a match: mysql-community-server
# yum module disable mysql

- 重新执行
# yum --enablerepo=mysql80-community install mysql-community-server

启动/停止服务
# service mysqld start
# service mysqld stop

查看服务状态
# service mysqld status

查看临时密码
# more /var/log/mysqld.log |grep temporary

修改密码
mysql> alter user user() identified by "new password";

查看密码验证策略
mysql> show variables like 'validate_password.%';

设置密码复杂度
mysql> set global validate_password.policy=0;

设置可远程访问主机
mysql> create user root@'%' identified by "password";

设置权限
mysql> grant all privileges on *.* to root@'%' with grant option;

刷新
mysql> flush privileges;
mysql> flush hosts;
```

> 热加载工具 https://github.com/cosmtrek/air
```
go get -u github.com/cosmtrek/air
# export PATH=$PATH:/Users/home/go/bin
# source ~/.zshrc

air init 
# 配置介绍 https://github.com/cosmtrek/air/blob/master/air_example.toml
```

> fr-editor 20211220
```
license 提示定位：
_init: function k()
var d = new Image;

底部 logo 定位
<span>Powered by</span>
```