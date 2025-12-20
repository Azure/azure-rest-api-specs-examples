package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/Licenses_CreateOrUpdate.json
func ExampleLicensesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLicensesClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", armavs.LicenseNameVmwareFirewall, armavs.License{}, nil)
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
	// res = armavs.LicensesClientCreateOrUpdateResponse{
	// 	License: &armavs.License{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/licenses/VmwareFirewall"),
	// 		Name: to.Ptr("VmwareFirewall"),
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/licenses"),
	// 		Properties: &armavs.VmwareFirewallLicenseProperties{
	// 			Kind: to.Ptr(armavs.LicenseKindVmwareFirewall),
	// 			EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-12-31T23:59:59Z"); return t}()),
	// 			Cores: to.Ptr[int32](16),
	// 			BroadcomSiteID: to.Ptr("123456"),
	// 			BroadcomContractNumber: to.Ptr("123456"),
	// 			ProvisioningState: to.Ptr(armavs.LicenseProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
