package armstorageactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageactions/armstorageactions"
)

// Generated from example definition: 2023-01-01/storageTasksList/ListStorageTasksByResourceGroup.json
func ExampleStorageTasksClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorageactions.NewClientFactory("1f31ba14-ce16-4281-b9b4-3e78da6e1616", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageTasksClient().NewListByResourceGroupPager("res6117", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armstorageactions.StorageTasksClientListByResourceGroupResponse{
		// 	StorageTasksListResult: armstorageactions.StorageTasksListResult{
		// 		Value: []*armstorageactions.StorageTask{
		// 			{
		// 				Name: to.Ptr("mytask1"),
		// 				Type: to.Ptr("Microsoft.StorageActions/storageTasks"),
		// 				ID: to.Ptr("/subscriptions/c86a9c18-8373-41fa-92d4-1d7bdc16977b/resourceGroups/res6117/providers/Microsoft.StorageActions/storageTasks/mytask1"),
		// 				Identity: &armstorageactions.ManagedServiceIdentity{
		// 					Type: to.Ptr(armstorageactions.ManagedServiceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("2fd475e8-8923-4597-842f-7ce1adfc6c4a"),
		// 					TenantID: to.Ptr("b4a2005c-32c1-434c-bbf0-ff486912fc75"),
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armstorageactions.StorageTaskProperties{
		// 					Description: to.Ptr("Storage task"),
		// 					Action: &armstorageactions.StorageTaskAction{
		// 						If: &armstorageactions.IfCondition{
		// 							Condition: to.Ptr("[[greater(Content-Length, '100')]]"),
		// 							Operations: []*armstorageactions.StorageTaskOperation{
		// 								{
		// 									Name: to.Ptr(armstorageactions.StorageTaskOperationNameDeleteBlob),
		// 									OnFailure: to.Ptr(armstorageactions.OnFailureBreak),
		// 									OnSuccess: to.Ptr(armstorageactions.OnSuccessContinue),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					CreationTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-05T23:06:50.7722358Z"); return t}()),
		// 					Enabled: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armstorageactions.ProvisioningStateSucceeded),
		// 					TaskVersion: to.Ptr[int64](1),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("mytask2"),
		// 				Type: to.Ptr("Microsoft.StorageActions/storageTasks"),
		// 				ID: to.Ptr("/subscriptions/c86a9c18-8373-41fa-92d4-1d7bdc16977b/resourceGroups/res6117/providers/Microsoft.StorageActions/storageTasks/mytask2"),
		// 				Identity: &armstorageactions.ManagedServiceIdentity{
		// 					Type: to.Ptr(armstorageactions.ManagedServiceIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armstorageactions.UserAssignedIdentity{
		// 						"/subscriptions/c86a9c18-8373-41fa-92d4-1d7bdc16977b/resourceGroups/res6117/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myUserAssignedIdentity": &armstorageactions.UserAssignedIdentity{
		// 							ClientID: to.Ptr("bbbbbbbb-0000-0000-0000-000000000000"),
		// 							PrincipalID: to.Ptr("aaaaaaaa-0000-0000-0000-000000000000"),
		// 						},
		// 					},
		// 				},
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armstorageactions.StorageTaskProperties{
		// 					Description: to.Ptr("Storage task"),
		// 					Action: &armstorageactions.StorageTaskAction{
		// 						Else: &armstorageactions.ElseCondition{
		// 							Operations: []*armstorageactions.StorageTaskOperation{
		// 								{
		// 									Name: to.Ptr(armstorageactions.StorageTaskOperationNameDeleteBlob),
		// 									OnFailure: to.Ptr(armstorageactions.OnFailureBreak),
		// 									OnSuccess: to.Ptr(armstorageactions.OnSuccessContinue),
		// 								},
		// 							},
		// 						},
		// 						If: &armstorageactions.IfCondition{
		// 							Condition: to.Ptr("[[equals(AccessTier, 'Cool')]]"),
		// 							Operations: []*armstorageactions.StorageTaskOperation{
		// 								{
		// 									Name: to.Ptr(armstorageactions.StorageTaskOperationNameSetBlobTier),
		// 									OnFailure: to.Ptr(armstorageactions.OnFailureBreak),
		// 									OnSuccess: to.Ptr(armstorageactions.OnSuccessContinue),
		// 									Parameters: map[string]*string{
		// 										"tier": to.Ptr("Hot"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					CreationTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-06T00:24:59.1441663Z"); return t}()),
		// 					Enabled: to.Ptr(true),
		// 					ProvisioningState: to.Ptr(armstorageactions.ProvisioningStateSucceeded),
		// 					TaskVersion: to.Ptr[int64](1),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
