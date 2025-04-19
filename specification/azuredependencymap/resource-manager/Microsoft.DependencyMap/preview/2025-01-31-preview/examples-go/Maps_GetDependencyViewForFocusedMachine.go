package armdependencymap_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dependencymap/armdependencymap"
)

// Generated from example definition: 2025-01-31-preview/Maps_GetDependencyViewForFocusedMachine.json
func ExampleMapsClient_BeginGetDependencyViewForFocusedMachine() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdependencymap.NewClientFactory("D6E58BDB-45F1-41EC-A884-1FC945058848", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMapsClient().BeginGetDependencyViewForFocusedMachine(ctx, "rgdependencyMap", "mapsTest1", armdependencymap.GetDependencyViewForFocusedMachineRequest{
		FocusedMachineID: to.Ptr("imzykeisagngrnfinbqtu"),
		Filters: &armdependencymap.VisualizationFilter{
			DateTime: &armdependencymap.DateTimeFilter{
				StartDateTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-29T07:35:15.336Z"); return t }()),
				EndDateTimeUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-29T07:35:15.336Z"); return t }()),
			},
			ProcessNameFilter: &armdependencymap.ProcessNameFilter{
				Operator: to.Ptr(armdependencymap.ProcessNameFilterOperatorContains),
				ProcessNames: []*string{
					to.Ptr("mnqtvduwzemjcvvmnnoqvcuemwhnz"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
