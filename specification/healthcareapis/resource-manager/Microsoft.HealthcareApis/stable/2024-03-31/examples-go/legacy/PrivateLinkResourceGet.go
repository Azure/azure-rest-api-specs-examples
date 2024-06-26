package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/PrivateLinkResourceGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "rgname", "service1", "fhir", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourceDescription = armhealthcareapis.PrivateLinkResourceDescription{
	// 	Name: to.Ptr("fhir"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/services/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/services/service1/privateLinkResources/fhir"),
	// 	Properties: &armhealthcareapis.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("fhir"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("fhir")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.azurehealthcareapis.com")},
	// 			},
	// 		}
}
