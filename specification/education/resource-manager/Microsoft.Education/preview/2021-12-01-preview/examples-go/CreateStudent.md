```go
package armeducation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/CreateStudent.json
func ExampleStudentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeducation.NewStudentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"{billingAccountName}",
		"{billingProfileName}",
		"{invoiceSectionName}",
		"{studentAlias}",
		armeducation.StudentDetails{
			Properties: &armeducation.StudentProperties{
				Budget: &armeducation.Amount{
					Currency: to.Ptr("USD"),
					Value:    to.Ptr[float32](100),
				},
				Email:          to.Ptr("test@contoso.com"),
				ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-09T22:13:21.795Z"); return t }()),
				FirstName:      to.Ptr("test"),
				LastName:       to.Ptr("user"),
				Role:           to.Ptr(armeducation.StudentRoleStudent),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feducation%2Farmeducation%2Fv0.1.0/sdk/resourcemanager/education/armeducation/README.md) on how to add the SDK to your project and authenticate.
