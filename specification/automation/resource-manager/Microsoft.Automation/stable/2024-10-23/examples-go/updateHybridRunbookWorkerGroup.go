package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/updateHybridRunbookWorkerGroup.json
func ExampleHybridRunbookWorkerGroupClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHybridRunbookWorkerGroupClient().Update(ctx, "rg", "testaccount", "TestHybridGroup", armautomation.HybridRunbookWorkerGroupCreateOrUpdateParameters{
		Properties: &armautomation.HybridRunbookWorkerGroupCreateOrUpdateProperties{
			Credential: &armautomation.RunAsCredentialAssociationProperty{
				Name: to.Ptr("myRunAsCredentialUpdatedName"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.HybridRunbookWorkerGroupClientUpdateResponse{
	// 	HybridRunbookWorkerGroup: armautomation.HybridRunbookWorkerGroup{
	// 		Name: to.Ptr("TestHybridGroup"),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/testaccount/hybridRunbookWorkerGroups/TestHybridGroup"),
	// 		Location: to.Ptr("East US 2"),
	// 		Properties: &armautomation.HybridRunbookWorkerGroupProperties{
	// 			Credential: &armautomation.RunAsCredentialAssociationProperty{
	// 				Name: to.Ptr("myRunAsCredentialUpdatedName"),
	// 			},
	// 			GroupType: to.Ptr(armautomation.GroupTypeEnumUser),
	// 		},
	// 		SystemData: &armautomation.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 			CreatedBy: to.Ptr("foo@contoso.com"),
	// 			CreatedByType: to.Ptr(armautomation.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-24T16:30:55+00:00"); return t}()),
	// 			LastModifiedBy: to.Ptr("foo@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armautomation.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
