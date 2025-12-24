package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be46becafeb29aa993898709e35759d3643b2809/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/DataExportGet.json
func ExampleDataExportsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataExportsClient().Get(ctx, "RgTest1", "DeWnTest1234", "export1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataExport = armoperationalinsights.DataExport{
	// 	Name: to.Ptr("export1"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces/export"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/RgTest1/providers/microsoft.operationalinsights/workspaces/DeWnTest1234/export/export1"),
	// 	Properties: &armoperationalinsights.DataExportProperties{
	// 		CreatedDate: to.Ptr("Sun, 12 Jan 2020 12:51:10 GMT"),
	// 		DataExportID: to.Ptr("d5233afc-7829-4b89-c594-08d7975e19a5"),
	// 		Destination: &armoperationalinsights.Destination{
	// 			Type: to.Ptr(armoperationalinsights.TypeEventHub),
	// 			ResourceID: to.Ptr("/subscriptions/192b9f85-a39a-4276-b96d-d5cd351703f9/resourceGroups/OIAutoRest1234/providers/Microsoft.EventHub/namespaces/test"),
	// 		},
	// 		Enable: to.Ptr(true),
	// 		LastModifiedDate: to.Ptr("Sun, 12 Jan 2020 12:51:10 GMT"),
	// 		TableNames: []*string{
	// 			to.Ptr("Heartbeat")},
	// 		},
	// 	}
}
