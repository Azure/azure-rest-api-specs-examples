package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ApplicationActionFetchHealth_example.json
func ExampleApplicationsClient_BeginFetchHealth() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginFetchHealth(ctx, "resRg", "myCluster", "myApp", armservicefabricmanagedclusters.ApplicationFetchHealthRequest{
		EventsHealthStateFilter:               to.Ptr(armservicefabricmanagedclusters.HealthFilterError),
		DeployedApplicationsHealthStateFilter: to.Ptr(armservicefabricmanagedclusters.HealthFilterError),
		ServicesHealthStateFilter:             to.Ptr(armservicefabricmanagedclusters.HealthFilterError),
		ExcludeHealthStatistics:               to.Ptr(true),
		Timeout:                               to.Ptr[int64](30),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
