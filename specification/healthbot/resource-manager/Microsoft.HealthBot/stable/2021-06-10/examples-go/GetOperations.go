package armhealthbot_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2021-06-10/examples/GetOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthbot.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.AvailableOperations = armhealthbot.AvailableOperations{
		// 	Value: []*armhealthbot.OperationDetail{
		// 		{
		// 			Name: to.Ptr("Microsoft.Healthbot/healthbots/read"),
		// 			Display: &armhealthbot.OperationDisplay{
		// 				Description: to.Ptr("Read Azure Health Bot"),
		// 				Operation: to.Ptr("Read Azure Health Bot"),
		// 				Provider: to.Ptr("Azure Health Bot"),
		// 				Resource: to.Ptr("Azure Health Bot"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Healthbot/healthbots/write"),
		// 			Display: &armhealthbot.OperationDisplay{
		// 				Description: to.Ptr("Writes Azure Health Bot"),
		// 				Operation: to.Ptr("Write Azure Health Bot"),
		// 				Provider: to.Ptr("Azure Health Bot"),
		// 				Resource: to.Ptr("Azure Health Bot"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Healthbot/healthbots/delete"),
		// 			Display: &armhealthbot.OperationDisplay{
		// 				Description: to.Ptr("Deletes Azure Health Bot"),
		// 				Operation: to.Ptr("Delete Azure Health Bot"),
		// 				Provider: to.Ptr("Azure Health Bot"),
		// 				Resource: to.Ptr("Azure Health Bot"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
