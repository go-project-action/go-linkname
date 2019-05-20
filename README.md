## Go LinkName

## 測試

```
go run cmd/main.go
```

## 学习知识点

### 1. 使用 `//go:linkname` 開頭不能有空白

`// go:linkname` 這樣會錯誤因為 `//` 跟 `g` 中間有空白

如果有了空白會變成報錯

```bash
# command-line-arguments
main.main: relocation target github.com/go-project-action/go-linkname/src/b.Hi not defined
```

### 2. 没有实现的包要加一个 `.s` 汇编档案绕过检查

試試看把 `b.s` 刪掉

```bash
# github.com/go-project-action/go-linkname/src/b
src/b/b.go:7:6: missing function body
```

### 3. 使用 `//go:linkname` 的包要加上 `_ "unsafe"`

把 `unsafe` 拿到會報錯

```bash
# github.com/go-project-action/go-linkname/src/a
src/a/a.go:5:3: //go:linkname only allowed in Go files that import "unsafe"
```

## 問題:

### 1. 为什么 `time.Sleep` 的实现要这么搞?

`timeSleep` 函數有用調用到 `runtime/internal` 的代碼，所以必須寫在 `runtime` 包裡面。

> `timeSleep` 用到 `assignBucket()` 函數，`assignBucket()` 函數用到 `timers struct`，`timers struct` 用到 `internal/cpu` 的 `CacheLinePadSize`。

如果 `timeSleep` 大寫也可以給別人調用。不過為了讓代碼分的更詳細一點，就在 `time` 下面加了一個 `Sleep` 函數，並且把 `timeSleep` 連接到 `time.Sleep`。
