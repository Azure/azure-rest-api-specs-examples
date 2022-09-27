package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/SimulateAlerts_example.json
func ExampleAlertsClient_BeginSimulate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAlertsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSimulate(ctx, "centralus", armsecurity.AlertSimulatorRequestBody{
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
