package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerSecurityAlertsCreateMax.json
func ExampleServerSecurityAlertPoliciesClient_BeginCreateOrUpdate_updateAServersThreatDetectionPolicyWithAllParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerSecurityAlertPoliciesClient().BeginCreateOrUpdate(ctx, "securityalert-4799", "securityalert-6440", armmysql.SecurityAlertPolicyNameDefault, armmysql.ServerSecurityAlertPolicy{
		Properties: &armmysql.SecurityAlertPolicyProperties{
			DisabledAlerts: []*string{
				to.Ptr("Access_Anomaly"),
				to.Ptr("Usage_Anomaly")},
			EmailAccountAdmins: to.Ptr(true),
			EmailAddresses: []*string{
				to.Ptr("testSecurityAlert@microsoft.com")},
			RetentionDays:           to.Ptr[int32](5),
			State:                   to.Ptr(armmysql.ServerSecurityAlertPolicyStateEnabled),
			StorageAccountAccessKey: to.Ptr("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
			StorageEndpoint:         to.Ptr("https://mystorage.blob.core.windows.net"),
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
	// res.ServerSecurityAlertPolicy = armmysql.ServerSecurityAlertPolicy{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers/securityAlertPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/securityalert-4799/providers/Microsoft.DBforMySQL/servers/securityalert-6440/securityAlertPolicies/default"),
	// 	Properties: &armmysql.SecurityAlertPolicyProperties{
	// 		DisabledAlerts: []*string{
	// 			to.Ptr("Access_Anomaly"),
	// 			to.Ptr("Usage_Anomaly")},
	// 			EmailAccountAdmins: to.Ptr(true),
	// 			EmailAddresses: []*string{
	// 				to.Ptr("testSecurityAlert@microsoft.com")},
	// 				RetentionDays: to.Ptr[int32](5),
	// 				State: to.Ptr(armmysql.ServerSecurityAlertPolicyStateEnabled),
	// 				StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 			},
	// 		}
}
