package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPSupportedSkus_Distributed.json
func ExampleClient_SAPSupportedSKU_sapSupportedSkusDistributed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().SAPSupportedSKU(ctx, "centralus", &armworkloads.ClientSAPSupportedSKUOptions{SAPSupportedSKU: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPSupportedResourceSKUsResult = armworkloads.SAPSupportedResourceSKUsResult{
	// 	SupportedSKUs: []*armworkloads.SAPSupportedSKU{
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(false),
	// 			VMSKU: to.Ptr("Standard_E4ds_v4"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(false),
	// 			VMSKU: to.Ptr("Standard_E8ds_v4"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(false),
	// 			VMSKU: to.Ptr("Standard_E16ds_v4"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(false),
	// 			VMSKU: to.Ptr("Standard_E20ds_v4"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(false),
	// 			VMSKU: to.Ptr("Standard_E32ds_v4"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(false),
	// 			VMSKU: to.Ptr("Standard_E48ds_v4"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(false),
	// 			VMSKU: to.Ptr("Standard_E64ds_v4"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M32Is"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M32ts"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M64Is"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M64ms"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M64s"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M128ms"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M128s"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M208ms_v2"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M208s_v2"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M416ms_v2"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(false),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M416s_v2"),
	// 	}},
	// }
}
