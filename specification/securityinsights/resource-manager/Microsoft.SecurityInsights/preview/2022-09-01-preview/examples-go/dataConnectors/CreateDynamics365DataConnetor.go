package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CreateDynamics365DataConnetor.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesADynamics365DataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewDataConnectorsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "c2541efb-c9a6-47fe-9501-87d1017d1512", &armsecurityinsights.Dynamics365DataConnector{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindDynamics365),
		Properties: &armsecurityinsights.Dynamics365DataConnectorProperties{
			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
			DataTypes: &armsecurityinsights.Dynamics365DataConnectorDataTypes{
				Dynamics365CdsActivities: &armsecurityinsights.Dynamics365DataConnectorDataTypesDynamics365CdsActivities{
					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
