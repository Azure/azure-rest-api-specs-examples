package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountSetManagementPolicy_LastAccessTimeBasedBlobActions.json
func ExampleManagementPoliciesClient_CreateOrUpdate_storageAccountSetManagementPolicy_LastAccessTimeBasedBlobActions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewManagementPoliciesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "res7687", "sto9699", armstorage.ManagementPolicyNameDefault, armstorage.ManagementPolicy{
		Properties: &armstorage.ManagementPolicyProperties{
			Policy: &armstorage.ManagementPolicySchema{
				Rules: []*armstorage.ManagementPolicyRule{
					{
						Name: to.Ptr("olcmtest"),
						Type: to.Ptr(armstorage.RuleTypeLifecycle),
						Definition: &armstorage.ManagementPolicyDefinition{
							Actions: &armstorage.ManagementPolicyAction{
								BaseBlob: &armstorage.ManagementPolicyBaseBlob{
									Delete: &armstorage.DateAfterModification{
										DaysAfterLastAccessTimeGreaterThan: to.Ptr[float32](1000),
									},
									EnableAutoTierToHotFromCool: to.Ptr(true),
									TierToArchive: &armstorage.DateAfterModification{
										DaysAfterLastAccessTimeGreaterThan: to.Ptr[float32](90),
									},
									TierToCool: &armstorage.DateAfterModification{
										DaysAfterLastAccessTimeGreaterThan: to.Ptr[float32](30),
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
									to.Ptr("olcmtestcontainer")},
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
	// TODO: use response item
	_ = res
}
