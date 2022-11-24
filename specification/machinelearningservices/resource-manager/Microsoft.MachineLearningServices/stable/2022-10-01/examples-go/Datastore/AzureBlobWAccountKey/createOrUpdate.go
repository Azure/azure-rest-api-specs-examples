package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Datastore/AzureBlobWAccountKey/createOrUpdate.json
func ExampleDatastoresClient_CreateOrUpdate_createOrUpdateDatastoreAzureBlobWAccountKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewDatastoresClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearning.Datastore{
		Properties: &armmachinelearning.AzureBlobDatastore{
			Description: to.Ptr("string"),
			Tags: map[string]*string{
				"string": to.Ptr("string"),
			},
			Credentials: &armmachinelearning.AccountKeyDatastoreCredentials{
				CredentialsType: to.Ptr(armmachinelearning.CredentialsTypeAccountKey),
				Secrets: &armmachinelearning.AccountKeyDatastoreSecrets{
					SecretsType: to.Ptr(armmachinelearning.SecretsTypeAccountKey),
					Key:         to.Ptr("string"),
				},
			},
			DatastoreType: to.Ptr(armmachinelearning.DatastoreTypeAzureBlob),
			AccountName:   to.Ptr("string"),
			ContainerName: to.Ptr("string"),
			Endpoint:      to.Ptr("core.windows.net"),
			Protocol:      to.Ptr("https"),
		},
	}, &armmachinelearning.DatastoresClientCreateOrUpdateOptions{SkipValidation: to.Ptr(false)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
