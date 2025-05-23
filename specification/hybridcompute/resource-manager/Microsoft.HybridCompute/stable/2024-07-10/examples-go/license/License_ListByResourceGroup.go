package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f41d0c9332078cb2ef07b749081d94915255ada5/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/license/License_ListByResourceGroup.json
func ExampleLicensesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLicensesClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.LicensesListResult = armhybridcompute.LicensesListResult{
		// 	Value: []*armhybridcompute.License{
		// 		{
		// 			Name: to.Ptr("{licenseName}"),
		// 			Type: to.Ptr("Microsoft.HybridCompute/licenses"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Licenses/{licenseName}"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armhybridcompute.LicenseProperties{
		// 				LicenseDetails: &armhybridcompute.LicenseDetails{
		// 					Type: to.Ptr(armhybridcompute.LicenseCoreTypePCore),
		// 					AssignedLicenses: to.Ptr[int32](8),
		// 					Edition: to.Ptr(armhybridcompute.LicenseEditionDatacenter),
		// 					ImmutableID: to.Ptr("<generated Guid>"),
		// 					Processors: to.Ptr[int32](6),
		// 					State: to.Ptr(armhybridcompute.LicenseStateActivated),
		// 					Target: to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
		// 				},
		// 				LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
		// 				TenantID: to.Ptr("{tenandId}"),
		// 			},
		// 	}},
		// }
	}
}
