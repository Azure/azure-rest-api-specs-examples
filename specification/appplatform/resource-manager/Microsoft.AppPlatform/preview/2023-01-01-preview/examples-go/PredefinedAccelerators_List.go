package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-01-01-preview/examples/PredefinedAccelerators_List.json
func ExamplePredefinedAcceleratorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPredefinedAcceleratorsClient().NewListPager("myResourceGroup", "myservice", "default", nil)
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
		// page.PredefinedAcceleratorResourceCollection = armappplatform.PredefinedAcceleratorResourceCollection{
		// 	Value: []*armappplatform.PredefinedAcceleratorResource{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/applicationAccelerators/predefinedAccelerators"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/applicationAccelerators/default/predefinedAccelerators"),
		// 			SystemData: &armappplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample-user"),
		// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample-user"),
		// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armappplatform.PredefinedAcceleratorProperties{
		// 				Description: to.Ptr("acc-desc"),
		// 				AcceleratorTags: []*string{
		// 					to.Ptr("tag-a"),
		// 					to.Ptr("tag-b")},
		// 					DisplayName: to.Ptr("acc-name"),
		// 					IconURL: to.Ptr("acc-icon"),
		// 					ProvisioningState: to.Ptr(armappplatform.PredefinedAcceleratorProvisioningStateSucceeded),
		// 					State: to.Ptr(armappplatform.PredefinedAcceleratorStateEnabled),
		// 				},
		// 				SKU: &armappplatform.SKU{
		// 					Name: to.Ptr("E0"),
		// 					Capacity: to.Ptr[int32](2),
		// 					Tier: to.Ptr("Enterprise"),
		// 				},
		// 		}},
		// 	}
	}
}
