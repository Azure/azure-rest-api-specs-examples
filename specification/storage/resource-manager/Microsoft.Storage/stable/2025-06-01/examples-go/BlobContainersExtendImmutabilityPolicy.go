package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4e9df3afd38a1cfa00a5d49419dce51bd014601f/specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/BlobContainersExtendImmutabilityPolicy.json
func ExampleBlobContainersClient_ExtendImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlobContainersClient().ExtendImmutabilityPolicy(ctx, "res6238", "sto232", "container5023", "8d59f830d0c3bf9", &armstorage.BlobContainersClientExtendImmutabilityPolicyOptions{Parameters: &armstorage.ImmutabilityPolicy{
		Properties: &armstorage.ImmutabilityPolicyProperty{
			ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](100),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ImmutabilityPolicy = armstorage.ImmutabilityPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/containers/immutabilityPolicies"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6238/providers/Microsoft.Storage/storageAccounts/sto232/blobServices/default/containers/container5023/immutabilityPolicies/default"),
	// 	Etag: to.Ptr("\"8d57a8b2ff50332\""),
	// 	Properties: &armstorage.ImmutabilityPolicyProperty{
	// 		ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](100),
	// 		State: to.Ptr(armstorage.ImmutabilityPolicyStateLocked),
	// 	},
	// }
}
