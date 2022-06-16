package armdatacatalog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datacatalog/armdatacatalog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datacatalog/resource-manager/Microsoft.DataCatalog/stable/2016-03-30/examples/UpdateADCCatalog.json
func ExampleADCCatalogsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatacatalog.NewADCCatalogsClient("12345678-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"exampleResourceGroup",
		"exampleCatalog",
		armdatacatalog.ADCCatalog{
			Location: to.Ptr("North US"),
			Tags: map[string]*string{
				"mykey":  to.Ptr("myvalue"),
				"mykey2": to.Ptr("myvalue2"),
			},
			Properties: &armdatacatalog.ADCCatalogProperties{
				Admins: []*armdatacatalog.Principals{
					{
						ObjectID: to.Ptr("99999999-9999-9999-999999999999"),
						Upn:      to.Ptr("myupn@microsoft.com"),
					}},
				EnableAutomaticUnitAdjustment: to.Ptr(false),
				SKU:                           to.Ptr(armdatacatalog.SKUTypeStandard),
				Units:                         to.Ptr[int32](1),
				Users: []*armdatacatalog.Principals{
					{
						ObjectID: to.Ptr("99999999-9999-9999-999999999999"),
						Upn:      to.Ptr("myupn@microsoft.com"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
