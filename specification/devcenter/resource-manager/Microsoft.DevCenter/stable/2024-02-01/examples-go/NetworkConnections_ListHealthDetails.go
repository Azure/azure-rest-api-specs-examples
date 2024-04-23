package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/NetworkConnections_ListHealthDetails.json
func ExampleNetworkConnectionsClient_NewListHealthDetailsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkConnectionsClient().NewListHealthDetailsPager("rg1", "uswest3network", &armdevcenter.NetworkConnectionsClientListHealthDetailsOptions{Top: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.HealthCheckStatusDetailsListResult = armdevcenter.HealthCheckStatusDetailsListResult{
		// 	Value: []*armdevcenter.HealthCheckStatusDetails{
		// 		{
		// 			Name: to.Ptr("latest"),
		// 			Type: to.Ptr("Microsoft.DevCenter/networkconnections/healthchecks"),
		// 			ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/networkconnections/uswest3network/healthchecks/latest"),
		// 			SystemData: &armdevcenter.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
		// 				CreatedBy: to.Ptr("System"),
		// 				CreatedByType: to.Ptr(armdevcenter.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("System"),
		// 				LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeApplication),
		// 			},
		// 			Properties: &armdevcenter.HealthCheckStatusDetailsProperties{
		// 				EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-03T12:43:15.000Z"); return t}()),
		// 				HealthChecks: []*armdevcenter.HealthCheck{
		// 					{
		// 						DisplayName: to.Ptr("Azure AD device sync"),
		// 						EndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-03T12:43:14.000Z"); return t}()),
		// 						StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-03T12:43:15.000Z"); return t}()),
		// 						Status: to.Ptr(armdevcenter.HealthCheckStatusPassed),
		// 				}},
		// 				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-03T12:43:14.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
