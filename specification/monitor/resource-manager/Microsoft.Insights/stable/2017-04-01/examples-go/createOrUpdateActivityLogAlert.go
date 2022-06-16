package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2017-04-01/examples/createOrUpdateActivityLogAlert.json
func ExampleActivityLogAlertsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewActivityLogAlertsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<activity-log-alert-name>",
		armmonitor.ActivityLogAlertResource{
			Resource: armmonitor.Resource{
				Location: to.StringPtr("<location>"),
				Tags:     map[string]*string{},
			},
			Properties: &armmonitor.ActivityLogAlert{
				Description: to.StringPtr("<description>"),
				Actions: &armmonitor.ActivityLogAlertActionList{
					ActionGroups: []*armmonitor.ActivityLogAlertActionGroup{
						{
							ActionGroupID: to.StringPtr("<action-group-id>"),
							WebhookProperties: map[string]*string{
								"sampleWebhookProperty": to.StringPtr("samplePropertyValue"),
							},
						}},
				},
				Condition: &armmonitor.ActivityLogAlertAllOfCondition{
					AllOf: []*armmonitor.ActivityLogAlertLeafCondition{
						{
							Equals: to.StringPtr("<equals>"),
							Field:  to.StringPtr("<field>"),
						},
						{
							Equals: to.StringPtr("<equals>"),
							Field:  to.StringPtr("<field>"),
						}},
				},
				Enabled: to.BoolPtr(true),
				Scopes: []*string{
					to.StringPtr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ActivityLogAlertResource.ID: %s\n", *res.ID)
}
