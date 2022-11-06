package main

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	kms "github.com/alibabacloud-go/kms-20160120/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/aliyun/credentials-go/credentials"
	log "github.com/sirupsen/logrus"
)

func main() {
	credentialConfig := &credential.Config{
		AccessKeyId:     tea.String("....."),
		AccessKeySecret: tea.String("...."),
		RoleArn:         tea.String("acs:ram::562383.....:role/ma...-ad...-role"),
		RoleSessionName: tea.String("test"),
		Type:            tea.String("ram_role_arn"),
	}

	credentials, err := credential.NewCredential(credentialConfig)
	if err != nil {
		log.WithError(err).Fatal("error creating Alibaba credentials")
	}

	config := &openapi.Config{
		Credential: credentials,
		RegionId:   tea.String("ap-southeast-1"),
	}
	options := &util.RuntimeOptions{}

	client, err := kms.NewClient(config)
	if err != nil {
		log.WithError(err).Fatal("error init KMS client")
	}

	request := &kms.CreateSecretRequest{
		SecretName: tea.String("sectest"),
		VersionId:  tea.String("v1"),
		SecretData: tea.String("test data"),
	}

	resp, err := client.CreateSecretWithOptions(request, options)
	if err != nil {
		log.WithError(err).Fatalf("error creating secret [%s] version [%s]", *request.SecretName, *request.VersionId)
	}

	log.Info(resp)
}
