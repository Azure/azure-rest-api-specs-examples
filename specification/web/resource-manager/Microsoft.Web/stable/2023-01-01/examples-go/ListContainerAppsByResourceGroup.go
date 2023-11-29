package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/ListContainerAppsByResourceGroup.json
func ExampleContainerAppsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsClient().NewListByResourceGroupPager("rg", nil)
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
		// page.ContainerAppCollection = armappservice.ContainerAppCollection{
		// 	Value: []*armappservice.ContainerApp{
		// 		{
		// 			Name: to.Ptr("testcontainerApp0"),
		// 			Type: to.Ptr("Microsoft.Web/containerApps"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/containerApps/testcontainerApp0"),
		// 			Kind: to.Ptr("containerApp"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappservice.ContainerAppProperties{
		// 				Configuration: &armappservice.Configuration{
		// 					Ingress: &armappservice.Ingress{
		// 						External: to.Ptr(true),
		// 						Fqdn: to.Ptr("testcontainerApp0.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
		// 						TargetPort: to.Ptr[int32](3000),
		// 						Transport: to.Ptr(armappservice.IngressTransportMethodAuto),
		// 					},
		// 				},
		// 				KubeEnvironmentID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/kubeEnvironments/demokube"),
		// 				LatestRevisionFqdn: to.Ptr("testcontainerApp0-pjxhsye.demokube-t24clv0g.eastus.containerApps.k4apps.io"),
		// 				ProvisioningState: to.Ptr(armappservice.ContainerAppProvisioningStateSucceeded),
		// 				Template: &armappservice.Template{
		// 					Containers: []*armappservice.Container{
		// 						{
		// 							Name: to.Ptr("testcontainerApp0"),
		// 							Image: to.Ptr("repo/testcontainerApp0:v4"),
		// 							Resources: &armappservice.ContainerResources{
		// 								CPU: to.Ptr[float64](0.2),
		// 								Memory: to.Ptr("100Mi"),
		// 							},
		// 					}},
		// 					Dapr: &armappservice.Dapr{
		// 						AppPort: to.Ptr[int32](3000),
		// 						Enabled: to.Ptr(true),
		// 					},
		// 					Scale: &armappservice.Scale{
		// 						MaxReplicas: to.Ptr[int32](5),
		// 						MinReplicas: to.Ptr[int32](1),
		// 						Rules: []*armappservice.ScaleRule{
		// 							{
		// 								Name: to.Ptr("httpscalingrule"),
		// 								HTTP: &armappservice.HTTPScaleRule{
		// 									Metadata: map[string]*string{
		// 										"concurrentRequests": to.Ptr("50"),
		// 									},
		// 								},
		// 						}},
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
