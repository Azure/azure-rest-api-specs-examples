package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/ContainerAppsFunctions_List.json
func ExampleContainerAppsFunctionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("12345678-1234-1234-1234-123456789abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsFunctionsClient().NewListPager("myResourceGroup", "myContainerApp", nil)
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
		// page = armappcontainers.ContainerAppsFunctionsClientListResponse{
		// 	ContainerAppsFunctionCollection: armappcontainers.ContainerAppsFunctionCollection{
		// 		Value: []*armappcontainers.ContainerAppsFunction{
		// 			{
		// 				Name: to.Ptr("myContainerApp/HttpExample"),
		// 				Type: to.Ptr("Microsoft.App/containerApps/functions"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789abc/resourceGroups/myResourceGroup/providers/Microsoft.App/containerApps/myContainerApp/functions/HttpExample"),
		// 				Properties: &armappcontainers.ContainerAppsFunctionProperties{
		// 					InvokeURLTemplate: to.Ptr("https://mycontainerapp.example-region.azurecontainerapps.io/api/HttpExample"),
		// 					IsDisabled: to.Ptr(false),
		// 					TriggerType: to.Ptr("httpTrigger"),
		// 					Language: to.Ptr("dotnet-isolated"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myContainerApp/TimerFunction"),
		// 				Type: to.Ptr("Microsoft.App/containerApps/functions"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789abc/resourceGroups/myResourceGroup/providers/Microsoft.App/containerApps/myContainerApp/functions/TimerFunction"),
		// 				Properties: &armappcontainers.ContainerAppsFunctionProperties{
		// 					InvokeURLTemplate: to.Ptr("https://mycontainerapp.example-region.azurecontainerapps.io/api/TimerFunction"),
		// 					IsDisabled: to.Ptr(false),
		// 					TriggerType: to.Ptr("timerTrigger"),
		// 					Language: to.Ptr("dotnet-isolated"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
