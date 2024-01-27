package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/685aad3f33d355c1d9c89d493ee9398865367bd8/specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/BuildpackBinding_List.json
func ExampleBuildpackBindingClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBuildpackBindingClient().NewListPager("myResourceGroup", "myservice", "default", "default", nil)
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
		// page.BuildpackBindingResourceCollection = armappplatform.BuildpackBindingResourceCollection{
		// 	Value: []*armappplatform.BuildpackBindingResource{
		// 		{
		// 			Name: to.Ptr("myBuildpackBinding"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring/buildServices/builders/buildpackBindings"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/buildServices/default/builders/default/buildpackBindings/myBuildpackBinding"),
		// 			SystemData: &armappplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample-user"),
		// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample-user"),
		// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armappplatform.BuildpackBindingProperties{
		// 				BindingType: to.Ptr(armappplatform.BindingTypeApplicationInsights),
		// 				LaunchProperties: &armappplatform.BuildpackBindingLaunchProperties{
		// 					Properties: map[string]*string{
		// 						"abc": to.Ptr("def"),
		// 						"any-string": to.Ptr("any-string"),
		// 						"sampling-rate": to.Ptr("12.0"),
		// 					},
		// 					Secrets: map[string]*string{
		// 						"connection-string": to.Ptr("*"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
