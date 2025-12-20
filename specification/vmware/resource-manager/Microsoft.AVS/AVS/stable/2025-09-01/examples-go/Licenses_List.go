package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/Licenses_List.json
func ExampleLicensesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLicensesClient().NewListPager("group1", "cloud1", nil)
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
		// page = armavs.LicensesClientListResponse{
		// 	LicenseListResult: armavs.LicenseListResult{
		// 		Value: []*armavs.License{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/licenses/vcf-license-001"),
		// 				Name: to.Ptr("VmwareFirewall"),
		// 				Type: to.Ptr("Microsoft.AVS/privateClouds/licenses"),
		// 				Properties: &armavs.VmwareFirewallLicenseProperties{
		// 					Kind: to.Ptr(armavs.LicenseKindVmwareFirewall),
		// 					ProvisioningState: to.Ptr(armavs.LicenseProvisioningStateSucceeded),
		// 					EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-01T00:00:00Z"); return t}()),
		// 					Cores: to.Ptr[int32](64),
		// 					BroadcomSiteID: to.Ptr("123456"),
		// 					BroadcomContractNumber: to.Ptr("123456"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
