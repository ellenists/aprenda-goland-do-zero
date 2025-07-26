# Construção do Projeto

## Setup Inicial
Para lidar com mais de um pacote é preciso criar módulos, isso pode ser feito com o seguinte comando:
```
go mod init <nome-do-modulo>
```
Isso criará um arquivo `go.mod` que será como um `package.json` no nosso projeto, gerenciando as dependências do projeto.
> Nunca edite o arquivo `go.mod`, ele é atualizado automaticamente!

Para executar o nosso primeiro programa:
``` go
func main() {
	fmt.Println("Olá mundo!")
}
```
Podemos executar no diretório onde se encontra o arquivo que possui a função `main`:
```
go run main.go
```
Isso irá exibir "Olá mundo!" no console. Se quisermos enviar este programa para alguém que não possui go instalado podemos usar o comando:
```
go build
```
Este comando compila o código fonte do programa para um arquivo executável. O arquivo executável segue o sistema operacional onde foi gerado, mas se você usa Windows e quer compartilhar um executável do seu programa com um amigo que usa Linux?
Para isso existe o suporte a **cross-compilation**. Isso significa que podemos definir para que plataforma o comando `go build` deve gerar o executável:
```
GOOS=linux GOARCH=amd64 go build -o meuapp-linux
```
Isso vai gerar um arquivo binário que poderá ser executado no sistema operacional Linux assim:
```
./meuapp-linux
```
Para utilizar pacotes externos no nosso projeto, primeiro precisamos baixar o pacote para ele, usando o comando:
```
go get <endereco-do-pacote>
```
Isso irá baixar o conteúdo do pacote e atualizar o arquivo `go.mod` listando esta nova dependência.
O comando `go mod tidy` é usado para limpar o arquivo `go mod`, garantindo que ele esteja consistente com o código fonte do projeto, removendo dependências não utilizadas e adicionando as necessárias.
Go é uma linguagem fortemente tipada, isso significa que os tipos de dados das variáveis são estritamente definidos e verificados pelo compilador, não permitindo operações entre tipos incompatíveis sem conversão explícita.
``` go
x := 10      // x é int
x = "texto"  // Erro: não pode atribuir string a uma variável int
```
Para operar com tipos diferentes, você precisa converter explicitamente:
``` go
x := 10.5       // float64
y := int(x)     // Conversão explícita para int
fmt.Println(y)  // Saída: 10
```
