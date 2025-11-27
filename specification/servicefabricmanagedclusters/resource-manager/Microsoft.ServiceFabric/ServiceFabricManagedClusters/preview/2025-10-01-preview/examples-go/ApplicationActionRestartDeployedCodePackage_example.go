package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ApplicationActionRestartDeployedCodePackage_example.json
func ExampleApplicationsClient_BeginRestartDeployedCodePackage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginRestartDeployedCodePackage(ctx, "resRg", "myCluster", "myApp", armservicefabricmanagedclusters.RestartDeployedCodePackageRequest{
		NodeName:                   to.Ptr("nt1_0"),
		ServiceManifestName:        to.Ptr("TestPkg"),
		ServicePackageActivationID: to.Ptr("sharedProcess"),
		CodePackageName:            to.Ptr("Code"),
		CodePackageInstanceID:      to.Ptr("133991326715515522"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
