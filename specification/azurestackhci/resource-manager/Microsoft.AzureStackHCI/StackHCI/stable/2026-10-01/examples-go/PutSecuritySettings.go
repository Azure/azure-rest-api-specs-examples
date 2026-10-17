package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/PutSecuritySettings.json
func ExampleSecuritySettingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
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
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.SecuritySettingsClientCreateOrUpdateResponse{
	// 	SecuritySetting: armazurestackhci.SecuritySetting{
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/clusters/securitySettings"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster/securitySettings/default"),
	// 		Properties: &armazurestackhci.SecurityProperties{
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			SecuredCoreComplianceAssignment: to.Ptr(armazurestackhci.ComplianceAssignmentTypeAudit),
	// 			SecurityComplianceStatus: &armazurestackhci.SecurityComplianceStatus{
	// 				DataAtRestEncrypted: to.Ptr(armazurestackhci.ComplianceStatusCompliant),
	// 				DataInTransitProtected: to.Ptr(armazurestackhci.ComplianceStatusCompliant),
	// 				LastUpdated: to.Ptr(time.Date(2023, time.November, 14, 7, 9, 44, 771000000, time.UTC)),
	// 				SecuredCoreCompliance: to.Ptr(armazurestackhci.ComplianceStatusCompliant),
	// 				WdacCompliance: to.Ptr(armazurestackhci.ComplianceStatusCompliant),
	// 			},
	// 			SmbEncryptionForIntraClusterTrafficComplianceAssignment: to.Ptr(armazurestackhci.ComplianceAssignmentTypeAudit),
	// 			WdacComplianceAssignment: to.Ptr(armazurestackhci.ComplianceAssignmentTypeApplyAndAutoCorrect),
	// 		},
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2021, time.January, 1, 17, 18, 19, 123456700, time.UTC)),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2021, time.January, 2, 17, 18, 19, 123456700, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
