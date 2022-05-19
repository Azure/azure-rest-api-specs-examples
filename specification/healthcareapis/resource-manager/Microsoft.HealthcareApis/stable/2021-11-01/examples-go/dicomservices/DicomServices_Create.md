Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhealthcareapis%2Farmhealthcareapis%2Fv1.0.0/sdk/resourcemanager/healthcareapis/armhealthcareapis/README.md) on how to add the SDK to your project and authenticate.

```go
package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_Create.json
func ExampleDicomServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhealthcareapis.NewDicomServicesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testRG",
		"workspace1",
		"blue",
		armhealthcareapis.DicomService{
			Location:   to.Ptr("westus"),
			Properties: &armhealthcareapis.DicomServiceProperties{},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
