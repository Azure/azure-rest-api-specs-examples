package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/dataConnectors/CreatePurviewAuditDataConnector.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesAPurviewAuditDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.PurviewAuditDataConnector{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindPurviewAudit),
		Properties: &armsecurityinsights.PurviewAuditDataConnectorProperties{
			ConnectorDefinitionName: to.Ptr("PowerAutomate"),
			DataTypes: &armsecurityinsights.PurviewAuditConnectorDataTypes{
				Logs: &armsecurityinsights.PurviewAuditConnectorDataTypesLogs{
					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
				},
			},
			DcrConfig: &armsecurityinsights.DCRConfiguration{
				DataCollectionEndpoint:        to.Ptr("https://microsoft-sentinel-datacollectionendpoint-123m.westeurope-1.ingest.monitor.azure.com"),
				DataCollectionRuleImmutableID: to.Ptr("dcr-de21b053bd5a44beb99a256c9db85023"),
				StreamName:                    to.Ptr("OFFICEPOWERAUTOMATE_RESTAPI"),
			},
			SourceType: to.Ptr("MicrosoftFlow"),
			TenantID:   to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientCreateOrUpdateResponse{
	// 	DataConnectorClassification: &armsecurityinsights.PurviewAuditDataConnector{
	// 		Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindPurviewAudit),
	// 		Properties: &armsecurityinsights.PurviewAuditDataConnectorProperties{
	// 			ConnectorDefinitionName: to.Ptr("PowerAutomate"),
	// 			DataTypes: &armsecurityinsights.PurviewAuditConnectorDataTypes{
	// 				Logs: &armsecurityinsights.PurviewAuditConnectorDataTypesLogs{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			SourceType: to.Ptr("MicrosoftFlow"),
	// 			TenantID: to.Ptr("MicrosoftFlow"),
	// 		},
	// 	},
	// }
}
