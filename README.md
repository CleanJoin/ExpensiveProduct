# ExpensiveProduct
<h3>Запуск:</h3>
go run main.go -input name_file <br>
  - input name_file (flag): указываем фаил с расширением (default = db.json)<br>
  файлы помещать в директорию src<br>
 
<h3> Docker:</h3>
 docker build -t myapp .
 docker run  --rm  myapp  go run main.go  -input db.json <br>
 
