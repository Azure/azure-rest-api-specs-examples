package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/listAllVMSSGuestConfigurationAssignmentReports.json
func ExampleAssignmentReportsVMSSClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armguestconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssignmentReportsVMSSClient().NewListPager("myResourceGroupName", "myVMSSName", "AuditSecureProtocol", nil)
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
		// page.AssignmentReportList = armguestconfiguration.AssignmentReportList{
		// 	Value: []*armguestconfiguration.AssignmentReport{
		// 		{
		// 			Name: to.Ptr("7367cbb8-ae99-47d0-a33b-a283564d2cb1"),
		// 			ID: to.Ptr("/subscriptions/mysubscriptionid/resourceGroups/myResourceGroupName/providers/Microsoft.Compute/virtualMachineScaleSets/myVMSSName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/AuditSecureProtocol/reports/7367cbb8-ae99-47d0-a33b-a283564d2cb1"),
		// 			Properties: &armguestconfiguration.AssignmentReportProperties{
		// 				Assignment: &armguestconfiguration.AssignmentInfo{
		// 					Name: to.Ptr("AuditSecureProtocol"),
		// 					Configuration: &armguestconfiguration.ConfigurationInfo{
		// 						Name: to.Ptr("AuditSecureProtocol"),
		// 					},
		// 				},
		// 				ComplianceStatus: to.Ptr(armguestconfiguration.ComplianceStatusCompliant),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T22:14:13.000Z"); return t}()),
		// 				ReportID: to.Ptr("7367cbb8-ae99-47d0-a33b-a283564d2cb1"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T22:13:53.000Z"); return t}()),
		// 				VM: &armguestconfiguration.VMInfo{
		// 					ID: to.Ptr("/subscriptions/mysubscriptionid/resourceGroups/myResourceGroupName/providers/Microsoft.Compute/virtualMachineScaleSets/myVMSSName/virtualMachines/1"),
		// 					UUID: to.Ptr("vmuuid"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("41ee2caf-48f9-4999-a793-82ec7c6beb2c"),
		// 			ID: to.Ptr("/subscriptions/mysubscriptionid/resourceGroups/myResourceGroupName/providers/Microsoft.Compute/virtualMachineScaleSets/myVMSSName/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/AuditSecureProtocol/reports/41ee2caf-48f9-4999-a793-82ec7c6beb2c"),
		// 			Properties: &armguestconfiguration.AssignmentReportProperties{
		// 				Assignment: &armguestconfiguration.AssignmentInfo{
		// 					Name: to.Ptr("AuditSecureProtocol"),
		// 					Configuration: &armguestconfiguration.ConfigurationInfo{
		// 						Name: to.Ptr("AuditSecureProtocol"),
		// 					},
		// 				},
		// 				ComplianceStatus: to.Ptr(armguestconfiguration.ComplianceStatusCompliant),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T20:14:13.000Z"); return t}()),
		// 				ReportID: to.Ptr("41ee2caf-48f9-4999-a793-82ec7c6beb2c"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T20:13:53.000Z"); return t}()),
		// 				VM: &armguestconfiguration.VMInfo{
		// 					ID: to.Ptr("/subscriptions/mysubscriptionid/resourceGroups/myResourceGroupName/providers/Microsoft.Compute/virtualMachineScaleSets/myVMSSName/virtualMachines/1"),
		// 					UUID: to.Ptr("vmuuid"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
