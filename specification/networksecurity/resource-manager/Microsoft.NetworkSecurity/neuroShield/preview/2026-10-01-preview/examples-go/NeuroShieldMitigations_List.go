package armnetworksecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networksecurity/armnetworksecurity"
)

// Generated from example definition: 2026-10-01-preview/NeuroShieldMitigations_List.json
func ExampleNeuroShieldMitigationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworksecurity.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNeuroShieldMitigationsClient().NewListPager("myResourceGroup", "myNeuroShield", "vm-01", nil)
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
		// page = armnetworksecurity.NeuroShieldMitigationsClientListResponse{
		// 	NeuroShieldMitigationListResult: armnetworksecurity.NeuroShieldMitigationListResult{
		// 		Value: []*armnetworksecurity.NeuroShieldMitigation{
		// 			{
		// 				Name: to.Ptr("mitigation-01"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/microsoft.networksecurity/neuroShields/myNeuroShield/protectedResources/vm-01/mitigations/mitigation-01"),
		// 				Type: to.Ptr("microsoft.networksecurity/neuroShields/protectedResources/mitigations"),
		// 				Properties: &armnetworksecurity.NeuroShieldMitigationProperties{
		// 					ProvisioningState: to.Ptr(armnetworksecurity.ResourceProvisioningStateSucceeded),
		// 					FindingID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/microsoft.networksecurity/neuroShields/myNeuroShield/protectedResources/vm-01/findings/CVE-2026-12345-vm-01"),
		// 					CVEID: to.Ptr("CVE-2026-12345"),
		// 					VirtualMachineARMID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myComputeRg/providers/Microsoft.Compute/virtualMachines/vm-01"),
		// 					Rank: to.Ptr[int32](1),
		// 					CommandsSummary: to.Ptr("Updates the affected kernel package to the latest approved version."),
		// 					IsRollbackable: to.Ptr(true),
		// 					Status: to.Ptr(armnetworksecurity.NeuroShieldMitigationStatusAwaitingApproval),
		// 					StatusDetails: to.Ptr("Waiting for customer approval."),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
