<div align="center">
    <h1>
        <a href="https://github.com/alicfeng/gutilib">
            Golang Library On Lampstand
        </a>
    </h1>
    <p>
      Golang底座依赖库
      <br/>
      有趣的灵魂，为工程和组件赋能，让你的组件更加标准化，让你的精力聚焦业务
    </p>
</div>



## 🪤 底座简介

顾名思义，它是一个基础库，封装 标准库不存在的函数而实质场景中常规且常见的作用函数 于此，只为业务代码更加简洁有力，释放你的精力聚焦在业务上，仅此而已。



## ✨ 支持特性

- [x] 简而有力的表达式

- [x]  集合常规场景函数

- [x] 文件类增强化函数



## 🚀 快速安装

```shell
go get github.com/alicfeng/gutilib@v0.0.4
```



## 📒 使用文档

#### 表达式

三元表达式

```go
// 支持任意类型的值类型
what := expression.Ternary(true, "hello", "hi") // what:hello
```



#### 集合

> 基于泛型支持任意键值类型

`map` 过滤

```go
// 定义我的成绩 Map
data := map[string]int64{
  "math":    88,
  "english": 90,
  "chinese": 60,
}

// core api:检索高于 80 分的各科成绩
result := collection.MapFilter(data, func(key string, value int64) bool {
  return value > 80
}) // map[english:90 math:88] --- PASS: TestCollectionMapFilter (0.00s)
```

`map` 断言是否存在指定的键

```go
// 定义我的成绩 Map
data := map[string]int64{
  "math":    88,
  "english": 90,
  "chinese": 60,
}

result := collection.MapExistKey(data, "math") // result : true
```

`map` 提取所有的键名

```go
// 定义我的成绩 Map
data := map[string]int64{
  "math":    88,
  "english": 90,
  "chinese": 60,
}

result := collection.MapKeys(data) // result : []string{"math","english","chinese"}
```

`map`获取所有的值

```go
// 定义我的成绩 Map
data := map[string]int64{
  "math":    88,
  "english": 90,
  "chinese": 60,
}

result := collection.MapValues(data) // result : []int64{88,90,60}
```

`map` 遍历

```go
// 定义我的成绩 Map
data := map[string]int64{
  "math":    88,
  "english": 90,
  "chinese": 60,
}

collection.MapForEach(data, func(key string,value int64){
  fmt.Print(key, value)
})
```

`map` 断言指定值是否被包含

```go
// 定义我的成绩 Map
data := map[string]int64{
  "math":    88,
  "english": 90,
  "chinese": 60,
}

result := collection.MapIsContain(data, 100) // result : false
```



#### 切片

切片过滤

```go
result := collection.SliceFilter([]int64{1, 2, 3}, func(index int, item int64) bool {
    return item == 2
})// result : [2]
```

切片截取

```go
result := collection.SliceSub([]int64{1, 2, 3}, 0, 1) // result : [1, 2]
```


#### 文件系统

判断路径是否存在

```go
var ok bool = filesystem.PathExist("/usr/local/ai/model.pb")
```