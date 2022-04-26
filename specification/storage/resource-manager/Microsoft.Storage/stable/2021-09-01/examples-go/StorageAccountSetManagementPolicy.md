Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.6.0/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.

```go
package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountSetManagementPolicy.json
func ExampleManagementPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armstorage.NewManagementPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.ManagementPolicyNameDefault,
		armstorage.ManagementPolicy{
			Properties: &armstorage.ManagementPolicyProperties{
				Policy: &armstorage.ManagementPolicySchema{
					Rules: []*armstorage.ManagementPolicyRule{
						{
							Name: to.Ptr("<name>"),
							Type: to.Ptr(armstorage.RuleTypeLifecycle),
							Definition: &armstorage.ManagementPolicyDefinition{
								Actions: &armstorage.ManagementPolicyAction{
									BaseBlob: &armstorage.ManagementPolicyBaseBlob{
										Delete: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Ptr[float32](1000),
										},
										TierToArchive: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Ptr[float32](90),
										},
										TierToCool: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Ptr[float32](30),
										},
									},
									Snapshot: &armstorage.ManagementPolicySnapShot{
										Delete: &armstorage.DateAfterCreation{
											DaysAfterCreationGreaterThan: to.Ptr[float32](30),
										},
									},
								},
								Filters: &armstorage.ManagementPolicyFilter{
									BlobTypes: []*string{
										to.Ptr("blockBlob")},
									PrefixMatch: []*string{
										to.Ptr("olcmtestcontainer1")},
								},
							},
							Enabled: to.Ptr(true),
						},
						{
							Name: to.Ptr("<name>"),
							Type: to.Ptr(armstorage.RuleTypeLifecycle),
							Definition: &armstorage.ManagementPolicyDefinition{
								Actions: &armstorage.ManagementPolicyAction{
									BaseBlob: &armstorage.ManagementPolicyBaseBlob{
										Delete: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Ptr[float32](1000),
										},
										TierToArchive: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Ptr[float32](90),
										},
										TierToCool: &armstorage.DateAfterModification{
											DaysAfterModificationGreaterThan: to.Ptr[float32](30),
										},
									},
								},
								Filters: &armstorage.ManagementPolicyFilter{
									BlobIndexMatch: []*armstorage.TagFilter{
										{
											Name:  to.Ptr("<name>"),
											Op:    to.Ptr("<op>"),
											Value: to.Ptr("<value>"),
										},
										{
											Name:  to.Ptr("<name>"),
											Op:    to.Ptr("<op>"),
											Value: to.Ptr("<value>"),
										}},
									BlobTypes: []*string{
										to.Ptr("blockBlob")},
									PrefixMatch: []*string{
										to.Ptr("olcmtestcontainer2")},
								},
							},
							Enabled: to.Ptr(true),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
