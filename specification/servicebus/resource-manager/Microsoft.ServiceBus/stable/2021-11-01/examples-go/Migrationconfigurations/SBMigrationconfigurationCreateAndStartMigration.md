Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicebus%2Farmservicebus%2Fv0.3.1/sdk/resourcemanager/servicebus/armservicebus/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicebus_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Migrationconfigurations/SBMigrationconfigurationCreateAndStartMigration.json
func ExampleMigrationConfigsClient_BeginCreateAndStartMigration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicebus.NewMigrationConfigsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateAndStartMigration(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		armservicebus.MigrationConfigurationName("$default"),
		armservicebus.MigrationConfigProperties{
			Properties: &armservicebus.MigrationConfigPropertiesProperties{
				PostMigrationName: to.StringPtr("<post-migration-name>"),
				TargetNamespace:   to.StringPtr("<target-namespace>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MigrationConfigsClientCreateAndStartMigrationResult)
}
```
