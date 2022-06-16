package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/createOrUpdateLogProfile.json
func ExampleLogProfilesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewLogProfilesClient("df602c9c-7aa0-407d-a6fb-eb20c8bd1192", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"Rac46PostSwapRG",
		armmonitor.LogProfileResource{
			Location: to.Ptr(""),
			Tags:     map[string]*string{},
			Properties: &armmonitor.LogProfileProperties{
				Categories: []*string{
					to.Ptr("Write"),
					to.Ptr("Delete"),
					to.Ptr("Action")},
				Locations: []*string{
					to.Ptr("global")},
				RetentionPolicy: &armmonitor.RetentionPolicy{
					Days:    to.Ptr[int32](3),
					Enabled: to.Ptr(true),
				},
				ServiceBusRuleID: to.Ptr(""),
				StorageAccountID: to.Ptr("/subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/JohnKemTest/providers/Microsoft.Storage/storageAccounts/johnkemtest8162"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
