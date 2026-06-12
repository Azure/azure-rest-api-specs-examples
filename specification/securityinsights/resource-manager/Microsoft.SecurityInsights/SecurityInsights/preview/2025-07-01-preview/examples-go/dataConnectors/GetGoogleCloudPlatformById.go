package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/dataConnectors/GetGoogleCloudPlatformById.json
func ExampleDataConnectorsClient_Get_getAGcpDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "GCP_fce27b90-d6f5-4d30-991a-af509a2b50a1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	DataConnectorClassification: &armsecurityinsights.GCPDataConnector{
	// 		Name: to.Ptr("GCP_fce27b90-d6f5-4d30-991a-af509a2b50a1"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/GCP_afef3743-0c88-469c-84ff-ca2e87dc1e48"),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindGCP),
	// 		Properties: &armsecurityinsights.GCPDataConnectorProperties{
	// 			Auth: &armsecurityinsights.GCPAuthProperties{
	// 				ProjectNumber: to.Ptr("123456789012"),
	// 				ServiceAccountEmail: to.Ptr("sentinel-service-account@project-id.iam.gserviceaccount.com"),
	// 				WorkloadIdentityProviderID: to.Ptr("sentinel-identity-provider"),
	// 			},
	// 			ConnectorDefinitionName: to.Ptr("GcpConnector"),
	// 			Request: &armsecurityinsights.GCPRequestProperties{
	// 				ProjectID: to.Ptr("project-id"),
	// 				SubscriptionNames: []*string{
	// 					to.Ptr("sentinel-subscription"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
