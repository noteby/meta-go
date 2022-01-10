#!/bin/sh

# 配置免密登录，
# 本机执行：ssh-keygen -t rsa 
# 查看公钥：cat ~/.ssh/id_rsa.pub
# 将公钥拷贝至服务器文件： ~/.ssh/authorized_keys

while read line; do
    eval "$line"
done < deploy.ini

projectdir=$projectdir
supervisorconf=$supervisorconf

hostname=$hostname
username=$username
gituser=$gituser
gitpass=$gitpass

expect << EOF
    set timeout 20

    spawn ssh $username@$hostname
    expect "$ " { send "cd $projectdir\r" }
    expect "$ " { send "git pull --rebase\r" }
    expect ": " { send "$gituser\r" }
    expect ": " { send "$gitpass\r" }
    expect "$ " { send "go build -o bin/ main.go\r" }    
    expect "$ " { send "supervisorctl -c $supervisorconf restart meta-go\r" }
    expect "$ " { send "exit\r" }
    exec sleep 3 
    
    expect eof
EOF