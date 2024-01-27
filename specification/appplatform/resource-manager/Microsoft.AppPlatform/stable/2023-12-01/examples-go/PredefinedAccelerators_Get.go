package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/685aad3f33d355c1d9c89d493ee9398865367bd8/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/PredefinedAccelerators_Get.json
func ExamplePredefinedAcceleratorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPredefinedAcceleratorsClient().Get(ctx, "myResourceGroup", "myservice", "default", "acc-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PredefinedAcceleratorResource = armappplatform.PredefinedAcceleratorResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/applicationAccelerators/predefinedAccelerators"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/applicationAccelerators/default/predefinedAccelerators/acc-name"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.PredefinedAcceleratorProperties{
	// 		Description: to.Ptr("acc-desc"),
	// 		AcceleratorTags: []*string{
	// 			to.Ptr("tag-a"),
	// 			to.Ptr("tag-b")},
	// 			DisplayName: to.Ptr("acc-name"),
	// 			IconURL: to.Ptr("acc-icon"),
	// 			ProvisioningState: to.Ptr(armappplatform.PredefinedAcceleratorProvisioningStateSucceeded),
	// 			State: to.Ptr(armappplatform.PredefinedAcceleratorStateEnabled),
	// 		},
	// 		SKU: &armappplatform.SKU{
	// 			Name: to.Ptr("E0"),
	// 			Capacity: to.Ptr[int32](2),
	// 			Tier: to.Ptr("Enterprise"),
	// 		},
	// 	}
}
