package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/SapLandscapeMonitor_Create.json
func ExampleSapLandscapeMonitorClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSapLandscapeMonitorClient().Create(ctx, "myResourceGroup", "mySapMonitor", armworkloads.SapLandscapeMonitor{
		Properties: &armworkloads.SapLandscapeMonitorProperties{
			Grouping: &armworkloads.SapLandscapeMonitorPropertiesGrouping{
				Landscape: []*armworkloads.SapLandscapeMonitorSidMapping{
					{
						Name: to.Ptr("Prod"),
						TopSid: []*string{
							to.Ptr("SID1"),
							to.Ptr("SID2")},
					}},
				SapApplication: []*armworkloads.SapLandscapeMonitorSidMapping{
					{
						Name: to.Ptr("ERP1"),
						TopSid: []*string{
							to.Ptr("SID1"),
							to.Ptr("SID2")},
					}},
			},
			TopMetricsThresholds: []*armworkloads.SapLandscapeMonitorMetricThresholds{
				{
					Name:   to.Ptr("Instance Availability"),
					Green:  to.Ptr[float32](90),
					Red:    to.Ptr[float32](50),
					Yellow: to.Ptr[float32](75),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SapLandscapeMonitor = armworkloads.SapLandscapeMonitor{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Workloads/monitors/sapLandscapeMonitor"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Workloads/monitors/myMonitor/sapLandscapeMonitor/default"),
	// 	SystemData: &armworkloads.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 	},
	// 	Properties: &armworkloads.SapLandscapeMonitorProperties{
	// 		Grouping: &armworkloads.SapLandscapeMonitorPropertiesGrouping{
	// 			Landscape: []*armworkloads.SapLandscapeMonitorSidMapping{
	// 				{
	// 					Name: to.Ptr("Prod"),
	// 					TopSid: []*string{
	// 						to.Ptr("SID1"),
	// 						to.Ptr("SID2")},
	// 				}},
	// 				SapApplication: []*armworkloads.SapLandscapeMonitorSidMapping{
	// 					{
	// 						Name: to.Ptr("ERP1"),
	// 						TopSid: []*string{
	// 							to.Ptr("SID1"),
	// 							to.Ptr("SID2")},
	// 					}},
	// 				},
	// 				TopMetricsThresholds: []*armworkloads.SapLandscapeMonitorMetricThresholds{
	// 					{
	// 						Name: to.Ptr("Instance Availability"),
	// 						Green: to.Ptr[float32](90),
	// 						Red: to.Ptr[float32](50),
	// 						Yellow: to.Ptr[float32](75),
	// 				}},
	// 			},
	// 		}
}
