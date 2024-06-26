package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/privatelink/PrivateLinkResourcesListByWorkspace.json
func ExampleWorkspacePrivateLinkResourcesClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacePrivateLinkResourcesClient().NewListByWorkspacePager("testRG", "workspace1", nil)
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
		// page.PrivateLinkResourceListResultDescription = armhealthcareapis.PrivateLinkResourceListResultDescription{
		// 	Value: []*armhealthcareapis.PrivateLinkResourceDescription{
		// 		{
		// 			Name: to.Ptr("healthcareworkspace"),
		// 			Type: to.Ptr("Microsoft.HealthcareApis/workspaces/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/workspaces/workspace1/privateLinkResources/healthcareworkspace"),
		// 			Properties: &armhealthcareapis.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("healthcareworkspace"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("fhirservices"),
		// 					to.Ptr("dicomservices")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.fhir.azurehealthcareapis.com"),
		// 						to.Ptr("privatelink.dicom.azurehealthcareapis.com")},
		// 					},
		// 			}},
		// 		}
	}
}
