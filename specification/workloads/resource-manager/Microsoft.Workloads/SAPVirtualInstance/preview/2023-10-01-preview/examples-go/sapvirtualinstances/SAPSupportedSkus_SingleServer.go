package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPSupportedSkus_SingleServer.json
func ExampleWorkloadsClient_SAPSupportedSKU_sapSupportedSkusSingleServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkloadsClient().SAPSupportedSKU(ctx, "centralus", &armworkloadssapvirtualinstance.WorkloadsClientSAPSupportedSKUOptions{SAPSupportedSKU: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SAPSupportedResourceSKUsResult = armworkloadssapvirtualinstance.SAPSupportedResourceSKUsResult{
	// 	SupportedSKUs: []*armworkloadssapvirtualinstance.SAPSupportedSKU{
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
	// 			IsAppServerCertified: to.Ptr(true),
	// 			IsDatabaseCertified: to.Ptr(true),
	// 			VMSKU: to.Ptr("Standard_M32ts"),
	// 	}},
	// }
}
