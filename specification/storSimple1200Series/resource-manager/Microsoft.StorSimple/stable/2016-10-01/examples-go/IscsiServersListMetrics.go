package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersListMetrics.json
func ExampleIscsiServersClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewIscsiServersClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListMetricsPager("HSDK-UGU4PITWNI",
		"HSDK-UGU4PITWNI",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		&armstorsimple1200series.IscsiServersClientListMetricsOptions{Filter: to.Ptr("startTime%20ge%20'2018-08-10T18:30:00Z'%20and%20endTime%20le%20'2018-08-11T18:30:00Z'")})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
