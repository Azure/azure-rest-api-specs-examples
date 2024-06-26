package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/EdmUploadPrivateLinkResourcesListByService.json
func ExamplePrivateLinkResourcesClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().ListByService(ctx, "rgname", "service1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourceListResult = armm365securityandcompliance.PrivateLinkResourceListResult{
	// 	Value: []*armm365securityandcompliance.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("fhir"),
	// 			Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForEDMUpload/privateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForEDMUpload/service1/privateLinkResources/fhir"),
	// 			Properties: &armm365securityandcompliance.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("fhir"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("fhir")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.compliance.microsoft.com")},
	// 					},
	// 			}},
	// 		}
}
