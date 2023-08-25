package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateDiagnostic.json
func ExampleDiagnosticClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiagnosticClient().CreateOrUpdate(ctx, "rg1", "apimService1", "applicationinsights", armapimanagement.DiagnosticContract{
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
	}, &armapimanagement.DiagnosticClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiagnosticContract = armapimanagement.DiagnosticContract{
	// 	Name: to.Ptr("applicationinsights"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/diagnostics"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/diagnostics/applicationinsights"),
	// 	Properties: &armapimanagement.DiagnosticContractProperties{
	// 		AlwaysLog: to.Ptr(armapimanagement.AlwaysLogAllErrors),
	// 		Backend: &armapimanagement.PipelineDiagnosticSettings{
	// 			Response: &armapimanagement.HTTPMessageDiagnostic{
	// 				Body: &armapimanagement.BodyDiagnosticSettings{
	// 					Bytes: to.Ptr[int32](512),
	// 				},
	// 				Headers: []*string{
	// 					to.Ptr("Content-type")},
	// 				},
	// 				Request: &armapimanagement.HTTPMessageDiagnostic{
	// 					Body: &armapimanagement.BodyDiagnosticSettings{
	// 						Bytes: to.Ptr[int32](512),
	// 					},
	// 					Headers: []*string{
	// 						to.Ptr("Content-type")},
	// 					},
	// 				},
	// 				Frontend: &armapimanagement.PipelineDiagnosticSettings{
	// 					Response: &armapimanagement.HTTPMessageDiagnostic{
	// 						Body: &armapimanagement.BodyDiagnosticSettings{
	// 							Bytes: to.Ptr[int32](512),
	// 						},
	// 						Headers: []*string{
	// 							to.Ptr("Content-type")},
	// 						},
	// 						Request: &armapimanagement.HTTPMessageDiagnostic{
	// 							Body: &armapimanagement.BodyDiagnosticSettings{
	// 								Bytes: to.Ptr[int32](512),
	// 							},
	// 							Headers: []*string{
	// 								to.Ptr("Content-type")},
	// 							},
	// 						},
	// 						LoggerID: to.Ptr("/loggers/applicationinsights"),
	// 						Sampling: &armapimanagement.SamplingSettings{
	// 							Percentage: to.Ptr[float64](50),
	// 							SamplingType: to.Ptr(armapimanagement.SamplingTypeFixed),
	// 						},
	// 					},
	// 				}
}
