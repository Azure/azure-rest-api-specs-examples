package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateDiagnostic.json
func ExampleDiagnosticClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapimanagement.NewDiagnosticClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"apimService1",
		"applicationinsights",
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
				LoggerID: to.Ptr("/loggers/azuremonitor"),
				Sampling: &armapimanagement.SamplingSettings{
					Percentage:   to.Ptr[float64](50),
					SamplingType: to.Ptr(armapimanagement.SamplingTypeFixed),
				},
			},
		},
		&armapimanagement.DiagnosticClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
