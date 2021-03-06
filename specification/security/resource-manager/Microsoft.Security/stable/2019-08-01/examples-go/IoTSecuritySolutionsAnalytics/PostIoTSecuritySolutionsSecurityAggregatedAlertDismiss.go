package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/PostIoTSecuritySolutionsSecurityAggregatedAlertDismiss.json
func ExampleIotSecuritySolutionsAnalyticsAggregatedAlertClient_Dismiss() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionsAnalyticsAggregatedAlertClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Dismiss(ctx,
		"IoTEdgeResources",
		"default",
		"IoT_Bruteforce_Fail/2019-02-02/dismiss",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
