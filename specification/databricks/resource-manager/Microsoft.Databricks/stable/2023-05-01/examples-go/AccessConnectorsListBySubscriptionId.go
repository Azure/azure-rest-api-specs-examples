package armdatabricks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databricks/armdatabricks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e1a87e1a5deb3f986ea1474d233d6680f1e19fc1/specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-05-01/examples/AccessConnectorsListBySubscriptionId.json
func ExampleAccessConnectorsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabricks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccessConnectorsClient().NewListBySubscriptionPager(nil)
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
		// page.AccessConnectorListResult = armdatabricks.AccessConnectorListResult{
		// 	Value: []*armdatabricks.AccessConnector{
		// 		{
		// 			Name: to.Ptr("myAccessConnector1"),
		// 			Type: to.Ptr("Microsoft.Databricks/accessConnectors"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Databricks/accessConnectors/myAccessConnector1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatabricks.ManagedServiceIdentity{
		// 				Type: to.Ptr(armdatabricks.ManagedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("7ad2bae1-37d0-413e-91f8-b0b7bef807fc"),
		// 				TenantID: to.Ptr("e3fe3f22-4b98-4c04-82cc-d8817d1b17da"),
		// 			},
		// 			Properties: &armdatabricks.AccessConnectorProperties{
		// 				ProvisioningState: to.Ptr(armdatabricks.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAccessConnector2"),
		// 			Type: to.Ptr("Microsoft.Databricks/accessConnectors"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Databricks/accessConnectors/myAccessConnector2"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatabricks.ManagedServiceIdentity{
		// 				Type: to.Ptr(armdatabricks.ManagedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("5619ff16-afe1-47e5-ae67-8393c6c3223d"),
		// 				TenantID: to.Ptr("e3fe3f22-4b98-4c04-82cc-d8817d1b17da"),
		// 			},
		// 			Properties: &armdatabricks.AccessConnectorProperties{
		// 				ProvisioningState: to.Ptr(armdatabricks.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAccessConnector3"),
		// 			Type: to.Ptr("Microsoft.Databricks/accessConnectors"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Databricks/accessConnectors/myAccessConnector3"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatabricks.ManagedServiceIdentity{
		// 				Type: to.Ptr(armdatabricks.ManagedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("04b25430-8db1-48a0-9c2f-32270ed63eef"),
		// 				TenantID: to.Ptr("e3fe3f22-4b98-4c04-82cc-d8817d1b17da"),
		// 			},
		// 			Properties: &armdatabricks.AccessConnectorProperties{
		// 				ProvisioningState: to.Ptr(armdatabricks.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAccessConnector4"),
		// 			Type: to.Ptr("Microsoft.Databricks/accessConnectors"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg3/providers/Microsoft.Databricks/accessConnectors/myAccessConnector4"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armdatabricks.ManagedServiceIdentity{
		// 				Type: to.Ptr(armdatabricks.ManagedServiceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("4856ceed-0a99-4df7-b9fc-35603650af06"),
		// 				TenantID: to.Ptr("e3fe3f22-4b98-4c04-82cc-d8817d1b17da"),
		// 			},
		// 			Properties: &armdatabricks.AccessConnectorProperties{
		// 				ProvisioningState: to.Ptr(armdatabricks.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
