package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateWorkspaceApiDiagnostic.json
func ExampleWorkspaceAPIDiagnosticClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceAPIDiagnosticClient().Update(ctx, "rg1", "apimService1", "wks1", "echo-api", "applicationinsights", "*", armapimanagement.DiagnosticUpdateContract{
		Properties: &armapimanagement.DiagnosticContractUpdateProperties{
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
			LoggerID: to.Ptr("/workspaces/wks1/loggers/applicationinsights"),
			Sampling: &armapimanagement.SamplingSettings{
				Percentage:   to.Ptr[float64](50),
				SamplingType: to.Ptr(armapimanagement.SamplingTypeFixed),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiagnosticContract = armapimanagement.DiagnosticContract{
	// 	Name: to.Ptr("applicationinsights"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis/diagnostics"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/echo-api/diagnostics/applicationinsights"),
	// 	Properties: &armapimanagement.DiagnosticContractProperties{
	// 		AlwaysLog: to.Ptr(armapimanagement.AlwaysLogAllErrors),
	// 		Backend: &armapimanagement.PipelineDiagnosticSettings{
	// 			Response: &armapimanagement.HTTPMessageDiagnostic{
	// 				Body: &armapimanagement.BodyDiagnosticSettings{
	// 					Bytes: to.Ptr[int32](100),
	// 				},
	// 				Headers: []*string{
	// 				},
	// 			},
	// 			Request: &armapimanagement.HTTPMessageDiagnostic{
	// 				Body: &armapimanagement.BodyDiagnosticSettings{
	// 					Bytes: to.Ptr[int32](100),
	// 				},
	// 				Headers: []*string{
	// 				},
	// 			},
	// 		},
	// 		Frontend: &armapimanagement.PipelineDiagnosticSettings{
	// 			Response: &armapimanagement.HTTPMessageDiagnostic{
	// 				Body: &armapimanagement.BodyDiagnosticSettings{
	// 					Bytes: to.Ptr[int32](100),
	// 				},
	// 				Headers: []*string{
	// 				},
	// 			},
	// 			Request: &armapimanagement.HTTPMessageDiagnostic{
	// 				Body: &armapimanagement.BodyDiagnosticSettings{
	// 					Bytes: to.Ptr[int32](100),
	// 				},
	// 				Headers: []*string{
	// 				},
	// 			},
	// 		},
	// 		HTTPCorrelationProtocol: to.Ptr(armapimanagement.HTTPCorrelationProtocolLegacy),
	// 		LogClientIP: to.Ptr(true),
	// 		LoggerID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/loggers/aisamplingtest"),
	// 		Sampling: &armapimanagement.SamplingSettings{
	// 			Percentage: to.Ptr[float64](100),
	// 			SamplingType: to.Ptr(armapimanagement.SamplingTypeFixed),
	// 		},
	// 	},
	// }
}
