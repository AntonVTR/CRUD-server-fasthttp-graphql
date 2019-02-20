This is probation to write a fasthttp + graphql and CRUD object

to run "go run ./main.go"


requests via browser
http://localhost:8080/?query=mutation+_{create(FirstName:"Inca",LastName:"Kola",Gender:"other",Position:"CEO",Salary:6000){id,FirstName,LastName,Gender, Position, Salary}}

http://localhost:8080/product?query={employer(id:106){id,FirstName,LastName,Gender,Position,Salary}}
		
http://localhost:8080/product?query={list{id,FirstName,LastName,Gender,Position,Salary}}
		
			
http://localhost:8080/?query=mutation+_{update(id:101,Salaty:9000){id,FirstName,LastName,Gender, Position, Salary}}
		
http://localhost:8080/product?query=mutation+_{delete(id:106){id,FirstName,LastName,Gender, Position, Salary}}

or terminal

curl 'http://localhost:8080/?query=mutation+_{create(FirstName:"Inca",LastName:"Kola",Gender:"other",Position:"CEO",Salary:6000){id,FirstName,LastName,Gender, Position, Salary}}'

curl 'http://localhost:8080/product?query={employer(id:106){id,FirstName,LastName,Gender,Position,Salary}}'

curl 'http://localhost:8080/product?query={list{id,FirstName,LastName,Gender,Position,Salary}}'

curl 'http://localhost:8080/?query=mutation+_{update(id:101,Salaty:9000){id,FirstName,LastName,Gender, Position, Salary}}'

curl 'http://localhost:8080/product?query=mutation+_{delete(id:106){id,FirstName,LastName,Gender, Position, Salary}}'
