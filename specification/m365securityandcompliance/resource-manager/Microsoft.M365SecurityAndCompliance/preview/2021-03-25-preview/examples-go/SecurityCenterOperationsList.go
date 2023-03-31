package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterOperationsList.json
func ExampleOperationsClient_NewListPager_listSecurityCenterOperations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armm365securityandcompliance.OperationListResult{
		// 	Value: []*armm365securityandcompliance.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/read"),
		// 			Display: &armm365securityandcompliance.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/write"),
		// 			Display: &armm365securityandcompliance.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/delete"),
		// 			Display: &armm365securityandcompliance.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/privateEndpointConnections/read"),
		// 			Display: &armm365securityandcompliance.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/privateEndpointConnections/write"),
		// 			Display: &armm365securityandcompliance.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/privateEndpointConnections/delete"),
		// 			Display: &armm365securityandcompliance.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/privateLinkResources/read"),
		// 			Display: &armm365securityandcompliance.OperationDisplay{
		// 			},
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
