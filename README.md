# Criação de uma VPC na AWS com Remote State usando Terraform

Este repositório contém um exemplo de como criar uma VPC na AWS usando Terraform e gerenciar o estado remoto da VPC com S3. Além disso, este repositório também inclui testes da VPC criada usando Terratest.

## Pré-requisitos

- Terraform
- Go
- Terratest
- Uma conta na AWS com as permissões necessárias para criar uma VPC.

## Como usar este exemplo

1. Clone este repositório em sua máquina local:
   Copy code
   git clone https://github.com/seu-usuario/sua-vpc-terraform.git
2. Acesse o diretório do repositório:
   cd sua-vpc-terraform
3. Inicialize o Terraform:
   terraform init
4. Crie um arquivo terraform.tfvars com as variáveis de ambiente necessárias para acessar sua conta na AWS. Por exemplo:
   access_key = "SUA_CHAVE_DE_ACESSO_AWS"
   secret_key = "SUA_CHAVE_SECRETA_AWS"
5. Execute o Terraform para criar a VPC:
   terraform apply
6. Para testar a VPC criada, execute o Terratest:
   go test -v -timeout 30m -run TestVpc test/

## Arquivos

- terraform_test.go: O arquivo de teste de integração que usa o Terratest para testar o Terraform.
- go.mod: O arquivo que define as dependências do projeto Go.
- main.tf: O arquivo principal do Terraform que contém a configuração da VPC.
- network.tf: O arquivo que define os recursos de rede da VPC, como sub-redes, route tables e Internet Gateway.
- outputs.tf: O arquivo que define as saídas da VPC, como o ID e o endereço IP público do Internet Gateway.
  test/terraform_test.go: O arquivo de teste do Terratest que verifica se a VPC foi criada corretamente.

## Considerações finais

Lembre-se de limpar os recursos criados na AWS quando você terminar de usar este exemplo, executando o comando terraform destroy e substituir os exemplos de código pelo seu próprio código e de adaptar o README de acordo com o seu projeto.
