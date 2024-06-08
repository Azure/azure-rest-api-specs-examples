package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f6f50c6388fd5836fa142384641b8353a99874ef/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ChangeDataCapture_ListByFactory.json
func ExampleChangeDataCaptureClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewChangeDataCaptureClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.ChangeDataCaptureListResponse = armdatafactory.ChangeDataCaptureListResponse{
		// 	Value: []*armdatafactory.ChangeDataCaptureResource{
		// 		{
		// 			Name: to.Ptr("exampleChangeDataCapture"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/adfcdcs"),
		// 			Etag: to.Ptr("4200eefe-0000-0100-0000-641aa97a0000"),
		// 			ID: to.Ptr("/subscriptions/d3bb3b2e-9a7e-4194-9960-5171bd192117/resourceGroups/amja-rg-03/providers/Microsoft.DataFactory/factories/amja-adf-04/adfcdcs/exampleChangeDataCapture"),
		// 			Properties: &armdatafactory.ChangeDataCapture{
		// 				Description: to.Ptr("Sample demo change data capture to transfer data from delimited (csv) to Azure SQL Database with automapped and non-automapped mappings."),
		// 				AllowVNetOverride: to.Ptr(false),
		// 				Status: to.Ptr("Stopped"),
		// 			},
		// 	}},
		// }
	}
}
