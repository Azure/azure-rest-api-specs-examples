package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getMetricAlertStatusByName.json
func ExampleMetricAlertsStatusClient_ListByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetricAlertsStatusClient().ListByName(ctx, "EastUs", "custom1", "cmVzb3VyY2VJZD0vc3Vic2NyaXB0aW9ucy8xNGRkZjBjNS03N2M1LTRiNTMtODRmNi1lMWZhNDNhZDY4ZjcvcmVzb3VyY2VHcm91cHMvZ2lndGVzdC9wcm92aWRlcnMvTWljcm9zb2Z0LkNvbXB1dGUvdmlydHVhbE1hY2hpbmVzL2dpZ3dhZG1l", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetricAlertStatusCollection = armmonitor.MetricAlertStatusCollection{
	// 	Value: []*armmonitor.MetricAlertStatus{
	// 		{
	// 			Name: to.Ptr("cmVzb3VyY2VJZD0vc3Vic2NyaXB0aW9ucy8xNGRkZjBjNS03N2M1LTRiNTMtODRmNi1lMWZhNDNhZDY4ZjcvcmVzb3VyY2VHcm91cHMvZ2lndGVzdC9wcm92aWRlcnMvTWljcm9zb2Z0LkNvbXB1dGUvdmlydHVhbE1hY2hpbmVzL2dpZ3dhZG1l"),
	// 			Type: to.Ptr("Microsoft.Insights/metricAlerts/status"),
	// 			ID: to.Ptr("/subscriptions/009f6022-67ec-423e-9aa7-691182870588/resourceGroups/EastUs/providers/microsoft.insights/metricAlerts/custom1"),
	// 			Properties: &armmonitor.MetricAlertStatusProperties{
	// 				Dimensions: map[string]*string{
	// 					"resourceId": to.Ptr("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme"),
	// 				},
	// 				Status: to.Ptr("Healthy"),
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-17T20:17:25.629Z"); return t}()),
	// 			},
	// 	}},
	// }
}
