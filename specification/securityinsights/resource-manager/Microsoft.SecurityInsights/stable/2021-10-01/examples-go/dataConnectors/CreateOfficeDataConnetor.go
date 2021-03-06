package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/CreateOfficeDataConnetor.json
func ExampleDataConnectorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewDataConnectorsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myRg",
		"myWorkspace",
		"73e01a99-5cd7-4139-a149-9f2736ff2ab5",
		&armsecurityinsights.OfficeDataConnector{
			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
			Kind: to.Ptr(armsecurityinsights.DataConnectorKindOffice365),
			Properties: &armsecurityinsights.OfficeDataConnectorProperties{
				DataTypes: &armsecurityinsights.OfficeDataConnectorDataTypes{
					Exchange: &armsecurityinsights.OfficeDataConnectorDataTypesExchange{
						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
					},
					SharePoint: &armsecurityinsights.OfficeDataConnectorDataTypesSharePoint{
						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
					},
					Teams: &armsecurityinsights.OfficeDataConnectorDataTypesTeams{
						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
					},
				},
				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
