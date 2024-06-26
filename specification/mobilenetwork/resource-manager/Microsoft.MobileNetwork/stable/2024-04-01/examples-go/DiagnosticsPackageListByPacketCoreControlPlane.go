package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/DiagnosticsPackageListByPacketCoreControlPlane.json
func ExampleDiagnosticsPackagesClient_NewListByPacketCoreControlPlanePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiagnosticsPackagesClient().NewListByPacketCoreControlPlanePager("rg1", "TestPacketCoreCP", nil)
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
		// page.DiagnosticsPackageListResult = armmobilenetwork.DiagnosticsPackageListResult{
		// 	Value: []*armmobilenetwork.DiagnosticsPackage{
		// 		{
		// 			Name: to.Ptr("dp1"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/diagnosticsPackages/pc1"),
		// 			Properties: &armmobilenetwork.DiagnosticsPackagePropertiesFormat{
		// 				ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
		// 				Status: to.Ptr(armmobilenetwork.DiagnosticsPackageStatus("Stopped")),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dp2"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/diagnosticsPackages/dp2"),
		// 			Properties: &armmobilenetwork.DiagnosticsPackagePropertiesFormat{
		// 				ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
		// 				Status: to.Ptr(armmobilenetwork.DiagnosticsPackageStatus("Stopped")),
		// 			},
		// 	}},
		// }
	}
}
