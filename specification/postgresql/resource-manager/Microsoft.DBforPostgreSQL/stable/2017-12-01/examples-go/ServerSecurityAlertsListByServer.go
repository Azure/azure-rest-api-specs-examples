package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerSecurityAlertsListByServer.json
func ExampleServerSecurityAlertPoliciesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerSecurityAlertPoliciesClient().NewListByServerPager("securityalert-4799", "securityalert-6440", nil)
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
		// page.ServerSecurityAlertPolicyListResult = armpostgresql.ServerSecurityAlertPolicyListResult{
		// 	Value: []*armpostgresql.ServerSecurityAlertPolicy{
		// 		{
		// 			Name: to.Ptr("Default"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/securityAlertPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/securityalert-4799/providers/Microsoft.DBforPostgreSQL/servers/securityalert-6440/securityAlertPolicies"),
		// 			Properties: &armpostgresql.SecurityAlertPolicyProperties{
		// 				DisabledAlerts: []*string{
		// 					to.Ptr("Access_Anomaly")},
		// 					EmailAccountAdmins: to.Ptr(true),
		// 					EmailAddresses: []*string{
		// 						to.Ptr("test@microsoft.com;user@microsoft.com")},
		// 						RetentionDays: to.Ptr[int32](0),
		// 						State: to.Ptr(armpostgresql.ServerSecurityAlertPolicyStateDisabled),
		// 						StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
		// 					},
		// 			}},
		// 		}
	}
}
