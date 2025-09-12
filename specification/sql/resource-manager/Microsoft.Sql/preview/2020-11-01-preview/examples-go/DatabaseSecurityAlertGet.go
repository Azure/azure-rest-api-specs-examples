package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8358c7473dfe057d84a6b6a921225063c040b31a/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseSecurityAlertGet.json
func ExampleDatabaseSecurityAlertPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseSecurityAlertPoliciesClient().Get(ctx, "securityalert-6852", "securityalert-2080", "testdb", armsql.SecurityAlertPolicyNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DatabaseSecurityAlertPolicy = armsql.DatabaseSecurityAlertPolicy{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/securityAlertPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/securityalert-6852/providers/Microsoft.Sql/servers/securityalert-2080/databases/testdb"),
	// 	Properties: &armsql.SecurityAlertsPolicyProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-03T04:41:33.937Z"); return t}()),
	// 		DisabledAlerts: []*string{
	// 			to.Ptr("Usage_Anomaly")},
	// 			EmailAccountAdmins: to.Ptr(true),
	// 			EmailAddresses: []*string{
	// 				to.Ptr("test@consoto.com"),
	// 				to.Ptr("user@consoto.com")},
	// 				RetentionDays: to.Ptr[int32](0),
	// 				State: to.Ptr(armsql.SecurityAlertsPolicyStateEnabled),
	// 				StorageAccountAccessKey: to.Ptr(""),
	// 			},
	// 			SystemData: &armsql.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-03T04:41:33.937Z"); return t}()),
	// 				CreatedBy: to.Ptr("string"),
	// 				CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-03T04:41:33.937Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("string"),
	// 				LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 			},
	// 		}
}
