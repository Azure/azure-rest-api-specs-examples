package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v3"
)

// Generated from example definition: 2025-04-01-preview/legacy/ServiceCreateMinimum.json
func ExampleServicesClient_BeginCreateOrUpdate_createOrUpdateAServiceWithMinimumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "rg1", "service2", armhealthcareapis.ServicesDescription{
		Kind:     to.Ptr(armhealthcareapis.KindFhirR4),
		Location: to.Ptr("westus2"),
		Properties: &armhealthcareapis.ServicesProperties{
			AccessPolicies: []*armhealthcareapis.ServiceAccessPolicyEntry{
				{
					ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
				},
			},
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhealthcareapis.ServicesClientCreateOrUpdateResponse{
	// 	ServicesDescription: armhealthcareapis.ServicesDescription{
	// 		Name: to.Ptr("service2"),
	// 		Type: to.Ptr("Microsoft.HealthcareApis/services"),
	// 		Etag: to.Ptr("etagvalue"),
	// 		ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/services/service2"),
	// 		Kind: to.Ptr(armhealthcareapis.KindFhirR4),
	// 		Location: to.Ptr("westus2"),
	// 		Properties: &armhealthcareapis.ServicesProperties{
	// 			AccessPolicies: []*armhealthcareapis.ServiceAccessPolicyEntry{
	// 				{
	// 					ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
	// 				},
	// 			},
	// 			AuthenticationConfiguration: &armhealthcareapis.ServiceAuthenticationConfigurationInfo{
	// 				Audience: to.Ptr("https://azurehealthcareapis.com"),
	// 				Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 				SmartProxyEnabled: to.Ptr(false),
	// 			},
	// 			CorsConfiguration: &armhealthcareapis.ServiceCorsConfigurationInfo{
	// 				AllowCredentials: to.Ptr(false),
	// 				Headers: []*string{
	// 				},
	// 				Methods: []*string{
	// 				},
	// 				Origins: []*string{
	// 				},
	// 			},
	// 			CosmosDbConfiguration: &armhealthcareapis.ServiceCosmosDbConfigurationInfo{
	// 				OfferThroughput: to.Ptr[int32](1000),
	// 			},
	// 			PrivateEndpointConnections: []*armhealthcareapis.PrivateEndpointConnection{
	// 			},
	// 			ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateCreating),
	// 			PublicNetworkAccess: to.Ptr(armhealthcareapis.PublicNetworkAccessDisabled),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
