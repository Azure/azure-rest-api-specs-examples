package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/GetOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationEntityListResult = armbotservice.OperationEntityListResult{
		// 	Value: []*armbotservice.OperationEntity{
		// 		{
		// 			Name: to.Ptr("Microsoft.BotService/botService/read"),
		// 			Display: &armbotservice.OperationDisplayInfo{
		// 				Description: to.Ptr("Read Bot Service"),
		// 				Operation: to.Ptr("Read Bot Service"),
		// 				Provider: to.Ptr("Microsoft Bot Service"),
		// 				Resource: to.Ptr("Bot Service"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.BotService/botService/write"),
		// 			Display: &armbotservice.OperationDisplayInfo{
		// 				Description: to.Ptr("Writes Bot Service"),
		// 				Operation: to.Ptr("Write Bot Service"),
		// 				Provider: to.Ptr("Microsoft Bot Service"),
		// 				Resource: to.Ptr("Bot Service"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.BotService/botService/delete"),
		// 			Display: &armbotservice.OperationDisplayInfo{
		// 				Description: to.Ptr("Deletes Bot Service"),
		// 				Operation: to.Ptr("Delete Bot Service"),
		// 				Provider: to.Ptr("Microsoft Bot Service"),
		// 				Resource: to.Ptr("Bot Service"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.BotService/botService/botName/read"),
		// 			Display: &armbotservice.OperationDisplayInfo{
		// 				Description: to.Ptr("Check bot name availability"),
		// 				Operation: to.Ptr("Check bot name availability"),
		// 				Provider: to.Ptr("Microsoft Bot Service"),
		// 				Resource: to.Ptr("Bot Service name availability"),
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
