Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.4.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobContainersPutImmutabilityPolicy.json
func ExampleBlobContainersClient_CreateOrUpdateImmutabilityPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewBlobContainersClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateImmutabilityPolicy(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<container-name>",
		&armstorage.BlobContainersClientCreateOrUpdateImmutabilityPolicyOptions{IfMatch: nil,
			Parameters: &armstorage.ImmutabilityPolicy{
				Properties: &armstorage.ImmutabilityPolicyProperty{
					AllowProtectedAppendWrites:            to.BoolPtr(true),
					ImmutabilityPeriodSinceCreationInDays: to.Int32Ptr(3),
				},
			},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BlobContainersClientCreateOrUpdateImmutabilityPolicyResult)
}
```
