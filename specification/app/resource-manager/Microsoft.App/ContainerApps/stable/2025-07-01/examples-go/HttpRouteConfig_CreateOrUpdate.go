package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b2965096067d6f8374b5485b0568fd36e7c9d099/specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/HttpRouteConfig_CreateOrUpdate.json
func ExampleHTTPRouteConfigClient_CreateOrUpdate_createOrUpdateHttpRoute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHTTPRouteConfigClient().CreateOrUpdate(ctx, "examplerg", "testcontainerenv", "httproutefriendlyname", &armappcontainers.HTTPRouteConfigClientCreateOrUpdateOptions{HTTPRouteConfigEnvelope: &armappcontainers.HTTPRouteConfig{
		Properties: &armappcontainers.HTTPRouteConfigProperties{
			CustomDomains: []*armappcontainers.CustomDomain{
				{
					Name:          to.Ptr("example.com"),
					BindingType:   to.Ptr(armappcontainers.BindingTypeSniEnabled),
					CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/certificates/certificate-1"),
				}},
			Rules: []*armappcontainers.HTTPRouteRule{
				{
					Description: to.Ptr("random-description"),
					Routes: []*armappcontainers.HTTPRoute{
						{
							Action: &armappcontainers.HTTPRouteAction{
								PrefixRewrite: to.Ptr("/v1/api"),
							},
							Match: &armappcontainers.HTTPRouteMatch{
								Path:          to.Ptr("/v1"),
								CaseSensitive: to.Ptr(true),
							},
						}},
					Targets: []*armappcontainers.HTTPRouteTarget{
						{
							ContainerApp: to.Ptr("capp-1"),
							Revision:     to.Ptr("capp-1--0000001"),
						}},
				}},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HTTPRouteConfig = armappcontainers.HTTPRouteConfig{
	// 	Type: to.Ptr("Microsoft.App/managedEnvironments/httpRouteConfigs"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/httpRouteConfigs/route-1"),
	// 	Properties: &armappcontainers.HTTPRouteConfigProperties{
	// 		CustomDomains: []*armappcontainers.CustomDomain{
	// 			{
	// 				Name: to.Ptr("example.com"),
	// 				BindingType: to.Ptr(armappcontainers.BindingTypeSniEnabled),
	// 				CertificateID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/testcontainerenv/certificates/certificate-1"),
	// 		}},
	// 		Fqdn: to.Ptr("app1.example.com"),
	// 		ProvisioningState: to.Ptr(armappcontainers.HTTPRouteProvisioningStateSucceeded),
	// 		Rules: []*armappcontainers.HTTPRouteRule{
	// 			{
	// 				Description: to.Ptr("random-description"),
	// 				Routes: []*armappcontainers.HTTPRoute{
	// 					{
	// 						Action: &armappcontainers.HTTPRouteAction{
	// 							PrefixRewrite: to.Ptr("/v1/api"),
	// 						},
	// 						Match: &armappcontainers.HTTPRouteMatch{
	// 							Path: to.Ptr("/v1"),
	// 							CaseSensitive: to.Ptr(true),
	// 						},
	// 				}},
	// 				Targets: []*armappcontainers.HTTPRouteTarget{
	// 					{
	// 						ContainerApp: to.Ptr("capp-1"),
	// 						Revision: to.Ptr("capp-1--0000001"),
	// 				}},
	// 		}},
	// 	},
	// }
}
