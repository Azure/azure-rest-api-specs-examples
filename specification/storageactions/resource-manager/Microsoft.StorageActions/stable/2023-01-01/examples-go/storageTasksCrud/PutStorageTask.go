package armstorageactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageactions/armstorageactions"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5b798125a6aa7d5152fe0e3dd595d8a76dcfa568/specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/storageTasksCrud/PutStorageTask.json
func ExampleStorageTasksClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorageactions.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageTasksClient().BeginCreate(ctx, "res4228", "mytask1", armstorageactions.StorageTask{
		Location: to.Ptr("westus"),
		Identity: &armstorageactions.ManagedServiceIdentity{
			Type: to.Ptr(armstorageactions.ManagedServiceIdentityTypeSystemAssigned),
		},
		Properties: &armstorageactions.StorageTaskProperties{
			Description: to.Ptr("My Storage task"),
			Action: &armstorageactions.StorageTaskAction{
				Else: &armstorageactions.ElseCondition{
					Operations: []*armstorageactions.StorageTaskOperation{
						{
							Name:      to.Ptr(armstorageactions.StorageTaskOperationNameDeleteBlob),
							OnFailure: to.Ptr(armstorageactions.OnFailureBreak),
							OnSuccess: to.Ptr(armstorageactions.OnSuccessContinue),
						}},
				},
				If: &armstorageactions.IfCondition{
					Condition: to.Ptr("[[equals(AccessTier, 'Cool')]]"),
					Operations: []*armstorageactions.StorageTaskOperation{
						{
							Name:      to.Ptr(armstorageactions.StorageTaskOperationNameSetBlobTier),
							OnFailure: to.Ptr(armstorageactions.OnFailureBreak),
							OnSuccess: to.Ptr(armstorageactions.OnSuccessContinue),
							Parameters: map[string]*string{
								"tier": to.Ptr("Hot"),
							},
						}},
				},
			},
			Enabled: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageTask = armstorageactions.StorageTask{
	// 	Name: to.Ptr("mytask1"),
	// 	Type: to.Ptr("Microsoft.StorageActions/storageTasks"),
	// 	ID: to.Ptr("/subscriptions/c86a9c18-8373-41fa-92d4-1d7bdc16977b/resourceGroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask1"),
	// 	Location: to.Ptr("westus"),
	// 	Identity: &armstorageactions.ManagedServiceIdentity{
	// 		Type: to.Ptr(armstorageactions.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("2fd475e8-8923-4597-842f-7ce1adfc6c4a"),
	// 		TenantID: to.Ptr("b4a2005c-32c1-434c-bbf0-ff486912fc75"),
	// 	},
	// 	Properties: &armstorageactions.StorageTaskProperties{
	// 		Description: to.Ptr("Storage task"),
	// 		Action: &armstorageactions.StorageTaskAction{
	// 			Else: &armstorageactions.ElseCondition{
	// 				Operations: []*armstorageactions.StorageTaskOperation{
	// 					{
	// 						Name: to.Ptr(armstorageactions.StorageTaskOperationNameDeleteBlob),
	// 						OnFailure: to.Ptr(armstorageactions.OnFailureBreak),
	// 						OnSuccess: to.Ptr(armstorageactions.OnSuccessContinue),
	// 				}},
	// 			},
	// 			If: &armstorageactions.IfCondition{
	// 				Condition: to.Ptr("[[equals(AccessTier, 'Cool')]]"),
	// 				Operations: []*armstorageactions.StorageTaskOperation{
	// 					{
	// 						Name: to.Ptr(armstorageactions.StorageTaskOperationNameSetBlobTier),
	// 						OnFailure: to.Ptr(armstorageactions.OnFailureBreak),
	// 						OnSuccess: to.Ptr(armstorageactions.OnSuccessContinue),
	// 						Parameters: map[string]*string{
	// 							"tier": to.Ptr("Hot"),
	// 						},
	// 				}},
	// 			},
	// 		},
	// 		CreationTimeInUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-06T00:24:59.144Z"); return t}()),
	// 		Enabled: to.Ptr(true),
	// 		ProvisioningState: to.Ptr(armstorageactions.ProvisioningStateSucceeded),
	// 		TaskVersion: to.Ptr[int64](1),
	// 	},
	// }
}
