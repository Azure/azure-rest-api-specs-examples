package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be0d1c383d683a8eb22ab5fd5b75e380ac3c2eea/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ChangeDataCapture_Update.json
func ExampleChangeDataCaptureClient_CreateOrUpdate_changeDataCaptureUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChangeDataCaptureClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleChangeDataCapture", armdatafactory.ChangeDataCaptureResource{
		Properties: &armdatafactory.ChangeDataCapture{
			Description:       to.Ptr("Sample demo change data capture to transfer data from delimited (csv) to Azure SQL Database. Updating table mappings."),
			AllowVNetOverride: to.Ptr(false),
			Status:            to.Ptr("Stopped"),
		},
	}, &armdatafactory.ChangeDataCaptureClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ChangeDataCaptureResource = armdatafactory.ChangeDataCaptureResource{
	// 	Name: to.Ptr("exampleChangeDataCapture"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/adfcdcs"),
	// 	Etag: to.Ptr("43008e18-0000-0100-0000-641aafca0000"),
	// 	ID: to.Ptr("/subscriptions/d3bb3b2e-9a7e-4194-9960-5171bd192117/resourceGroups/amja-rg-03/providers/Microsoft.DataFactory/factories/amja-adf-04/adfcdcs/exampleChangeDataCapture"),
	// 	Properties: &armdatafactory.ChangeDataCapture{
	// 		Description: to.Ptr("Sample demo change data capture to transfer data from delimited (csv) to Azure SQL Database. Updating table mappings."),
	// 		AllowVNetOverride: to.Ptr(false),
	// 		Status: to.Ptr("Stopped"),
	// 	},
	// }
}
