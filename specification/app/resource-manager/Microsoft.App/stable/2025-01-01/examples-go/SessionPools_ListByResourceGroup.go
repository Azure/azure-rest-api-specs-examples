package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/SessionPools_ListByResourceGroup.json
func ExampleContainerAppsSessionPoolsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsSessionPoolsClient().NewListByResourceGroupPager("rg", nil)
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
		// page.SessionPoolCollection = armappcontainers.SessionPoolCollection{
		// 	Value: []*armappcontainers.SessionPool{
		// 		{
		// 			Name: to.Ptr("testsessionpool"),
		// 			Type: to.Ptr("Microsoft.App/sessionPools"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/sessionPools/testsessionpool"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.SessionPoolProperties{
		// 				ContainerType: to.Ptr(armappcontainers.ContainerTypeCustomContainer),
		// 				CustomContainerTemplate: &armappcontainers.CustomContainerTemplate{
		// 					Containers: []*armappcontainers.SessionContainer{
		// 						{
		// 							Name: to.Ptr("testinitcontainer"),
		// 							Args: []*string{
		// 								to.Ptr("-c"),
		// 								to.Ptr("while true; do echo hello; sleep 10;done")},
		// 								Command: []*string{
		// 									to.Ptr("/bin/sh")},
		// 									Image: to.Ptr("repo/testcontainer:v4"),
		// 									Resources: &armappcontainers.SessionContainerResources{
		// 										CPU: to.Ptr[float64](0.25),
		// 										Memory: to.Ptr("0.5Gi"),
		// 									},
		// 							}},
		// 							Ingress: &armappcontainers.SessionIngress{
		// 								TargetPort: to.Ptr[int32](80),
		// 							},
		// 						},
		// 						DynamicPoolConfiguration: &armappcontainers.DynamicPoolConfiguration{
		// 							LifecycleConfiguration: &armappcontainers.LifecycleConfiguration{
		// 								CooldownPeriodInSeconds: to.Ptr[int32](600),
		// 								LifecycleType: to.Ptr(armappcontainers.LifecycleTypeTimed),
		// 							},
		// 						},
		// 						EnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube"),
		// 						NodeCount: to.Ptr[int32](1),
		// 						PoolManagementEndpoint: to.Ptr("https://testsessionpool.agreeableriver-3d30edf1.eastus.azurecontainerapps.io"),
		// 						PoolManagementType: to.Ptr(armappcontainers.PoolManagementTypeDynamic),
		// 						ProvisioningState: to.Ptr(armappcontainers.SessionPoolProvisioningStateSucceeded),
		// 						ScaleConfiguration: &armappcontainers.ScaleConfiguration{
		// 							MaxConcurrentSessions: to.Ptr[int32](500),
		// 							ReadySessionInstances: to.Ptr[int32](100),
		// 						},
		// 						SessionNetworkConfiguration: &armappcontainers.SessionNetworkConfiguration{
		// 							Status: to.Ptr(armappcontainers.SessionNetworkStatusEgressEnabled),
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
