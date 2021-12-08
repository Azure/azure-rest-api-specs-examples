Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatacatalog%2Farmdatacatalog%2Fv0.1.0/sdk/resourcemanager/datacatalog/armdatacatalog/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatacatalog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datacatalog/armdatacatalog"
)

// x-ms-original-file: specification/datacatalog/resource-manager/Microsoft.DataCatalog/stable/2016-03-30/examples/UpdateADCCatalog.json
func ExampleADCCatalogsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatacatalog.NewADCCatalogsClient("<subscription-id>",
		"<catalog-name>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		armdatacatalog.ADCCatalog{
			Resource: armdatacatalog.Resource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"mykey":  to.StringPtr("myvalue"),
					"mykey2": to.StringPtr("myvalue2"),
				},
			},
			Properties: &armdatacatalog.ADCCatalogProperties{
				Admins: []*armdatacatalog.Principals{
					{
						ObjectID: to.StringPtr("<object-id>"),
						Upn:      to.StringPtr("<upn>"),
					}},
				EnableAutomaticUnitAdjustment: to.BoolPtr(false),
				SKU:                           armdatacatalog.SKUTypeStandard.ToPtr(),
				Units:                         to.Int32Ptr(1),
				Users: []*armdatacatalog.Principals{
					{
						ObjectID: to.StringPtr("<object-id>"),
						Upn:      to.StringPtr("<upn>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ADCCatalog.ID: %s\n", *res.ID)
}
```
