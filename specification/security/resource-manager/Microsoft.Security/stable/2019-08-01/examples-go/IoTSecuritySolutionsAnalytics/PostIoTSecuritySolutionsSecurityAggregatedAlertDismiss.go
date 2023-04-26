package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/PostIoTSecuritySolutionsSecurityAggregatedAlertDismiss.json
func ExampleIotSecuritySolutionsAnalyticsAggregatedAlertClient_Dismiss() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewIotSecuritySolutionsAnalyticsAggregatedAlertClient().Dismiss(ctx, "IoTEdgeResources", "default", "IoT_Bruteforce_Fail/2019-02-02/dismiss", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
