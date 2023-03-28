package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86ead567acadc5a059949bca607a5e702610551f/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/ApplicationAccelerators_List.json
func ExampleApplicationAcceleratorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationAcceleratorsClient().NewListPager("myResourceGroup", "myservice", nil)
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
		// page.ApplicationAcceleratorResourceCollection = armappplatform.ApplicationAcceleratorResourceCollection{
		// 	Value: []*armappplatform.ApplicationAcceleratorResource{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/applicationAccelerators"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/applicationAccelerators/default"),
		// 			SystemData: &armappplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample-user"),
		// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample-user"),
		// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armappplatform.ApplicationAcceleratorProperties{
		// 				Components: []*armappplatform.ApplicationAcceleratorComponent{
		// 					{
		// 						Name: to.Ptr("component1"),
		// 						Instances: []*armappplatform.ApplicationAcceleratorInstance{
		// 							{
		// 								Name: to.Ptr("instance1"),
		// 								Status: to.Ptr("Running"),
		// 						}},
		// 						ResourceRequests: &armappplatform.ApplicationAcceleratorResourceRequests{
		// 							CPU: to.Ptr("1"),
		// 							InstanceCount: to.Ptr[int32](1),
		// 							Memory: to.Ptr("1Gi"),
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armappplatform.ApplicationAcceleratorProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armappplatform.SKU{
		// 				Name: to.Ptr("E0"),
		// 				Capacity: to.Ptr[int32](2),
		// 				Tier: to.Ptr("Enterprise"),
		// 			},
		// 	}},
		// }
	}
}
