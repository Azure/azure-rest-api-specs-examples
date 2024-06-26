package armsphere_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sphere/armsphere"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/09c37754dac91874ff689ed1e60effb4268c8669/specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetDeployments.json
func ExampleDeploymentsClient_NewListByDeviceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeploymentsClient().NewListByDeviceGroupPager("MyResourceGroup1", "MyCatalog1", "MyProduct1", "myDeviceGroup1", &armsphere.DeploymentsClientListByDeviceGroupOptions{Filter: nil,
		Top:         nil,
		Skip:        nil,
		Maxpagesize: nil,
	})
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
		// page.DeploymentListResult = armsphere.DeploymentListResult{
		// 	Value: []*armsphere.Deployment{
		// 		{
		// 			Name: to.Ptr("MyDeployment1"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1/products/MyProduct1/deviceGroups/MyDeviceGroup1/deployments/MyDeployment1"),
		// 			Properties: &armsphere.DeploymentProperties{
		// 				DeployedImages: []*armsphere.Image{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1/images/MyImage1"),
		// 						Properties: &armsphere.ImageProperties{
		// 							Description: to.Ptr("description"),
		// 							ComponentID: to.Ptr("componentId"),
		// 							Image: to.Ptr("dGVzdGltYWdl"),
		// 							ImageType: to.Ptr(armsphere.ImageType("ImageType")),
		// 							ProvisioningState: to.Ptr(armsphere.ProvisioningStateSucceeded),
		// 							RegionalDataBoundary: to.Ptr(armsphere.RegionalDataBoundaryNone),
		// 							URI: to.Ptr("imageUri"),
		// 						},
		// 				}},
		// 				DeploymentDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-30T21:51:39.269Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
