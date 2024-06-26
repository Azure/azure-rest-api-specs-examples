package armwindowsiot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsiot/armwindowsiot"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/windowsiot/resource-manager/Microsoft.WindowsIoT/stable/2019-06-01/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwindowsiot.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationListResult = armwindowsiot.OperationListResult{
		// 	Value: []*armwindowsiot.OperationEntity{
		// 		{
		// 			Name: to.Ptr("Microsoft.WindowsIoT/Services/write"),
		// 			Display: &armwindowsiot.OperationDisplayInfo{
		// 				Description: to.Ptr("Creates a Windows IoT Subscription with the specified parameters or update the properties or tags or adds custom domain for the specified Windows IoT Subscription."),
		// 				Operation: to.Ptr("Create/Update Windows IoT Subscription"),
		// 				Provider: to.Ptr("Windows IoT"),
		// 				Resource: to.Ptr("Windows IoT Services"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.WindowsIoT/Services/delete"),
		// 			Display: &armwindowsiot.OperationDisplayInfo{
		// 				Description: to.Ptr("Deletes an existing Windows IoT Subscription."),
		// 				Operation: to.Ptr("Delete Windows IoT Subscription"),
		// 				Provider: to.Ptr("Windows IoT"),
		// 				Resource: to.Ptr("Windows IoT Services"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.WindowsIoT/checkNameAvailability/read"),
		// 			Display: &armwindowsiot.OperationDisplayInfo{
		// 				Description: to.Ptr("Checks that account name is valid and is not in use."),
		// 				Operation: to.Ptr("Check Name Availability"),
		// 				Provider: to.Ptr("Windows IoT"),
		// 				Resource: to.Ptr("Name Availability"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.WindowsIoT/Services/read"),
		// 			Display: &armwindowsiot.OperationDisplayInfo{
		// 				Description: to.Ptr("Returns the list of Windows IoT Services or gets the properties for the specified Windows IoT Subscription."),
		// 				Operation: to.Ptr("List/Get Windows IoT Subscription(s)"),
		// 				Provider: to.Ptr("Windows IoT"),
		// 				Resource: to.Ptr("Windows IoT Services"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.WindowsIoT/operations/read"),
		// 			Display: &armwindowsiot.OperationDisplayInfo{
		// 				Description: to.Ptr("Polls the status of an asynchronous operation."),
		// 				Operation: to.Ptr("Poll Asynchronous Operation"),
		// 				Provider: to.Ptr("Windows IoT"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
