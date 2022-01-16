Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.3.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountSetManagementPolicy.json
func ExampleManagementPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewManagementPoliciesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.ManagementPolicyName("default"),
		armstorage.ManagementPolicy{
			Properties: &armstorage.ManagementPolicyProperties{
				Policy: &armstorage.ManagementPolicySchema{
					Rules: []*armstorage.ManagementPolicyRule{
						{
							Name: to.StringPtr("<name>"),
							Type: armstorage.RuleType("Lifecycle").ToPtr(),
							Definition: &armstorage.ManagementPolicyDefinition{
								Actions: &armstorage.ManagementPolicyAction{
									BaseBlob: &armstorage.ManagementPolicyBaseBlob{
										Delete: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Float32Ptr(1000),
										},
										TierToArchive: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Float32Ptr(90),
										},
										TierToCool: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Float32Ptr(30),
										},
									},
									Snapshot: &armstorage.ManagementPolicySnapShot{
										Delete: &armstorage.DateAfterCreation{
											DaysAfterCreationGreaterThan: to.Float32Ptr(30),
										},
									},
								},
								Filters: &armstorage.ManagementPolicyFilter{
									BlobTypes: []*string{
										to.StringPtr("blockBlob")},
									PrefixMatch: []*string{
										to.StringPtr("olcmtestcontainer1")},
								},
							},
							Enabled: to.BoolPtr(true),
						},
						{
							Name: to.StringPtr("<name>"),
							Type: armstorage.RuleType("Lifecycle").ToPtr(),
							Definition: &armstorage.ManagementPolicyDefinition{
								Actions: &armstorage.ManagementPolicyAction{
									BaseBlob: &armstorage.ManagementPolicyBaseBlob{
										Delete: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Float32Ptr(1000),
										},
										TierToArchive: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Float32Ptr(90),
										},
										TierToCool: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Float32Ptr(30),
										},
									},
								},
								Filters: &armstorage.ManagementPolicyFilter{
									BlobIndexMatch: []*armstorage.TagFilter{
										{
											Name:  to.StringPtr("<name>"),
											Op:    to.StringPtr("<op>"),
											Value: to.StringPtr("<value>"),
										},
										{
											Name:  to.StringPtr("<name>"),
											Op:    to.StringPtr("<op>"),
											Value: to.StringPtr("<value>"),
										}},
									BlobTypes: []*string{
										to.StringPtr("blockBlob")},
									PrefixMatch: []*string{
										to.StringPtr("olcmtestcontainer2")},
								},
							},
							Enabled: to.BoolPtr(true),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ManagementPoliciesClientCreateOrUpdateResult)
}
```
