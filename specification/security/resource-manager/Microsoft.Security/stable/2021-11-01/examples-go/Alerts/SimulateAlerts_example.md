```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2021-11-01/examples/Alerts/SimulateAlerts_example.json
func ExampleAlertsClient_Simulate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAlertsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Simulate(ctx,
		"centralus",
		armsecurity.AlertSimulatorRequestBody{
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
					to.Ptr(armsecurity.BundleTypeVirtualMachines)},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.7.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.
