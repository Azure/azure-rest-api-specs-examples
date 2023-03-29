package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSecurityAlertCreateMin.json
func ExampleManagedDatabaseSecurityAlertPoliciesClient_CreateOrUpdate_updateADatabasesThreatDetectionPolicyWithMinimalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedDatabaseSecurityAlertPoliciesClient().CreateOrUpdate(ctx, "securityalert-4799", "securityalert-6440", "testdb", armsql.SecurityAlertPolicyNameDefault, armsql.ManagedDatabaseSecurityAlertPolicy{
		Properties: &armsql.SecurityAlertPolicyProperties{
			State: to.Ptr(armsql.SecurityAlertPolicyStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedDatabaseSecurityAlertPolicy = armsql.ManagedDatabaseSecurityAlertPolicy{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstance/databases/securityAlertPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/securityalert-4799/providers/Microsoft.Sql/managedInstances/securityalert-6440/databases/testdb"),
	// 	Properties: &armsql.SecurityAlertPolicyProperties{
	// 		DisabledAlerts: []*string{
	// 		},
	// 		EmailAccountAdmins: to.Ptr(true),
	// 		EmailAddresses: []*string{
	// 		},
	// 		RetentionDays: to.Ptr[int32](0),
	// 		State: to.Ptr(armsql.SecurityAlertPolicyStateEnabled),
	// 		StorageAccountAccessKey: to.Ptr(""),
	// 		StorageEndpoint: to.Ptr(""),
	// 	},
	// }
}
