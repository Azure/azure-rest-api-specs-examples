package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/dataConnectors/GetDataConnectors.json
func ExampleDataConnectorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataConnectorsClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page.DataConnectorList = armsecurityinsights.DataConnectorList{
		// 	Value: []armsecurityinsights.DataConnectorClassification{
		// 		&armsecurityinsights.ASCDataConnector{
		// 			Name: to.Ptr("763f9fa1-c2d3-4fa2-93e9-bccd4899aa12"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/763f9fa1-c2d3-4fa2-93e9-bccd4899aa12"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindAzureSecurityCenter),
		// 			Properties: &armsecurityinsights.ASCDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
		// 					Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				SubscriptionID: to.Ptr("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0"),
		// 			},
		// 		},
		// 		&armsecurityinsights.TIDataConnector{
		// 			Name: to.Ptr("c345bf40-8509-4ed2-b947-50cb773aaf04"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/c345bf40-8509-4ed2-b947-50cb773aaf04"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindThreatIntelligence),
		// 			Properties: &armsecurityinsights.TIDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.TIDataConnectorDataTypes{
		// 					Indicators: &armsecurityinsights.TIDataConnectorDataTypesIndicators{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 		},
		// 		&armsecurityinsights.AADDataConnector{
		// 			Name: to.Ptr("f0cd27d2-5f03-4c06-ba31-d2dc82dcb51d"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/f0cd27d2-5f03-4c06-ba31-d2dc82dcb51d"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindAzureActiveDirectory),
		// 			Properties: &armsecurityinsights.AADDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
		// 					Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 		},
		// 		&armsecurityinsights.OfficeDataConnector{
		// 			Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindOffice365),
		// 			Properties: &armsecurityinsights.OfficeDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.OfficeDataConnectorDataTypes{
		// 					Exchange: &armsecurityinsights.OfficeDataConnectorDataTypesExchange{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 					SharePoint: &armsecurityinsights.OfficeDataConnectorDataTypesSharePoint{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 		},
		// 		&armsecurityinsights.MCASDataConnector{
		// 			Name: to.Ptr("b96d014d-b5c2-4a01-9aba-a8058f629d42"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/b96d014d-b5c2-4a01-9aba-a8058f629d42"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindMicrosoftCloudAppSecurity),
		// 			Properties: &armsecurityinsights.MCASDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.MCASDataConnectorDataTypes{
		// 					Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 					DiscoveryLogs: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 		},
		// 		&armsecurityinsights.AATPDataConnector{
		// 			Name: to.Ptr("07e42cb3-e658-4e90-801c-efa0f29d3d44"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/07e42cb3-e658-4e90-801c-efa0f29d3d44"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindAzureAdvancedThreatProtection),
		// 			Properties: &armsecurityinsights.AATPDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
		// 					Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 		},
		// 		&armsecurityinsights.AwsCloudTrailDataConnector{
		// 			Name: to.Ptr("c345bf40-8509-4ed2-b947-50cb773aaf04"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/c345bf40-8509-4ed2-b947-50cb773aaf04"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindAmazonWebServicesCloudTrail),
		// 			Properties: &armsecurityinsights.AwsCloudTrailDataConnectorProperties{
		// 				AwsRoleArn: to.Ptr("myAwsRoleArn"),
		// 				DataTypes: &armsecurityinsights.AwsCloudTrailDataConnectorDataTypes{
		// 					Logs: &armsecurityinsights.AwsCloudTrailDataConnectorDataTypesLogs{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		&armsecurityinsights.MDATPDataConnector{
		// 			Name: to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Kind: to.Ptr(armsecurityinsights.DataConnectorKindMicrosoftDefenderAdvancedThreatProtection),
		// 			Properties: &armsecurityinsights.MDATPDataConnectorProperties{
		// 				DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
		// 					Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
		// 						State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
		// 					},
		// 				},
		// 				TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
		// 			},
		// 	}},
		// }
	}
}
