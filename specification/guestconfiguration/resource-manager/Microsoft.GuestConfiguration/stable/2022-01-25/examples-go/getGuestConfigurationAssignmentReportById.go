package armguestconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/guestconfiguration/resource-manager/Microsoft.GuestConfiguration/stable/2022-01-25/examples/getGuestConfigurationAssignmentReportById.json
func ExampleAssignmentReportsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armguestconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentReportsClient().Get(ctx, "myResourceGroupName", "AuditSecureProtocol", "7367cbb8-ae99-47d0-a33b-a283564d2cb1", "myvm", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssignmentReport = armguestconfiguration.AssignmentReport{
	// 	Name: to.Ptr("7367cbb8-ae99-47d0-a33b-a283564d2cb1"),
	// 	ID: to.Ptr("/subscriptions/mysubscriptionid/resourceGroups/myResourceGroupName/providers/Microsoft.Compute/virtualMachines/myvm/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/AuditSecureProtocol/reports/7367cbb8-ae99-47d0-a33b-a283564d2cb1"),
	// 	Properties: &armguestconfiguration.AssignmentReportProperties{
	// 		Assignment: &armguestconfiguration.AssignmentInfo{
	// 			Name: to.Ptr("AuditSecureProtocol"),
	// 			Configuration: &armguestconfiguration.ConfigurationInfo{
	// 				Name: to.Ptr("AuditSecureProtocol"),
	// 				Version: to.Ptr("1.0.0.0"),
	// 			},
	// 		},
	// 		ComplianceStatus: to.Ptr(armguestconfiguration.ComplianceStatusCompliant),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T22:14:13.000Z"); return t}()),
	// 		ReportID: to.Ptr("7367cbb8-ae99-47d0-a33b-a283564d2cb1"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T22:13:53.000Z"); return t}()),
	// 		VM: &armguestconfiguration.VMInfo{
	// 			ID: to.Ptr("/subscriptions/mysubscriptionid/resourceGroups/myResourceGroupName/providers/Microsoft.Compute/virtualMachines/myvm"),
	// 			UUID: to.Ptr("vmuuid"),
	// 		},
	// 		Details: &armguestconfiguration.AssignmentReportDetails{
	// 			ComplianceStatus: to.Ptr(armguestconfiguration.ComplianceStatusCompliant),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T22:14:13.000Z"); return t}()),
	// 			JobID: to.Ptr("7367cbb8-ae99-47d0-a33b-a283564d2cb1"),
	// 			OperationType: to.Ptr(armguestconfiguration.TypeConsistency),
	// 			Resources: []*armguestconfiguration.AssignmentReportResource{
	// 				{
	// 					ComplianceStatus: to.Ptr(armguestconfiguration.ComplianceStatusCompliant),
	// 					Properties: map[string]any{
	// 						"ConfigurationName": "IsWebServerSecure",
	// 						"DependsOn": nil,
	// 						"IsSingleInstance": "Yes",
	// 						"ModuleName": "SecureProtocolWebServer",
	// 						"ModuleVersion": "1.0.0.3",
	// 						"Protocols":[]any{
	// 							map[string]any{
	// 								"Ensure": "Absent",
	// 								"Protocol": "SSL 2.0",
	// 							},
	// 							map[string]any{
	// 								"Ensure": "Absent",
	// 								"Protocol": "SSL 3.0",
	// 							},
	// 							map[string]any{
	// 								"Ensure": "Absent",
	// 								"Protocol": "TLS 1.0",
	// 							},
	// 							map[string]any{
	// 								"Ensure": "Absent",
	// 								"Protocol": "PCT 1.0",
	// 							},
	// 							map[string]any{
	// 								"Ensure": "Absent",
	// 								"Protocol": "Multi-Protocol Unified Hello",
	// 							},
	// 							map[string]any{
	// 								"Ensure": "Absent",
	// 								"Protocol": "TLS 1.1",
	// 							},
	// 							map[string]any{
	// 								"Ensure": "Absent",
	// 								"Protocol": "TLS 1.2",
	// 							},
	// 						},
	// 						"PsDscRunAsCredential": nil,
	// 						"Reasons": nil,
	// 						"ResourceId": "[SecureWebServer]s1",
	// 						"SourceInfo": nil,
	// 					},
	// 					Reasons: []*armguestconfiguration.AssignmentReportResourceComplianceReason{
	// 						{
	// 							Code: to.Ptr("DSC::RESOURCE::SUCCESS"),
	// 							Phrase: to.Ptr("Operation successful."),
	// 					}},
	// 			}},
	// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-29T22:13:53.000Z"); return t}()),
	// 		},
	// 	},
	// }
}
