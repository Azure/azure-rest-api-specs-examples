package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v11"
)

// Generated from example definition: 2018-06-01/Credentials_Create.json
func ExampleCredentialOperationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCredentialOperationsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleCredential", armdatafactory.CredentialResource{
		Properties: &armdatafactory.ManagedIdentityCredential{
			Type: to.Ptr("ManagedIdentity"),
			TypeProperties: &armdatafactory.ManagedIdentityTypeProperties{
				ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatafactory.CredentialOperationsClientCreateOrUpdateResponse{
	// 	CredentialResource: &armdatafactory.CredentialResource{
	// 		Name: to.Ptr("exampleCredential"),
	// 		Type: to.Ptr("Microsoft.DataFactory/factories/credentials"),
	// 		Etag: to.Ptr("0a0062d4-0000-0000-0000-5b245bcf0000"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/credentials/exampleCredential"),
	// 		Properties: &armdatafactory.ManagedIdentityCredential{
	// 			Type: to.Ptr("ManagedIdentity"),
	// 			TypeProperties: &armdatafactory.ManagedIdentityTypeProperties{
	// 				ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami"),
	// 			},
	// 		},
	// 	},
	// }
}
