Go projects => Modules
go.mod ==> requirements.txt in Python
go build para criar um executável do meu código
go build -o hello para especificar o nome do executável
go fmt formata o código fonte. Sempre utilizar antes de fazer um build ou implantar o código fonte no GitHub.
Go coloca o ponto e vírgula automaticamente. Não é necessário fazer.
go vet checa se estamos seguindo boas práticas de programação

Go usa makefile para automatizar tarefas de build de programas (arquivo Makefile)

Deixe seus objetivos claros ao escrever um programa em Go

Literais são especificados conforme a base.
    0b for binary
    0o for octal
    0x for hexadecimal

Literais do tipo Runa representam um caracter usando áspas simples, por exemplo: '\n'

Tipos compostos (objetos)

Arrays
Não é possível alterar o tamanho de uma array dinamicamente
Não é possível usar uma variável para determinar o tamanho de uma Array, ex: var x [tamanho]int. Funciona se usar var x[3]int.
Quase todas as limitações de Arrays são resolvidas através de Slices (fatias).

Slices
Podem crescer conforme precisamos.
Não precisa especificar tamanho na declaração: var x = []int{10,20,30}
nil significa ausência de valor para Slices e outros tipos de dados

y := []int{20,30,40}
x = append(x, y...) => os ... permitem a inclusão de todos os itens da Slice y na Slice x.

