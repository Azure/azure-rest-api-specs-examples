package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Datastore/AzureDataLakeGen1WServicePrincipal/createOrUpdate.json
func ExampleDatastoresClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewDatastoresClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"test-rg",
		"my-aml-workspace",
		"string",
		armmachinelearning.DatastoreData{
			Properties: &armmachinelearning.AzureDataLakeGen1Datastore{
				Description: to.Ptr("string"),
				Tags: map[string]*string{
					"string": to.Ptr("string"),
				},
				Credentials: &armmachinelearning.ServicePrincipalDatastoreCredentials{
					CredentialsType: to.Ptr(armmachinelearning.CredentialsTypeServicePrincipal),
					AuthorityURL:    to.Ptr("string"),
					ClientID:        to.Ptr("00000000-1111-2222-3333-444444444444"),
					ResourceURL:     to.Ptr("string"),
					Secrets: &armmachinelearning.ServicePrincipalDatastoreSecrets{
						SecretsType:  to.Ptr(armmachinelearning.SecretsTypeServicePrincipal),
						ClientSecret: to.Ptr("string"),
					},
					TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
				},
				DatastoreType: to.Ptr(armmachinelearning.DatastoreTypeAzureDataLakeGen1),
				StoreName:     to.Ptr("string"),
			},
		},
		&armmachinelearning.DatastoresClientCreateOrUpdateOptions{SkipValidation: to.Ptr(false)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
