package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerSecurityAlertsCreateMin.json
func ExampleServerSecurityAlertPoliciesClient_BeginCreateOrUpdate_updateAServersThreatDetectionPolicyWithMinimalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerSecurityAlertPoliciesClient().BeginCreateOrUpdate(ctx, "securityalert-4799", "securityalert-6440", armmariadb.SecurityAlertPolicyNameDefault, armmariadb.ServerSecurityAlertPolicy{
		Properties: &armmariadb.SecurityAlertPolicyProperties{
			EmailAccountAdmins: to.Ptr(true),
			State:              to.Ptr(armmariadb.ServerSecurityAlertPolicyStateDisabled),
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
	// res.ServerSecurityAlertPolicy = armmariadb.ServerSecurityAlertPolicy{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.DBforMariaDB/servers/securityAlertPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/securityalert-4799/providers/Microsoft.DBforMariaDB/servers/securityalert-6440/securityAlertPolicies/default"),
	// 	Properties: &armmariadb.SecurityAlertPolicyProperties{
	// 		DisabledAlerts: []*string{
	// 		},
	// 		EmailAccountAdmins: to.Ptr(true),
	// 		EmailAddresses: []*string{
	// 		},
	// 		RetentionDays: to.Ptr[int32](0),
	// 		State: to.Ptr(armmariadb.ServerSecurityAlertPolicyStateEnabled),
	// 		StorageEndpoint: to.Ptr(""),
	// 	},
	// }
}
