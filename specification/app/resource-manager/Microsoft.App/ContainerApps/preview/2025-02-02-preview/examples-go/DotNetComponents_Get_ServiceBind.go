package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/DotNetComponents_Get_ServiceBind.json
func ExampleDotNetComponentsClient_Get_getNetComponentWithServiceBinds() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDotNetComponentsClient().Get(ctx, "examplerg", "myenvironment", "mydotnetcomponent", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DotNetComponent = armappcontainers.DotNetComponent{
	// 	Name: to.Ptr("mydotnetcomponent"),
	// 	Type: to.Ptr("Microsoft.App/managedEnvironments/dotNetComponents"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/dotNetComponents/mydotnetcomponent"),
	// 	Properties: &armappcontainers.DotNetComponentProperties{
	// 		ComponentType: to.Ptr(armappcontainers.DotNetComponentTypeAspireDashboard),
	// 		Configurations: []*armappcontainers.DotNetComponentConfigurationProperty{
	// 			{
	// 				PropertyName: to.Ptr("dashboard-theme"),
	// 				Value: to.Ptr("dark"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armappcontainers.DotNetComponentProvisioningStateSucceeded),
	// 		ServiceBinds: []*armappcontainers.DotNetComponentServiceBind{
	// 			{
	// 				Name: to.Ptr("yellowcat"),
	// 				ServiceID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/dotNetComponents/yellowcat"),
	// 		}},
	// 	},
	// }
}
