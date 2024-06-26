package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/fhirservices/FhirServices_List.json
func ExampleFhirServicesClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFhirServicesClient().NewListByWorkspacePager("testRG", "workspace1", nil)
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
		// page.FhirServiceCollection = armhealthcareapis.FhirServiceCollection{
		// 	Value: []*armhealthcareapis.FhirService{
		// 		{
		// 			Name: to.Ptr("fhirservice1"),
		// 			Type: to.Ptr("Microsoft.HealthcareApis/workspaces/fhirservices"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/fhirservices/fhirservice1"),
		// 			Properties: &armhealthcareapis.FhirServiceProperties{
		// 				ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("fhirservice2"),
		// 			Type: to.Ptr("Microsoft.HealthcareApis/workspaces/fhirservices"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/fhirservices/fhirservice2"),
		// 			Properties: &armhealthcareapis.FhirServiceProperties{
		// 				ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
