package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/privateLinkScopeOperationStatuses.json
func ExamplePrivateLinkScopeOperationStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkScopeOperationStatusClient().Get(ctx, "713192d7-503f-477a-9cfe-4efc3ee2bd11", "MyResourceGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatus = armmonitor.OperationStatus{
	// 	Name: to.Ptr("713192d7-503f-477a-9cfe-4efc3ee2bd11"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-01T16:13:13.933Z"); return t}()),
	// 	ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/privateLinkScopeOperationStatuses/713192d7-503f-477a-9cfe-4efc3ee2bd11"),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-01T13:13:13.933Z"); return t}()),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
