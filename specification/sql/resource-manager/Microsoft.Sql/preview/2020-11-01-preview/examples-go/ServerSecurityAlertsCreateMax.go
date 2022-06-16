package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerSecurityAlertsCreateMax.json
func ExampleServerSecurityAlertPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewServerSecurityAlertPoliciesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"securityalert-4799",
		"securityalert-6440",
		armsql.SecurityAlertPolicyNameDefault,
		armsql.ServerSecurityAlertPolicy{
			Properties: &armsql.SecurityAlertsPolicyProperties{
				DisabledAlerts: []*string{
					to.Ptr("Access_Anomaly"),
					to.Ptr("Usage_Anomaly")},
				EmailAccountAdmins: to.Ptr(true),
				EmailAddresses: []*string{
					to.Ptr("testSecurityAlert@microsoft.com")},
				RetentionDays:           to.Ptr[int32](5),
				State:                   to.Ptr(armsql.SecurityAlertsPolicyStateEnabled),
				StorageAccountAccessKey: to.Ptr("sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD=="),
				StorageEndpoint:         to.Ptr("https://mystorage.blob.core.windows.net"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
