package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolSecurityAlertWithAllParameters.json
func ExampleSQLPoolSecurityAlertPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSecurityAlertPoliciesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"securityalert-4799",
		"securityalert-6440",
		"testdb",
		armsynapse.SecurityAlertPolicyNameDefault,
		armsynapse.SQLPoolSecurityAlertPolicy{
			Properties: &armsynapse.SecurityAlertPolicyProperties{
				DisabledAlerts: []*string{
					to.Ptr("Sql_Injection"),
					to.Ptr("Usage_Anomaly")},
				EmailAccountAdmins: to.Ptr(true),
				EmailAddresses: []*string{
					to.Ptr("test@microsoft.com"),
					to.Ptr("user@microsoft.com")},
				RetentionDays:           to.Ptr[int32](6),
				State:                   to.Ptr(armsynapse.SecurityAlertPolicyStateEnabled),
				StorageAccountAccessKey: to.Ptr("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
				StorageEndpoint:         to.Ptr("https://mystorage.blob.core.windows.net"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
