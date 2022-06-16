package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/ClustersUpdate.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armoperationalinsights.NewClustersClient("00000000-0000-0000-0000-00000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"oiautorest6685",
		"oiautorest6685",
		armoperationalinsights.ClusterPatch{
			Identity: &armoperationalinsights.Identity{
				Type: to.Ptr(armoperationalinsights.IdentityTypeUserAssigned),
				UserAssignedIdentities: map[string]*armoperationalinsights.UserIdentityProperties{
					"/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/oiautorest6685/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity": {},
				},
			},
			Properties: &armoperationalinsights.ClusterPatchProperties{
				KeyVaultProperties: &armoperationalinsights.KeyVaultProperties{
					KeyName:     to.Ptr("aztest2170cert"),
					KeyRsaSize:  to.Ptr[int32](1024),
					KeyVaultURI: to.Ptr("https://aztest2170.vault.azure.net"),
					KeyVersion:  to.Ptr("654ft6c4e63845cbb50fd6fg51540429"),
				},
			},
			SKU: &armoperationalinsights.ClusterSKU{
				Name:     to.Ptr(armoperationalinsights.ClusterSKUNameEnumCapacityReservation),
				Capacity: to.Ptr(armoperationalinsights.CapacityTenHundred),
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("val1"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
