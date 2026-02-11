package armnginx_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v3"
)

// Generated from example definition: 2025-03-01-preview/Configurations_Analysis.json
func ExampleConfigurationsClient_Analysis() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnginx.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationsClient().Analysis(ctx, "myResourceGroup", "myDeployment", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnginx.ConfigurationsClientAnalysisResponse{
	// 	AnalysisResult: &armnginx.AnalysisResult{
	// 		Data: &armnginx.AnalysisResultData{
	// 			Errors: []*armnginx.AnalysisDiagnostic{
	// 				{
	// 					Description: to.Ptr("Directives outside the http context are not allowed"),
	// 					Directive: to.Ptr("worker_processes"),
	// 					File: to.Ptr("/etc/nginx/nginx.conf"),
	// 					ID: to.Ptr("config-analysis-error-1"),
	// 					Line: to.Ptr[float32](2),
	// 					Message: to.Ptr("You are not allowed to set the worker_processes directive"),
	// 					Rule: to.Ptr("nginx-azure-load-balancer-allowed-directives"),
	// 				},
	// 			},
	// 		},
	// 		Status: to.Ptr("FAILED"),
	// 	},
	// }
}
