package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01a71545e82bb98b8137d3038150c436d46a59ed/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Credentials_Create.json
func ExampleCredentialOperationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCredentialOperationsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleCredential", armdatafactory.ManagedIdentityCredentialResource{
		Properties: &armdatafactory.ManagedIdentityCredential{
			Type: to.Ptr("ManagedIdentity"),
			TypeProperties: &armdatafactory.ManagedIdentityTypeProperties{
				ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami"),
			},
		},
	}, &armdatafactory.CredentialOperationsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedIdentityCredentialResource = armdatafactory.ManagedIdentityCredentialResource{
	// 	Name: to.Ptr("exampleCredential"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/credentials"),
	// 	Etag: to.Ptr("0a0062d4-0000-0000-0000-5b245bcf0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/credentials/exampleCredential"),
	// 	Properties: &armdatafactory.ManagedIdentityCredential{
	// 		Type: to.Ptr("ManagedIdentity"),
	// 		TypeProperties: &armdatafactory.ManagedIdentityTypeProperties{
	// 			ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami"),
	// 		},
	// 	},
	// }
}
