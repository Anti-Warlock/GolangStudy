# Go新建工程的步骤

***Tips***

***Go Module是源自1.11,之前的模式都是采用Go Path,就是必须在Go Root下的src文件夹新建工程,自由度不够高。所以下面粗略概括下采用Go Module的模式***

## Go Modules初始化项目步骤
* [1.](#1) 开启G0 Modules模块
    + a. 通过Go Env 查看修改Go Modules是否开启<br/> 
    `go env -w GO111MODULE=on`
* [2.](#2) 初始化项目
    + a. 创建一个项目(不要求在$GOPATH/src)<br/>
    `mkdir -p /path/projectName`  
    + b. 创建go.mod文件,同时指定当前项目的模块名称<br/>
    `go mod init projectName`
    + c. 下载依赖库<br/>
        + <c> 手动下载依赖<br>
        `go get github.example.com/project`
        + <c> 自动下载依赖<br>
        `go run example.go`
* [3.](#3) 生成go.mod和go.sum文件       
    + a. 作用<br/>
        &emsp;1.go.mod中的`//indirect`表示间接依赖<br/>
        &emsp;1.罗列当前项目的依赖版本<br/>
        &emsp;2.h1.hash表示整体项目的全部文件校验和来生成的Hash值,如果不存在,表示其中某些依赖可能用不上<br/>
        




