package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80065490402157d0df0dd37ab347c651b22eb576/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armdeviceupdate.OperationListResult{
		// 	Value: []*armdeviceupdate.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.DeviceUpdate/accounts/read"),
		// 			Display: &armdeviceupdate.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of Device Update Accounts"),
		// 				Operation: to.Ptr("Get/List Accounts"),
		// 				Provider: to.Ptr("Microsoft.DeviceUpdate"),
		// 				Resource: to.Ptr("Device Update Account"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DeviceUpdate/accounts/write"),
		// 			Display: &armdeviceupdate.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates a Device Update Account"),
		// 				Operation: to.Ptr("Create/Update Account"),
		// 				Provider: to.Ptr("Microsoft.DeviceUpdate"),
		// 				Resource: to.Ptr("Device Update Account"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
