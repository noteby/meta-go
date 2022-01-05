# meta-go

> IDE: VSCode https://code.visualstudio.com


> Go模块代理 https://goproxy.cn
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```


> 热加载工具 https://github.com/cosmtrek/air
```
go get -u github.com/cosmtrek/air
# export PATH=$PATH:/Users/home/go/bin
# source ~/.zshrc

air init 
# 配置介绍 https://github.com/cosmtrek/air/blob/master/air_example.toml
```
