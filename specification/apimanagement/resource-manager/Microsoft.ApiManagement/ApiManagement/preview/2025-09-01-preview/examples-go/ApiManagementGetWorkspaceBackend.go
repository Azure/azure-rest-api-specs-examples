package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementGetWorkspaceBackend.json
func ExampleWorkspaceBackendClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceBackendClient().Get(ctx, "rg1", "apimService1", "wks1", "sfbackend", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.WorkspaceBackendClientGetResponse{
	// 	BackendContract: armapimanagement.BackendContract{
	// 		Name: to.Ptr("sfbackend"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/backends"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/backends/sfbackend"),
	// 		Properties: &armapimanagement.BackendContractProperties{
	// 			Description: to.Ptr("Service Fabric Test App 1"),
	// 			Properties: &armapimanagement.BackendProperties{
	// 				ServiceFabricCluster: &armapimanagement.BackendServiceFabricClusterProperties{
	// 					ClientCertificateID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/certificates/cert1"),
	// 					ManagementEndpoints: []*string{
	// 						to.Ptr("https://somecluster.com"),
	// 					},
	// 					MaxPartitionResolutionRetries: to.Ptr[int32](5),
	// 					ServerX509Names: []*armapimanagement.X509CertificateName{
	// 						{
	// 							Name: to.Ptr("ServerCommonName1"),
	// 							IssuerCertificateThumbprint: to.Ptr("IssuerCertificateThumbprint1"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			URL: to.Ptr("fabric:/mytestapp/mytestservice"),
	// 			Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
	// 		},
	// 	},
	// }
}
