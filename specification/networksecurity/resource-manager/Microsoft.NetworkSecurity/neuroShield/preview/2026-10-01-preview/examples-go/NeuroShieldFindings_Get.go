package armnetworksecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networksecurity/armnetworksecurity"
)

// Generated from example definition: 2026-10-01-preview/NeuroShieldFindings_Get.json
func ExampleNeuroShieldFindingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworksecurity.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNeuroShieldFindingsClient().Get(ctx, "myResourceGroup", "myNeuroShield", "vm-01", "CVE-2026-12345-vm-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetworksecurity.NeuroShieldFindingsClientGetResponse{
	// 	NeuroShieldFinding: armnetworksecurity.NeuroShieldFinding{
	// 		Name: to.Ptr("CVE-2026-12345-vm-01"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/microsoft.networksecurity/neuroShields/myNeuroShield/protectedResources/vm-01/findings/CVE-2026-12345-vm-01"),
	// 		Type: to.Ptr("microsoft.networksecurity/neuroShields/protectedResources/findings"),
	// 		Properties: &armnetworksecurity.NeuroShieldFindingProperties{
	// 			ProvisioningState: to.Ptr(armnetworksecurity.ResourceProvisioningStateSucceeded),
	// 			CVEID: to.Ptr("CVE-2026-12345"),
	// 			VirtualMachineARMID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myComputeRg/providers/Microsoft.Compute/virtualMachines/vm-01"),
	// 			Severity: to.Ptr(armnetworksecurity.NeuroShieldSeverityThresholdHigh),
	// 			Description: to.Ptr("Kernel package vulnerability detected on vm-01."),
	// 			Reason: to.Ptr("The installed kernel package is within the affected version range."),
	// 			DetectedTime: to.Ptr(time.Date(2026, time.May, 28, 10, 15, 0, 0, time.UTC)),
	// 			Status: to.Ptr(armnetworksecurity.NeuroShieldFindingStatusActive),
	// 		},
	// 	},
	// }
}
