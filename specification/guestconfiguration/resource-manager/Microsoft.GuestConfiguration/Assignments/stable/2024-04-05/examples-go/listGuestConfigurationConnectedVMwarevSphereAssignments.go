package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration/v2"
)

// Generated from example definition: 2024-04-05/listGuestConfigurationConnectedVMwarevSphereAssignments.json
func ExampleConnectedVMwarevSphereAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armguestconfiguration.NewClientFactory("mySubscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectedVMwarevSphereAssignmentsClient().NewListPager("myResourceGroupName", "myVMName", nil)
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
		// page = armguestconfiguration.ConnectedVMwarevSphereAssignmentsClientListResponse{
		// 	AssignmentList: armguestconfiguration.AssignmentList{
		// 		Value: []*armguestconfiguration.Assignment{
		// 			{
		// 				Name: to.Ptr("AuditSecureProtocol2"),
		// 				ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.ConnectedVMwarevSphere/virtualmachines/myvm/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/AuditSecureProtocol2"),
		// 				Location: to.Ptr("centraluseuap"),
		// 				Properties: &armguestconfiguration.AssignmentProperties{
		// 					AssignmentHash: to.Ptr("content hash"),
		// 					ComplianceStatus: to.Ptr(armguestconfiguration.ComplianceStatusCompliant),
		// 					GuestConfiguration: &armguestconfiguration.Navigation{
		// 						Name: to.Ptr("AuditSecureProtocol2"),
		// 						ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
		// 						},
		// 						ContentHash: to.Ptr("content hash"),
		// 						ContentURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/builtinconfig/AuditSecureProtocol2/AuditSecureProtocol2_1.0.0.3.zip"),
		// 						Version: to.Ptr("1.0.0.3"),
		// 					},
		// 					LastComplianceStatusChecked: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T22:14:13Z"); return t}()),
		// 					LatestReportID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.ConnectedVMwarevSphere/virtualmachines/myvm/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/AuditSecureProtocol2/reports/7367cbb8-ae99-47d0-a33b-a283564d2cb1"),
		// 					ProvisioningState: nil,
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("myAssignment"),
		// 				ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.ConnectedVMwarevSphere/virtualmachines/myvm/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/myAssignment"),
		// 				Location: to.Ptr("centraluseuap"),
		// 				Properties: &armguestconfiguration.AssignmentProperties{
		// 					AssignmentHash: to.Ptr("content hash"),
		// 					ComplianceStatus: to.Ptr(armguestconfiguration.ComplianceStatusCompliant),
		// 					GuestConfiguration: &armguestconfiguration.Navigation{
		// 						Name: to.Ptr("myAssignment"),
		// 						ConfigurationParameter: []*armguestconfiguration.ConfigurationParameter{
		// 						},
		// 						ContentHash: to.Ptr("content hash"),
		// 						ContentURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/builtinconfig/myAssignment/myAssignment.0.0.3.zip"),
		// 						Version: to.Ptr("1.0.0.3"),
		// 					},
		// 					LastComplianceStatusChecked: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T22:14:13Z"); return t}()),
		// 					LatestReportID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.ConnectedVMwarevSphere/virtualmachines/myvm/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/myAssignment/reports/7367cbb8-ae99-47d0-a33b-a283564d2cb1"),
		// 					ProvisioningState: nil,
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
