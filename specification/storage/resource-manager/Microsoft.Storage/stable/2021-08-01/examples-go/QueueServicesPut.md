```go
package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/QueueServicesPut.json
func ExampleQueueServicesClient_SetServiceProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewQueueServicesClient("<subscription-id>", cred, nil)
	res, err := client.SetServiceProperties(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.QueueServiceProperties{
			QueueServiceProperties: &armstorage.QueueServicePropertiesProperties{
				Cors: &armstorage.CorsRules{
					CorsRules: []*armstorage.CorsRule{
						{
							AllowedHeaders: []*string{
								to.StringPtr("x-ms-meta-abc"),
								to.StringPtr("x-ms-meta-data*"),
								to.StringPtr("x-ms-meta-target*")},
							AllowedMethods: []*armstorage.CorsRuleAllowedMethodsItem{
								armstorage.CorsRuleAllowedMethodsItem("GET").ToPtr(),
								armstorage.CorsRuleAllowedMethodsItem("HEAD").ToPtr(),
								armstorage.CorsRuleAllowedMethodsItem("POST").ToPtr(),
								armstorage.CorsRuleAllowedMethodsItem("OPTIONS").ToPtr(),
								armstorage.CorsRuleAllowedMethodsItem("MERGE").ToPtr(),
								armstorage.CorsRuleAllowedMethodsItem("PUT").ToPtr()},
							AllowedOrigins: []*string{
								to.StringPtr("http://www.contoso.com"),
								to.StringPtr("http://www.fabrikam.com")},
							ExposedHeaders: []*string{
								to.StringPtr("x-ms-meta-*")},
							MaxAgeInSeconds: to.Int32Ptr(100),
						},
						{
							AllowedHeaders: []*string{
								to.StringPtr("*")},
							AllowedMethods: []*armstorage.CorsRuleAllowedMethodsItem{
								armstorage.CorsRuleAllowedMethodsItem("GET").ToPtr()},
							AllowedOrigins: []*string{
								to.StringPtr("*")},
							ExposedHeaders: []*string{
								to.StringPtr("*")},
							MaxAgeInSeconds: to.Int32Ptr(2),
						},
						{
							AllowedHeaders: []*string{
								to.StringPtr("x-ms-meta-12345675754564*")},
							AllowedMethods: []*armstorage.CorsRuleAllowedMethodsItem{
								armstorage.CorsRuleAllowedMethodsItem("GET").ToPtr(),
								armstorage.CorsRuleAllowedMethodsItem("PUT").ToPtr()},
							AllowedOrigins: []*string{
								to.StringPtr("http://www.abc23.com"),
								to.StringPtr("https://www.fabrikam.com/*")},
							ExposedHeaders: []*string{
								to.StringPtr("x-ms-meta-abc"),
								to.StringPtr("x-ms-meta-data*"),
								to.StringPtr("x-ms-meta-target*")},
							MaxAgeInSeconds: to.Int32Ptr(2000),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.QueueServicesClientSetServicePropertiesResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstorage%2Farmstorage%2Fv0.4.1/sdk/resourcemanager/storage/armstorage/README.md) on how to add the SDK to your project and authenticate.
