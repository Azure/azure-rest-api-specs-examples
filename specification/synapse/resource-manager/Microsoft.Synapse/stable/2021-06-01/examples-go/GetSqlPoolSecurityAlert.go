package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolSecurityAlert.json
func ExampleSQLPoolSecurityAlertPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLPoolSecurityAlertPoliciesClient().Get(ctx, "securityalert-6852", "securityalert-2080", "testdb", armsynapse.SecurityAlertPolicyNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLPoolSecurityAlertPolicy = armsynapse.SQLPoolSecurityAlertPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/securityAlertPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/securityalert-6852/providers/Microsoft.Synapse/workspaces/securityalert-2080/sqlPools/testdb"),
	// 	Properties: &armsynapse.SecurityAlertPolicyProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-08T00:00:00.000Z"); return t}()),
	// 		DisabledAlerts: []*string{
	// 			to.Ptr("Usage_Anomaly")},
	// 			EmailAccountAdmins: to.Ptr(true),
	// 			EmailAddresses: []*string{
	// 				to.Ptr("test@microsoft.com"),
	// 				to.Ptr("user@microsoft.com")},
	// 				RetentionDays: to.Ptr[int32](0),
	// 				State: to.Ptr(armsynapse.SecurityAlertPolicyStateEnabled),
	// 				StorageAccountAccessKey: to.Ptr(""),
	// 			},
	// 		}
}
