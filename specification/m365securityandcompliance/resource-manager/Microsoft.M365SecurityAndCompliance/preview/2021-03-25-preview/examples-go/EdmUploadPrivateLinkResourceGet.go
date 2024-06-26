package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/EdmUploadPrivateLinkResourceGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.PrivateLinkResource = armm365securityandcompliance.PrivateLinkResource{
	// 	Name: to.Ptr("fhir"),
	// 	Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForEDMUpload/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForEDMUpload/service1/privateLinkResources/fhir"),
	// 	Properties: &armm365securityandcompliance.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("fhir"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("fhir")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.security.microsoft.com")},
	// 			},
	// 			SystemData: &armm365securityandcompliance.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 				CreatedBy: to.Ptr("sove"),
	// 				CreatedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("sove"),
	// 				LastModifiedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 			},
	// 		}
}
