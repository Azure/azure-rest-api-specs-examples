package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/AppServiceEnvironments_ListMultiRolePoolInstanceMetricDefinitions.json
func ExampleEnvironmentsClient_NewListMultiRolePoolInstanceMetricDefinitionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnvironmentsClient().NewListMultiRolePoolInstanceMetricDefinitionsPager("test-rg", "test-ase", "10.7.1.8", nil)
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
		// page.ResourceMetricDefinitionCollection = armappservice.ResourceMetricDefinitionCollection{
		// 	Value: []*armappservice.ResourceMetricDefinition{
		// 		{
		// 			Name: to.Ptr("CpuPercentage"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/metricdefinitions/cpupercentage"),
		// 		},
		// 		{
		// 			Name: to.Ptr("MemoryPercentage"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/metricdefinitions/MemoryPercentage"),
		// 		},
		// 		{
		// 			Name: to.Ptr("DiskQueueLength"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/metricdefinitions/DiskQueueLength"),
		// 		},
		// 		{
		// 			Name: to.Ptr("HttpQueueLength"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/metricdefinitions/HttpQueueLength"),
		// 		},
		// 		{
		// 			Name: to.Ptr("BytesReceived"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/metricdefinitions/BytesReceived"),
		// 		},
		// 		{
		// 			Name: to.Ptr("BytesSent"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/metricdefinitions/BytesSent"),
		// 	}},
		// }
	}
}
