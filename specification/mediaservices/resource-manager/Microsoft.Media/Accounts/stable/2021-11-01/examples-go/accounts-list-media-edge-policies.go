package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2021-11-01/examples/accounts-list-media-edge-policies.json
func ExampleClient_ListEdgePolicies() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ListEdgePolicies(ctx, "contoso", "contososports", armmediaservices.ListEdgePoliciesInput{
		DeviceID: to.Ptr("contosiothubhost_contosoiotdevice"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EdgePolicies = armmediaservices.EdgePolicies{
	// 	UsageDataCollectionPolicy: &armmediaservices.EdgeUsageDataCollectionPolicy{
	// 		DataCollectionFrequency: to.Ptr("PT10M"),
	// 		DataReportingFrequency: to.Ptr("PT1H"),
	// 		EventHubDetails: &armmediaservices.EdgeUsageDataEventHub{
	// 			Name: to.Ptr("ams-x"),
	// 			Namespace: to.Ptr("ams-y-1-1"),
	// 			Token: to.Ptr("SharedAccessSignature sr=sb%3a%2f%2fams-x.servicebus.windows.net%2fams-y-1-1%2fPublishers%2f96a510a1-0565-492a-a67f-83d1aed1d1f6_SampleDeviceId&sig=signature_value%3d&se=1584073736&skn=EdgeUsageData"),
	// 		},
	// 		MaxAllowedUnreportedUsageDuration: to.Ptr("PT36H"),
	// 	},
	// }
}
