package armdatacatalog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datacatalog/armdatacatalog"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/datacatalog/resource-manager/Microsoft.DataCatalog/stable/2016-03-30/examples/CreateOrUpdateADCCatalog.json
func ExampleADCCatalogsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatacatalog.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewADCCatalogsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleCatalog", armdatacatalog.ADCCatalog{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ADCCatalog = armdatacatalog.ADCCatalog{
	// 	Name: to.Ptr("exampleCatalog"),
	// 	Type: to.Ptr("Microsoft.DataCatalog/catalogs"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataCatalog/catalogs/exampleCatalog"),
	// 	Location: to.Ptr("North US"),
	// 	Tags: map[string]*string{
	// 		"mykey": to.Ptr("myvalue"),
	// 		"mykey2": to.Ptr("myvalue2"),
	// 	},
	// 	Properties: &armdatacatalog.ADCCatalogProperties{
	// 		Admins: []*armdatacatalog.Principals{
	// 			{
	// 				ObjectID: to.Ptr("99999999-9999-9999-999999999999"),
	// 				Upn: to.Ptr("myupn@microsoft.com"),
	// 		}},
	// 		EnableAutomaticUnitAdjustment: to.Ptr(false),
	// 		SKU: to.Ptr(armdatacatalog.SKUTypeStandard),
	// 		SuccessfullyProvisioned: to.Ptr(true),
	// 		Units: to.Ptr[int32](1),
	// 		Users: []*armdatacatalog.Principals{
	// 			{
	// 				ObjectID: to.Ptr("99999999-9999-9999-999999999999"),
	// 				Upn: to.Ptr("myupn@microsoft.com"),
	// 		}},
	// 	},
	// }
}
