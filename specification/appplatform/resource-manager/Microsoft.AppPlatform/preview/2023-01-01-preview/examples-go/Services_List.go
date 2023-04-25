package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-01-01-preview/examples/Services_List.json
func ExampleServicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListPager("myResourceGroup", nil)
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
		// page.ServiceResourceList = armappplatform.ServiceResourceList{
		// 	Value: []*armappplatform.ServiceResource{
		// 		{
		// 			Name: to.Ptr("myservice"),
		// 			Type: to.Ptr("Microsoft.AppPlatform/Spring"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice"),
		// 			SystemData: &armappplatform.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 				CreatedBy: to.Ptr("sample-user"),
		// 				CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("sample-user"),
		// 				LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Properties: &armappplatform.ClusterResourceProperties{
		// 				NetworkProfile: &armappplatform.NetworkProfile{
		// 					IngressConfig: &armappplatform.IngressConfig{
		// 						ReadTimeoutInSeconds: to.Ptr[int32](300),
		// 					},
		// 					OutboundIPs: &armappplatform.NetworkProfileOutboundIPs{
		// 						PublicIPs: []*string{
		// 							to.Ptr("20.39.3.173"),
		// 							to.Ptr("40.64.67.13")},
		// 						},
		// 						RequiredTraffics: []*armappplatform.RequiredTraffic{
		// 							{
		// 								Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
		// 								IPs: []*string{
		// 									to.Ptr("20.62.211.25"),
		// 									to.Ptr("52.188.47.226")},
		// 									Port: to.Ptr[int32](443),
		// 									Protocol: to.Ptr("TCP"),
		// 								},
		// 								{
		// 									Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
		// 									IPs: []*string{
		// 										to.Ptr("20.62.211.25"),
		// 										to.Ptr("52.188.47.226")},
		// 										Port: to.Ptr[int32](1194),
		// 										Protocol: to.Ptr("UDP"),
		// 									},
		// 									{
		// 										Direction: to.Ptr(armappplatform.TrafficDirectionOutbound),
		// 										IPs: []*string{
		// 											to.Ptr("20.62.211.25"),
		// 											to.Ptr("52.188.47.226")},
		// 											Port: to.Ptr[int32](9000),
		// 											Protocol: to.Ptr("TCP"),
		// 									}},
		// 								},
		// 								ProvisioningState: to.Ptr(armappplatform.ProvisioningStateSucceeded),
		// 								ServiceID: to.Ptr("12345678abcd1234abcd12345678abcd"),
		// 							},
		// 							SKU: &armappplatform.SKU{
		// 								Name: to.Ptr("S0"),
		// 								Tier: to.Ptr("Standard"),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("myservice1"),
		// 							Type: to.Ptr("Microsoft.AppPlatform/Spring"),
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice1"),
		// 							SystemData: &armappplatform.SystemData{
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
		// 								CreatedBy: to.Ptr("sample-user"),
		// 								CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("sample-user"),
		// 								LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
		// 							},
		// 							Location: to.Ptr("eastus"),
		// 							Tags: map[string]*string{
		// 								"key1": to.Ptr("value1"),
		// 							},
		// 							Properties: &armappplatform.ClusterResourceProperties{
		// 								InfraResourceGroup: to.Ptr("myenvironment_SpringApps_12345678abcd1234abcd12345678abc1"),
		// 								ManagedEnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.App/managedEnvironments/myenvironment"),
		// 								ProvisioningState: to.Ptr(armappplatform.ProvisioningStateSucceeded),
		// 								ServiceID: to.Ptr("12345678abcd1234abcd12345678abc1"),
		// 							},
		// 							SKU: &armappplatform.SKU{
		// 								Name: to.Ptr("S0"),
		// 								Tier: to.Ptr("StandardGen2"),
		// 							},
		// 					}},
		// 				}
	}
}
