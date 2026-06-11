package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/Datastore/AzureFileWAccountKey/createOrUpdate.json
func ExampleDatastoresClient_CreateOrUpdate_createOrUpdateDatastoreAzureFileStoreWAccountKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatastoresClient().CreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearning.Datastore{
		Properties: &armmachinelearning.AzureFileDatastore{
			Description: to.Ptr("string"),
			AccountName: to.Ptr("string"),
			Credentials: &armmachinelearning.AccountKeyDatastoreCredentials{
				CredentialsType: to.Ptr(armmachinelearning.CredentialsTypeAccountKey),
				Secrets: &armmachinelearning.AccountKeyDatastoreSecrets{
					Key:         to.Ptr("string"),
					SecretsType: to.Ptr(armmachinelearning.SecretsTypeAccountKey),
				},
			},
			DatastoreType: to.Ptr(armmachinelearning.DatastoreTypeAzureFile),
			Endpoint:      to.Ptr("string"),
			FileShareName: to.Ptr("string"),
			Properties:    nil,
			Tags: map[string]*string{
				"string": to.Ptr("string"),
			},
			Protocol: to.Ptr("string"),
		},
	}, &armmachinelearning.DatastoresClientCreateOrUpdateOptions{
		SkipValidation: to.Ptr(false)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.DatastoresClientCreateOrUpdateResponse{
	// 	Datastore: armmachinelearning.Datastore{
	// 		Name: to.Ptr("string"),
	// 		Type: to.Ptr("string"),
	// 		ID: to.Ptr("string"),
	// 		Properties: &armmachinelearning.AzureFileDatastore{
	// 			Description: to.Ptr("string"),
	// 			AccountName: to.Ptr("string"),
	// 			Credentials: &armmachinelearning.AccountKeyDatastoreCredentials{
	// 				CredentialsType: to.Ptr(armmachinelearning.CredentialsTypeAccountKey),
	// 				Secrets: &armmachinelearning.AccountKeyDatastoreSecrets{
	// 					SecretsType: to.Ptr(armmachinelearning.SecretsTypeAccountKey),
	// 				},
	// 			},
	// 			DatastoreType: to.Ptr(armmachinelearning.DatastoreTypeAzureFile),
	// 			Endpoint: to.Ptr("string"),
	// 			FileShareName: to.Ptr("string"),
	// 			Properties: nil,
	// 			Tags: map[string]*string{
	// 				"string": to.Ptr("string"),
	// 			},
	// 			Protocol: to.Ptr("string"),
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
