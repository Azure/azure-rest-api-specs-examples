package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/SapVirtualInstances_InvokeSapSupportedSku_Distributed.json
func ExampleSAPVirtualInstancesClient_GetSapSupportedSKU_sapSupportedSkUsForDistributedNonHaEnvironment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSAPVirtualInstancesClient().GetSapSupportedSKU(ctx, "centralus", armworkloadssapvirtualinstance.SAPSupportedSKUsRequest{
		AppLocation:    to.Ptr("eastus"),
		SapProduct:     to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
		Environment:    to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeProd),
		DatabaseType:   to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
		DeploymentType: to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armworkloadssapvirtualinstance.SAPVirtualInstancesClientGetSapSupportedSKUResponse{
	// 	SAPSupportedResourceSKUsResult: &armworkloadssapvirtualinstance.SAPSupportedResourceSKUsResult{
	// 		SupportedSKUs: []*armworkloadssapvirtualinstance.SAPSupportedSKU{
	// 			{
	// 				IsAppServerCertified: to.Ptr(true),
	// 				IsDatabaseCertified: to.Ptr(false),
	// 				VMSKU: to.Ptr("Standard_E4ds_v4"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(true),
	// 				IsDatabaseCertified: to.Ptr(false),
	// 				VMSKU: to.Ptr("Standard_E8ds_v4"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(true),
	// 				IsDatabaseCertified: to.Ptr(false),
	// 				VMSKU: to.Ptr("Standard_E16ds_v4"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(true),
	// 				IsDatabaseCertified: to.Ptr(false),
	// 				VMSKU: to.Ptr("Standard_E20ds_v4"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(true),
	// 				IsDatabaseCertified: to.Ptr(false),
	// 				VMSKU: to.Ptr("Standard_E32ds_v4"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(true),
	// 				IsDatabaseCertified: to.Ptr(false),
	// 				VMSKU: to.Ptr("Standard_E48ds_v4"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(true),
	// 				IsDatabaseCertified: to.Ptr(false),
	// 				VMSKU: to.Ptr("Standard_E64ds_v4"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M32Is"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M32ts"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M64Is"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M64ms"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M64s"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M128ms"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M128s"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M208ms_v2"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M208s_v2"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M416ms_v2"),
	// 			},
	// 			{
	// 				IsAppServerCertified: to.Ptr(false),
	// 				IsDatabaseCertified: to.Ptr(true),
	// 				VMSKU: to.Ptr("Standard_M416s_v2"),
	// 			},
	// 		},
	// 	},
	// }
}
