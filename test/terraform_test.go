package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAwsVpc(t *testing.T) {
	t.Parallel()

	// Opções de configuração do Terratest
	terraformOptions := &terraform.Options{
		// O caminho para o diretório que contém o código do Terraform
		TerraformDir: "path/to/terraform/code",

		// Define valores para as variáveis do Terraform
		Vars: map[string]interface{}{
			"access_key": "ACCESS_KEY",
			"secret_key": "SECRET_KEY",
			"region":     "REGION",
			"cidr_block": "CIDR_BLOCK",
		},
	}

	// Executa o "terraform init" e o "terraform apply"
	terraform.InitAndApply(t, terraformOptions)

	// Verifica se a VPC foi criada com sucesso
	vpcID := terraform.Output(t, terraformOptions, "vpc_id")
	assert.NotEmpty(t, vpcID)
}
