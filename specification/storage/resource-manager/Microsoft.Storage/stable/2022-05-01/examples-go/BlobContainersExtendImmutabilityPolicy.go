package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/BlobContainersExtendImmutabilityPolicy.json
func ExampleBlobContainersClient_ExtendImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewBlobContainersClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ExtendImmutabilityPolicy(ctx, "res6238", "sto232", "container5023", "\"8d59f830d0c3bf9\"", &armstorage.BlobContainersClientExtendImmutabilityPolicyOptions{Parameters: &armstorage.ImmutabilityPolicy{
		Properties: &armstorage.ImmutabilityPolicyProperty{
			ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](100),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
