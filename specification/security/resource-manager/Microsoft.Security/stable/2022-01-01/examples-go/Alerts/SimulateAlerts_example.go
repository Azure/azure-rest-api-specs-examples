package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/SimulateAlerts_example.json
func ExampleAlertsClient_BeginSimulate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAlertsClient().BeginSimulate(ctx, "centralus", armsecurity.AlertSimulatorRequestBody{
		Properties: &armsecurity.AlertSimulatorBundlesRequestProperties{
			Kind: to.Ptr(armsecurity.KindBundles),
			Bundles: []*armsecurity.BundleType{
				to.Ptr(armsecurity.BundleTypeAppServices),
				to.Ptr(armsecurity.BundleTypeDNS),
				to.Ptr(armsecurity.BundleTypeKeyVaults),
				to.Ptr(armsecurity.BundleTypeKubernetesService),
				to.Ptr(armsecurity.BundleTypeResourceManager),
				to.Ptr(armsecurity.BundleTypeSQLServers),
				to.Ptr(armsecurity.BundleTypeStorageAccounts),
				to.Ptr(armsecurity.BundleTypeVirtualMachines),
				to.Ptr(armsecurity.BundleTypeCosmosDbs)},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
