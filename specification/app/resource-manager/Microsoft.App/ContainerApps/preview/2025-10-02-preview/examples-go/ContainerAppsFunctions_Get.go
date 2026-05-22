package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/ContainerAppsFunctions_Get.json
func ExampleContainerAppsFunctionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("12345678-1234-1234-1234-123456789abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsFunctionsClient().Get(ctx, "myResourceGroup", "myContainerApp", "HttpExample", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.ContainerAppsFunctionsClientGetResponse{
	// 	ContainerAppsFunction: armappcontainers.ContainerAppsFunction{
	// 		Name: to.Ptr("myContainerApp/HttpExample"),
	// 		Type: to.Ptr("Microsoft.App/containerApps/functions"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789abc/resourceGroups/myResourceGroup/providers/Microsoft.App/containerApps/myContainerApp/functions/HttpExample"),
	// 		Properties: &armappcontainers.ContainerAppsFunctionProperties{
	// 			InvokeURLTemplate: to.Ptr("https://mycontainerapp.example-region.azurecontainerapps.io/api/HttpExample"),
	// 			IsDisabled: to.Ptr(false),
	// 			TriggerType: to.Ptr("httpTrigger"),
	// 			Language: to.Ptr("dotnet-isolated"),
	// 		},
	// 	},
	// }
}
