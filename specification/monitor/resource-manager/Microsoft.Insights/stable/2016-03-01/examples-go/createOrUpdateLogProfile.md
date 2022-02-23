Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.4.1/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/createOrUpdateLogProfile.json
func ExampleLogProfilesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewLogProfilesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<log-profile-name>",
		armmonitor.LogProfileResource{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armmonitor.LogProfileProperties{
				Categories: []*string{
					to.StringPtr("Write"),
					to.StringPtr("Delete"),
					to.StringPtr("Action")},
				Locations: []*string{
					to.StringPtr("global")},
				RetentionPolicy: &armmonitor.RetentionPolicy{
					Days:    to.Int32Ptr(3),
					Enabled: to.BoolPtr(true),
				},
				ServiceBusRuleID: to.StringPtr("<service-bus-rule-id>"),
				StorageAccountID: to.StringPtr("<storage-account-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LogProfilesClientCreateOrUpdateResult)
}
```
