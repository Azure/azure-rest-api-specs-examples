package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateBackendServiceFabric.json
func ExampleBackendClient_CreateOrUpdate_apiManagementCreateBackendServiceFabric() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackendClient().CreateOrUpdate(ctx, "rg1", "apimService1", "sfbackend", armapimanagement.BackendContract{
		Properties: &armapimanagement.BackendContractProperties{
			Description: to.Ptr("Service Fabric Test App 1"),
			Properties: &armapimanagement.BackendProperties{
				ServiceFabricCluster: &armapimanagement.BackendServiceFabricClusterProperties{
					ClientCertificateID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
					ManagementEndpoints: []*string{
						to.Ptr("https://somecluster.com")},
					MaxPartitionResolutionRetries: to.Ptr[int32](5),
					ServerX509Names: []*armapimanagement.X509CertificateName{
						{
							Name:                        to.Ptr("ServerCommonName1"),
							IssuerCertificateThumbprint: to.Ptr("IssuerCertificateThumbprint1"),
						}},
				},
			},
			URL:      to.Ptr("fabric:/mytestapp/mytestservice"),
			Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
		},
	}, &armapimanagement.BackendClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackendContract = armapimanagement.BackendContract{
	// 	Name: to.Ptr("sfbackend"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/sfbackend"),
	// 	Properties: &armapimanagement.BackendContractProperties{
	// 		Description: to.Ptr("Service Fabric Test App 1"),
	// 		Properties: &armapimanagement.BackendProperties{
	// 			ServiceFabricCluster: &armapimanagement.BackendServiceFabricClusterProperties{
	// 				ClientCertificateID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
	// 				ManagementEndpoints: []*string{
	// 					to.Ptr("https://somecluster.com")},
	// 					MaxPartitionResolutionRetries: to.Ptr[int32](5),
	// 					ServerX509Names: []*armapimanagement.X509CertificateName{
	// 						{
	// 							Name: to.Ptr("ServerCommonName1"),
	// 							IssuerCertificateThumbprint: to.Ptr("IssuerCertificateThumbprint1"),
	// 					}},
	// 				},
	// 			},
	// 			URL: to.Ptr("fabric:/mytestapp/mytestservice"),
	// 			Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
	// 		},
	// 	}
}
