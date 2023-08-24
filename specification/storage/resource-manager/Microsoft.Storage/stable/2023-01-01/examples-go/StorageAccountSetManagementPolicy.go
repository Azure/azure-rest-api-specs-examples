package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountSetManagementPolicy.json
func ExampleManagementPoliciesClient_CreateOrUpdate_storageAccountSetManagementPolicies() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementPoliciesClient().CreateOrUpdate(ctx, "res7687", "sto9699", armstorage.ManagementPolicyNameDefault, armstorage.ManagementPolicy{
		Properties: &armstorage.ManagementPolicyProperties{
			Policy: &armstorage.ManagementPolicySchema{
				Rules: []*armstorage.ManagementPolicyRule{
					{
						Name: to.Ptr("olcmtest1"),
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
						Name: to.Ptr("olcmtest2"),
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
										Name:  to.Ptr("tag1"),
										Op:    to.Ptr("=="),
										Value: to.Ptr("val1"),
									},
									{
										Name:  to.Ptr("tag2"),
										Op:    to.Ptr("=="),
										Value: to.Ptr("val2"),
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementPolicy = armstorage.ManagementPolicy{
	// 	Name: to.Ptr("DefaultManagementPolicy"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/managementPolicies"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res7231/providers/Microsoft.Storage/storageAccounts/sto288/managementPolicies/default"),
	// 	Properties: &armstorage.ManagementPolicyProperties{
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-08T02:53:39.0932539Z"); return t}()),
	// 		Policy: &armstorage.ManagementPolicySchema{
	// 			Rules: []*armstorage.ManagementPolicyRule{
	// 				{
	// 					Name: to.Ptr("olcmtest1"),
	// 					Type: to.Ptr(armstorage.RuleTypeLifecycle),
	// 					Definition: &armstorage.ManagementPolicyDefinition{
	// 						Actions: &armstorage.ManagementPolicyAction{
	// 							BaseBlob: &armstorage.ManagementPolicyBaseBlob{
	// 								Delete: &armstorage.DateAfterModification{
	// 									DaysAfterModificationGreaterThan: to.Ptr[float32](1000),
	// 								},
	// 								TierToArchive: &armstorage.DateAfterModification{
	// 									DaysAfterModificationGreaterThan: to.Ptr[float32](90),
	// 								},
	// 								TierToCool: &armstorage.DateAfterModification{
	// 									DaysAfterModificationGreaterThan: to.Ptr[float32](30),
	// 								},
	// 							},
	// 							Snapshot: &armstorage.ManagementPolicySnapShot{
	// 								Delete: &armstorage.DateAfterCreation{
	// 									DaysAfterCreationGreaterThan: to.Ptr[float32](30),
	// 								},
	// 							},
	// 						},
	// 						Filters: &armstorage.ManagementPolicyFilter{
	// 							BlobTypes: []*string{
	// 								to.Ptr("blockBlob")},
	// 								PrefixMatch: []*string{
	// 									to.Ptr("olcmtestcontainer1")},
	// 								},
	// 							},
	// 							Enabled: to.Ptr(true),
	// 						},
	// 						{
	// 							Name: to.Ptr("olcmtest2"),
	// 							Type: to.Ptr(armstorage.RuleTypeLifecycle),
	// 							Definition: &armstorage.ManagementPolicyDefinition{
	// 								Actions: &armstorage.ManagementPolicyAction{
	// 									BaseBlob: &armstorage.ManagementPolicyBaseBlob{
	// 										Delete: &armstorage.DateAfterModification{
	// 											DaysAfterModificationGreaterThan: to.Ptr[float32](1000),
	// 										},
	// 										TierToArchive: &armstorage.DateAfterModification{
	// 											DaysAfterModificationGreaterThan: to.Ptr[float32](90),
	// 										},
	// 										TierToCool: &armstorage.DateAfterModification{
	// 											DaysAfterModificationGreaterThan: to.Ptr[float32](30),
	// 										},
	// 									},
	// 								},
	// 								Filters: &armstorage.ManagementPolicyFilter{
	// 									BlobIndexMatch: []*armstorage.TagFilter{
	// 										{
	// 											Name: to.Ptr("tag1"),
	// 											Op: to.Ptr("=="),
	// 											Value: to.Ptr("val1"),
	// 										},
	// 										{
	// 											Name: to.Ptr("tag2"),
	// 											Op: to.Ptr("=="),
	// 											Value: to.Ptr("val2"),
	// 									}},
	// 									BlobTypes: []*string{
	// 										to.Ptr("blockBlob")},
	// 										PrefixMatch: []*string{
	// 											to.Ptr("olcmtestcontainer2")},
	// 										},
	// 									},
	// 									Enabled: to.Ptr(true),
	// 							}},
	// 						},
	// 					},
	// 				}
}
