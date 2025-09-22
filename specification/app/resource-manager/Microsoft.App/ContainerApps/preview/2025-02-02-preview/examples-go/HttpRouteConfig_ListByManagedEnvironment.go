package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/HttpRouteConfig_ListByManagedEnvironment.json
func ExampleHTTPRouteConfigClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHTTPRouteConfigClient().NewListPager("examplerg", "testcontainerenv", nil)
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
		// page.HTTPRouteConfigCollection = armappcontainers.HTTPRouteConfigCollection{
		// 	Value: []*armappcontainers.HTTPRouteConfig{
		// 		{
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments/httpRouteConfigs"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/httpRouteConfigs/route-1"),
		// 			Properties: &armappcontainers.HTTPRouteConfigProperties{
		// 				CustomDomains: []*armappcontainers.CustomDomain{
		// 					{
		// 						Name: to.Ptr("example.com"),
		// 						BindingType: to.Ptr(armappcontainers.BindingTypeSniEnabled),
		// 						CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/certificates/certificate-1"),
		// 				}},
		// 				Fqdn: to.Ptr("app1.example.com"),
		// 				ProvisioningState: to.Ptr(armappcontainers.HTTPRouteProvisioningStateSucceeded),
		// 				Rules: []*armappcontainers.HTTPRouteRule{
		// 					{
		// 						Routes: []*armappcontainers.HTTPRoute{
		// 							{
		// 								Action: &armappcontainers.HTTPRouteAction{
		// 									PrefixRewrite: to.Ptr("/v1/api"),
		// 								},
		// 								Match: &armappcontainers.HTTPRouteMatch{
		// 									CaseSensitive: to.Ptr(true),
		// 									Prefix: to.Ptr("/api"),
		// 								},
		// 						}},
		// 						Targets: []*armappcontainers.HTTPRouteTarget{
		// 							{
		// 								ContainerApp: to.Ptr("containerApp1"),
		// 						}},
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("Microsoft.App/managedEnvironments/httpRouteConfigs"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/httpRouteConfigs/route-2"),
		// 			Properties: &armappcontainers.HTTPRouteConfigProperties{
		// 				CustomDomains: []*armappcontainers.CustomDomain{
		// 					{
		// 						Name: to.Ptr("example.com"),
		// 						BindingType: to.Ptr(armappcontainers.BindingTypeSniEnabled),
		// 						CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/certificates/certificate-2"),
		// 				}},
		// 				Fqdn: to.Ptr("app2.example.com"),
		// 				ProvisioningState: to.Ptr(armappcontainers.HTTPRouteProvisioningStateSucceeded),
		// 				Rules: []*armappcontainers.HTTPRouteRule{
		// 					{
		// 						Routes: []*armappcontainers.HTTPRoute{
		// 							{
		// 								Action: &armappcontainers.HTTPRouteAction{
		// 									PrefixRewrite: to.Ptr("/v2/api"),
		// 								},
		// 								Match: &armappcontainers.HTTPRouteMatch{
		// 									CaseSensitive: to.Ptr(false),
		// 									Prefix: to.Ptr("/api"),
		// 								},
		// 						}},
		// 						Targets: []*armappcontainers.HTTPRouteTarget{
		// 							{
		// 								ContainerApp: to.Ptr("containerApp2"),
		// 								Revision: to.Ptr("rev-2"),
		// 								Weight: to.Ptr[int32](50),
		// 							},
		// 							{
		// 								ContainerApp: to.Ptr("containerApp2"),
		// 								Revision: to.Ptr("rev-3"),
		// 								Weight: to.Ptr[int32](50),
		// 						}},
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
