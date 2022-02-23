Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseAutomaticTuningUpdateMax.json
func ExampleDatabaseAutomaticTuningClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewDatabaseAutomaticTuningClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		armsql.DatabaseAutomaticTuning{
			Properties: &armsql.DatabaseAutomaticTuningProperties{
				DesiredState: armsql.AutomaticTuningModeAuto.ToPtr(),
				Options: map[string]*armsql.AutomaticTuningOptions{
					"createIndex": {
						DesiredState: armsql.AutomaticTuningOptionModeDesiredOff.ToPtr(),
					},
					"dropIndex": {
						DesiredState: armsql.AutomaticTuningOptionModeDesiredOn.ToPtr(),
					},
					"forceLastGoodPlan": {
						DesiredState: armsql.AutomaticTuningOptionModeDesiredDefault.ToPtr(),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DatabaseAutomaticTuningClientUpdateResult)
}
```
