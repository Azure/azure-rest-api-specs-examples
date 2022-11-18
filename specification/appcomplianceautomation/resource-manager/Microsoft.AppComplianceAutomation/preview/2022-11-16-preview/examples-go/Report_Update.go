package armappcomplianceautomation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Report_Update.json
func ExampleReportClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappcomplianceautomation.NewReportClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "testReportName", armappcomplianceautomation.ReportResourcePatch{
		Properties: &armappcomplianceautomation.ReportProperties{
			OfferGUID: to.Ptr("0000"),
			Resources: []*armappcomplianceautomation.ResourceMetadata{
				{
					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"),
					Tags: map[string]*string{
						"key1": to.Ptr("value1"),
					},
				}},
			TimeZone:    to.Ptr("GMT Standard Time"),
			TriggerTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-04T05:11:56.197Z"); return t }()),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
