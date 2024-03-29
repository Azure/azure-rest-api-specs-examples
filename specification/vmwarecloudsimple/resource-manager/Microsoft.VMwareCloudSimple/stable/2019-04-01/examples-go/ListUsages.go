package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListUsages.json
func ExampleUsagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListPager("westus2", &armvmwarecloudsimple.UsagesClientListOptions{Filter: nil})
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
		// page.UsageListResponse = armvmwarecloudsimple.UsageListResponse{
		// 	Value: []*armvmwarecloudsimple.Usage{
		// 		{
		// 			Name: &armvmwarecloudsimple.UsageName{
		// 				LocalizedValue: to.Ptr("CS28-Node"),
		// 				Value: to.Ptr("general"),
		// 			},
		// 			CurrentValue: to.Ptr[int32](1),
		// 			Limit: to.Ptr[int32](5),
		// 			Unit: to.Ptr(armvmwarecloudsimple.UsageCountCount),
		// 		},
		// 		{
		// 			Name: &armvmwarecloudsimple.UsageName{
		// 				LocalizedValue: to.Ptr("CS36-Node"),
		// 				Value: to.Ptr("large"),
		// 			},
		// 			CurrentValue: to.Ptr[int32](0),
		// 			Limit: to.Ptr[int32](5),
		// 			Unit: to.Ptr(armvmwarecloudsimple.UsageCountCount),
		// 	}},
		// }
	}
}
