package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersExtendImmutabilityPolicy.json
func ExampleBlobContainersClient_ExtendImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.ExtendImmutabilityPolicy(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		"<if-match>",
		&armstorage.BlobContainersClientExtendImmutabilityPolicyOptions{Parameters: &armstorage.ImmutabilityPolicy{
			Properties: &armstorage.ImmutabilityPolicyProperty{
				ImmutabilityPeriodSinceCreationInDays: to.Int32Ptr(100),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientExtendImmutabilityPolicyResult)
}
