package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PutSecuritySettings.json
func ExampleSecuritySettingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSecuritySettingsClient().BeginCreateOrUpdate(ctx, "test-rg", "myCluster", "default", armazurestackhci.SecuritySetting{
		Properties: &armazurestackhci.SecurityProperties{
			SecuredCoreComplianceAssignment:                         to.Ptr(armazurestackhci.ComplianceAssignmentTypeAudit),
			SmbEncryptionForIntraClusterTrafficComplianceAssignment: to.Ptr(armazurestackhci.ComplianceAssignmentTypeAudit),
			WdacComplianceAssignment:                                to.Ptr(armazurestackhci.ComplianceAssignmentTypeApplyAndAutoCorrect),
		},
	}, nil)
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
	// res.SecuritySetting = armazurestackhci.SecuritySetting{
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/clusters/securitySettings"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster/securitySettings/default"),
	// 	SystemData: &armazurestackhci.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 	},
	// 	Properties: &armazurestackhci.SecurityProperties{
	// 		ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 		SecuredCoreComplianceAssignment: to.Ptr(armazurestackhci.ComplianceAssignmentTypeAudit),
	// 		SecurityComplianceStatus: &armazurestackhci.SecurityComplianceStatus{
	// 			DataAtRestEncrypted: to.Ptr(armazurestackhci.ComplianceStatusCompliant),
	// 			DataInTransitProtected: to.Ptr(armazurestackhci.ComplianceStatusCompliant),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-14T07:09:44.771Z"); return t}()),
	// 			SecuredCoreCompliance: to.Ptr(armazurestackhci.ComplianceStatusCompliant),
	// 			WdacCompliance: to.Ptr(armazurestackhci.ComplianceStatusCompliant),
	// 		},
	// 		SmbEncryptionForIntraClusterTrafficComplianceAssignment: to.Ptr(armazurestackhci.ComplianceAssignmentTypeAudit),
	// 		WdacComplianceAssignment: to.Ptr(armazurestackhci.ComplianceAssignmentTypeApplyAndAutoCorrect),
	// 	},
	// }
}
