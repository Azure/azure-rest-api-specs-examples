package armcustomerlockbox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerlockbox/armcustomerlockbox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/customerlockbox/resource-manager/Microsoft.CustomerLockbox/preview/2018-02-28-preview/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerlockbox.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armcustomerlockbox.OperationListResult{
		// 	Value: []*armcustomerlockbox.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerLockbox/register/action"),
		// 			Display: &armcustomerlockbox.OperationDisplay{
		// 				Description: to.Ptr("Register Provider Microsoft.CustomerLockbox"),
		// 				Operation: to.Ptr("Register Provider Microsoft.CustomerLockboxx"),
		// 				Provider: to.Ptr("Microsoft Customer Lockbox"),
		// 				Resource: to.Ptr("Microsoft Customer Lockbox"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerLockbox/operations/read"),
		// 			Display: &armcustomerlockbox.OperationDisplay{
		// 				Description: to.Ptr("Read Lockbox Operations"),
		// 				Operation: to.Ptr("Read Lockbox Operations"),
		// 				Provider: to.Ptr("Microsoft Customer Lockbox"),
		// 				Resource: to.Ptr("Lockbox Operations"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerLockbox/requests/read"),
		// 			Display: &armcustomerlockbox.OperationDisplay{
		// 				Description: to.Ptr("Read Lockbox Request"),
		// 				Operation: to.Ptr("Read Lockbox Request"),
		// 				Provider: to.Ptr("Microsoft Customer Lockbox"),
		// 				Resource: to.Ptr("Customer Lockbox Request"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerLockbox/requests/activitylog/CreateLockboxRequest/action"),
		// 			Display: &armcustomerlockbox.OperationDisplay{
		// 				Description: to.Ptr("Create Lockbox Request"),
		// 				Operation: to.Ptr("Create Lockbox Request"),
		// 				Provider: to.Ptr("Microsoft Customer Lockbox"),
		// 				Resource: to.Ptr("Customer Lockbox Request"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerLockbox/requests/activitylog/ApproveLockboxRequest/action"),
		// 			Display: &armcustomerlockbox.OperationDisplay{
		// 				Description: to.Ptr("Approve Lockbox Request"),
		// 				Operation: to.Ptr("Approve Lockbox Request"),
		// 				Provider: to.Ptr("Microsoft Customer Lockbox"),
		// 				Resource: to.Ptr("Customer Lockbox Request"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerLockbox/requests/activitylog/DenyLockboxRequest/action"),
		// 			Display: &armcustomerlockbox.OperationDisplay{
		// 				Description: to.Ptr("Deny Lockbox Request"),
		// 				Operation: to.Ptr("Deny Lockbox Request"),
		// 				Provider: to.Ptr("Microsoft Customer Lockbox"),
		// 				Resource: to.Ptr("Customer Lockbox Request"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerLockbox/requests/activitylog/ExpireLockboxRequest/action"),
		// 			Display: &armcustomerlockbox.OperationDisplay{
		// 				Description: to.Ptr("Expire Lockbox Request"),
		// 				Operation: to.Ptr("Expire Lockbox Request"),
		// 				Provider: to.Ptr("Microsoft Customer Lockbox"),
		// 				Resource: to.Ptr("Customer Lockbox Request"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerLockbox/requests/activitylog/CancelLockboxRequest/action"),
		// 			Display: &armcustomerlockbox.OperationDisplay{
		// 				Description: to.Ptr("Cancel Lockbox Request"),
		// 				Operation: to.Ptr("Cancel Lockbox Request"),
		// 				Provider: to.Ptr("Microsoft Customer Lockbox"),
		// 				Resource: to.Ptr("Customer Lockbox Request"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerLockbox/requests/activitylog/AutoApproveLockboxRequest/action"),
		// 			Display: &armcustomerlockbox.OperationDisplay{
		// 				Description: to.Ptr("AutoApprove Lockbox Request"),
		// 				Operation: to.Ptr("AutoApprove Lockbox Request"),
		// 				Provider: to.Ptr("Microsoft Customer Lockbox"),
		// 				Resource: to.Ptr("Customer Lockbox Request"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerLockbox/requests/activitylog/AutoDenyLockboxRequest/action"),
		// 			Display: &armcustomerlockbox.OperationDisplay{
		// 				Description: to.Ptr("AutoDeny Lockbox Request"),
		// 				Operation: to.Ptr("AutoDeny Lockbox Request"),
		// 				Provider: to.Ptr("Microsoft Customer Lockbox"),
		// 				Resource: to.Ptr("Customer Lockbox Request"),
		// 			},
		// 			IsDataAction: to.Ptr("false"),
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
