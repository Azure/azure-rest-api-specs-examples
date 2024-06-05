package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/license/License_Update.json
func ExampleLicensesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLicensesClient().BeginUpdate(ctx, "myResourceGroup", "{licenseName}", armhybridcompute.LicenseUpdate{
		Properties: &armhybridcompute.LicenseUpdateProperties{
			LicenseDetails: &armhybridcompute.LicenseUpdatePropertiesLicenseDetails{
				Type:       to.Ptr(armhybridcompute.LicenseCoreTypePCore),
				Edition:    to.Ptr(armhybridcompute.LicenseEditionDatacenter),
				Processors: to.Ptr[int32](6),
				State:      to.Ptr(armhybridcompute.LicenseStateActivated),
				Target:     to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
			},
			LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.License = armhybridcompute.License{
	// 	Name: to.Ptr("{licenseName}"),
	// 	Type: to.Ptr("Microsoft.HybridCompute/licenses"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/licenses/{licenseName}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armhybridcompute.LicenseProperties{
	// 		LicenseDetails: &armhybridcompute.LicenseDetails{
	// 			Type: to.Ptr(armhybridcompute.LicenseCoreTypePCore),
	// 			AssignedLicenses: to.Ptr[int32](8),
	// 			Edition: to.Ptr(armhybridcompute.LicenseEditionDatacenter),
	// 			ImmutableID: to.Ptr("<generated Guid>"),
	// 			Processors: to.Ptr[int32](6),
	// 			State: to.Ptr(armhybridcompute.LicenseStateActivated),
	// 			Target: to.Ptr(armhybridcompute.LicenseTargetWindowsServer2012),
	// 		},
	// 		LicenseType: to.Ptr(armhybridcompute.LicenseTypeESU),
	// 		TenantID: to.Ptr("{tenandId}"),
	// 	},
	// }
}
