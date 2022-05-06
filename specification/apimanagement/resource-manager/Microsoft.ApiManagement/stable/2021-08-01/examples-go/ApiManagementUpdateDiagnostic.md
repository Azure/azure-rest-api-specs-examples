Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapimanagement%2Farmapimanagement%2Fv0.5.0/sdk/resourcemanager/apimanagement/armapimanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateDiagnostic.json
func ExampleDiagnosticClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewDiagnosticClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<diagnostic-id>",
		"<if-match>",
		armapimanagement.DiagnosticContract{
			Properties: &armapimanagement.DiagnosticContractProperties{
				AlwaysLog: to.Ptr(armapimanagement.AlwaysLogAllErrors),
				Backend: &armapimanagement.PipelineDiagnosticSettings{
					Response: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
					Request: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
				},
				Frontend: &armapimanagement.PipelineDiagnosticSettings{
					Response: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
					Request: &armapimanagement.HTTPMessageDiagnostic{
						Body: &armapimanagement.BodyDiagnosticSettings{
							Bytes: to.Ptr[int32](512),
						},
						Headers: []*string{
							to.Ptr("Content-type")},
					},
				},
				LoggerID: to.Ptr("<logger-id>"),
				Sampling: &armapimanagement.SamplingSettings{
					Percentage:   to.Ptr[float64](50),
					SamplingType: to.Ptr(armapimanagement.SamplingTypeFixed),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
