package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration/v2"
)

// Generated from example definition: 2024-04-05/getVMSSGuestConfigurationAssignment.json
func ExampleAssignmentsVMSSClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armguestconfiguration.NewClientFactory("mySubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsVMSSClient().Get(ctx, "myResourceGroupName", "myVMSSName", "SecureProtocol", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armguestconfiguration.AssignmentsVMSSClientGetResponse{
	// 	Assignment: armguestconfiguration.Assignment{
	// 		Name: to.Ptr("AuditSecureProtocol"),
	// 		ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.Compute/virtualMachineScaleSets/myVMSSName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/AuditSecureProtocol"),
	// 		Location: to.Ptr("centraluseuap"),
	// 		Properties: &armguestconfiguration.AssignmentProperties{
	// 			AssignmentHash: to.Ptr("E0D8941DD713F284284561648C00C18FA76C8602943C7CD38AFD73B56AE4C35F.E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855"),
	// 			ComplianceStatus: to.Ptr(armguestconfiguration.ComplianceStatusCompliant),
	// 			GuestConfiguration: &armguestconfiguration.Navigation{
	// 				Name: to.Ptr("AuditSecureProtocol"),
	// 				ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
	// 				},
	// 				ContentHash: to.Ptr("content hash"),
	// 				ContentManagedIdentity: to.Ptr("test_identity"),
	// 				ContentURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/builtinconfig/AuditSecureProtocol/AuditSecureProtocol_1.0.0.3.zip"),
	// 				Version: to.Ptr("1.0.0.3"),
	// 			},
	// 			LastComplianceStatusChecked: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T22:14:13Z"); return t}()),
	// 			LatestReportID: nil,
	// 			ProvisioningState: to.Ptr(armguestconfiguration.ProvisioningStateSucceeded),
	// 			ResourceType: to.Ptr("VMSS"),
	// 		},
	// 	},
	// }
}
