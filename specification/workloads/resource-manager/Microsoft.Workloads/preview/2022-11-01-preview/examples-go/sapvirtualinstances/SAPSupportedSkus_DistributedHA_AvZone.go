package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPSupportedSkus_DistributedHA_AvZone.json
func ExampleClient_SAPSupportedSKU_sapSupportedSkusDistributedHaAvZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armworkloads.NewClient("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.SAPSupportedSKU(ctx, "centralus", &armworkloads.ClientSAPSupportedSKUOptions{SAPSupportedSKU: nil})
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
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_E32ds_v4"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_E48ds_v4"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_E64ds_v4"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M32Is"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M32ts"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M64Is"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M64ms"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M64s"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M128ms"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M128s"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M208ms_v2"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M208s_v2"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M416ms_v2"),
	// 		},
	// 		{
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M416s_v2"),
	// 	}},
	// }
}
