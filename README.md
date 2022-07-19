# cloc
统计指定目录下的各类语言代码文件的行数，支持的特性：
- 支持指定类型的代码文件
- 支持分开统计有效的代码行数、注释行数、空白行数
- 支持按照不同的字段进行排序展示

## 支持的语言

| language | extension |
| --- | --- |
| javaScript | js, jsx, mjs, cjs |
| JSON | json |
| TypeScript | tsx, ts |
| HTML | html, htm |
| SCSS | scss |
| CSS | css |
| Golang | go |

## 示例

```shell
cloc /src -sort code -order asc
```
会打印输出为：
```
│----.gitignore
│----LICENSE
│----README.md
│----cloc
│----dir_command.go ---------->[codes: 112]
│----go.mod
│----go.sum
│----main.go ---------->[codes: 58]
│----options.go ---------->[codes: 112]
│----options_test.go ---------->[codes: 37]
│----page_command.go ---------->[codes: 44]
│----page_judge.go ---------->[codes: 65]
│----page_point.go ---------->[codes: 67]
│----testdata
│----│----dir
│----│----│----heihei
│----│----│----│----here
│----│----│----hello.go ---------->[codes: 0]
│----│----│----text
│----│----heihei.go ---------->[codes: 37]
│----│----hello.html ---------->[codes: 0]
│----│----text
│----util.go ---------->[codes: 20]

[codes total]: 552

```

## 参数说明

| sort | 说明 |
| ----- | ----- |
| files | 文件数量 |
| code | 代码数量 |
| comment | 注释行数 |
| blank | 空白行数 |


| order | 说明 |
| ----- | ----- |
| desc | 降序 |
| asc | 升序 |

`order`不支持统计单个代码文件
