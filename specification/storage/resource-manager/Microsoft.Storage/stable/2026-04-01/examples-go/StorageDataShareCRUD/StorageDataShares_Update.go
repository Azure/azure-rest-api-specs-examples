package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-04-01/StorageDataShareCRUD/StorageDataShares_Update.json
func ExampleDataSharesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataSharesClient().BeginUpdate(ctx, "testrg", "teststorageaccount", "testdatashare", armstorage.DataShareUpdate{
		Properties: &armstorage.DataSharePropertiesUpdate{
			Description: to.Ptr("New dummy data share"),
			AccessPolicies: []*armstorage.DataShareAccessPolicy{
				{
					PrincipalID: to.Ptr("00000000-0000-0000-0000-123456781234"),
					TenantID:    to.Ptr("00000000-0000-0000-0000-987654321987"),
					Permission:  to.Ptr(armstorage.StorageDataShareAccessPolicyPermissionRead),
				},
			},
			Assets: []*armstorage.DataShareAsset{
				{
					AssetPath:   to.Ptr("/container/folder1/bar"),
					DisplayName: to.Ptr("virtualBar"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.DataSharesClientUpdateResponse{
	// 	DataShare: armstorage.DataShare{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount/dataShares/testdatashare"),
	// 		Name: to.Ptr("testdatashare"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/dataShares"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armstorage.DataShareProperties{
	// 			DataShareIdentifier: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			Description: to.Ptr("New dummy data share"),
	// 			DataShareURI: to.Ptr("azds://eastus:testdatashare:00000000-0000-0000-0000-000000000000"),
	// 			AccessPolicies: []*armstorage.DataShareAccessPolicy{
	// 				{
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-123456781234"),
	// 					TenantID: to.Ptr("00000000-0000-0000-0000-987654321987"),
	// 					Permission: to.Ptr(armstorage.StorageDataShareAccessPolicyPermissionRead),
	// 				},
	// 			},
	// 			Assets: []*armstorage.DataShareAsset{
	// 				{
	// 					AssetPath: to.Ptr("/container/folder1/bar"),
	// 					DisplayName: to.Ptr("virtualBar"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armstorage.NativeDataSharingProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
