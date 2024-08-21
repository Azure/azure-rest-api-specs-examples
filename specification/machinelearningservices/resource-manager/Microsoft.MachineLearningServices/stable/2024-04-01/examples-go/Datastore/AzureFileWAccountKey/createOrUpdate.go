package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9778042723206fbc582306dcb407bddbd73df005/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Datastore/AzureFileWAccountKey/createOrUpdate.json
func ExampleDatastoresClient_CreateOrUpdate_createOrUpdateDatastoreAzureFileStoreWAccountKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatastoresClient().CreateOrUpdate(ctx, "test-rg", "my-aml-workspace", "string", armmachinelearning.Datastore{
		Properties: &armmachinelearning.AzureFileDatastore{
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
			DatastoreType: to.Ptr(armmachinelearning.DatastoreTypeAzureFile),
			AccountName:   to.Ptr("string"),
			Endpoint:      to.Ptr("string"),
			FileShareName: to.Ptr("string"),
			Protocol:      to.Ptr("string"),
		},
	}, &armmachinelearning.DatastoresClientCreateOrUpdateOptions{SkipValidation: to.Ptr(false)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Datastore = armmachinelearning.Datastore{
	// 	Name: to.Ptr("string"),
	// 	Type: to.Ptr("string"),
	// 	ID: to.Ptr("string"),
	// 	SystemData: &armmachinelearning.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T12:34:56.999Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmachinelearning.AzureFileDatastore{
	// 		Description: to.Ptr("string"),
	// 		Tags: map[string]*string{
	// 			"string": to.Ptr("string"),
	// 		},
	// 		Credentials: &armmachinelearning.AccountKeyDatastoreCredentials{
	// 			CredentialsType: to.Ptr(armmachinelearning.CredentialsTypeAccountKey),
	// 			Secrets: &armmachinelearning.AccountKeyDatastoreSecrets{
	// 				SecretsType: to.Ptr(armmachinelearning.SecretsTypeAccountKey),
	// 			},
	// 		},
	// 		DatastoreType: to.Ptr(armmachinelearning.DatastoreTypeAzureFile),
	// 		AccountName: to.Ptr("string"),
	// 		Endpoint: to.Ptr("string"),
	// 		FileShareName: to.Ptr("string"),
	// 		Protocol: to.Ptr("string"),
	// 	},
	// }
}
