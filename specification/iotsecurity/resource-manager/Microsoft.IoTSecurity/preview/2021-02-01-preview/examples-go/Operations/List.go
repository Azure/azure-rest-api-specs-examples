package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/Operations/List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotsecurity.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationList = armiotsecurity.OperationList{
		// 	Value: []*armiotsecurity.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTSecurity/unregister/action"),
		// 			Display: &armiotsecurity.OperationDisplay{
		// 				Description: to.Ptr("Unregisters the subscription for Azure Defender for IoT"),
		// 				Operation: to.Ptr("Unregister Subscription"),
		// 				Provider: to.Ptr("Microsoft IoT Security"),
		// 				Resource: to.Ptr("Subscription"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTSecurity/register/action"),
		// 			Display: &armiotsecurity.OperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for Azure Defender for IoT"),
		// 				Operation: to.Ptr("Register Subscription"),
		// 				Provider: to.Ptr("Microsoft IoT Security"),
		// 				Resource: to.Ptr("Subscription"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTSecurity/defenderSettings/read"),
		// 			Display: &armiotsecurity.OperationDisplay{
		// 				Description: to.Ptr("Gets IoT Defender Settings"),
		// 				Operation: to.Ptr("Get IoT Defender Settings"),
		// 				Provider: to.Ptr("Microsoft IoT Security"),
		// 				Resource: to.Ptr("IoT Defender Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTSecurity/defenderSettings/write"),
		// 			Display: &armiotsecurity.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates IoT Defender Settings"),
		// 				Operation: to.Ptr("Create or update IoT Defender Settings"),
		// 				Provider: to.Ptr("Microsoft IoT Security"),
		// 				Resource: to.Ptr("IoT Defender Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTSecurity/defenderSettings/delete"),
		// 			Display: &armiotsecurity.OperationDisplay{
		// 				Description: to.Ptr("Deletes IoT Defender Settings"),
		// 				Operation: to.Ptr("Delete IoT Defender Settings"),
		// 				Provider: to.Ptr("Microsoft IoT Security"),
		// 				Resource: to.Ptr("IoT Defender Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTSecurity/defenderSettings/packageDownloads/action"),
		// 			Display: &armiotsecurity.OperationDisplay{
		// 				Description: to.Ptr("Gets downloadable IoT Defender packages information"),
		// 				Operation: to.Ptr("Get downloadable IoT Defender packages information"),
		// 				Provider: to.Ptr("Microsoft IoT Security"),
		// 				Resource: to.Ptr("IoT Defender Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.IoTSecurity/defenderSettings/downloadManagerActivation/action"),
		// 			Display: &armiotsecurity.OperationDisplay{
		// 				Description: to.Ptr("Download manager activation file"),
		// 				Operation: to.Ptr("Download manager activation file"),
		// 				Provider: to.Ptr("Microsoft IoT Security"),
		// 				Resource: to.Ptr("IoT Defender Settings"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 	}},
		// }
	}
}
