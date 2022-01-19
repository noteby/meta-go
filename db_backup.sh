#!/bin/sh

# 配置免密登录，
# 本机执行：ssh-keygen -t rsa 
# 查看公钥：cat ~/.ssh/id_rsa.pub
# 将公钥拷贝至服务器文件： ~/.ssh/authorized_keys

while read line; do
    eval "$line"
done < db_backup.ini

backupdir=$backupdir
backuppath=$backuppath

hostname=$hostname
username=$username

dbname=$dbname
dbuser=$dbuser
dbpass=$dbpass

localbackupdir=$localbackupdir

# 数据库备份
expect<<-EOF
    set timeout 20

    spawn ssh $username@$hostname
    expect "$ " { send "mkdir -p $backupdir\r" }
    expect "$ " { send "mysqldump -u $dbuser -p $dbname > $backuppath\r" }
    expect "password:" { send "$dbpass\r" }
    expect "$ " { send "exit\r" }
    exec sleep 3
    expect eof
EOF

# 拷贝至本地
expect<<-EOF
    set timeout 20

    spawn scp $username@$hostname:$backuppath $localbackupdir
    exec sleep 3 
    expect eof
EOF