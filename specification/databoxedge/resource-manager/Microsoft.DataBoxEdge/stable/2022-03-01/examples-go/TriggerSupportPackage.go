package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/TriggerSupportPackage.json
func ExampleSupportPackagesClient_BeginTriggerSupportPackage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSupportPackagesClient().BeginTriggerSupportPackage(ctx, "testedgedevice", "GroupForEdgeAutomation", armdataboxedge.TriggerSupportPackageRequest{
		Properties: &armdataboxedge.SupportPackageRequestProperties{
			Include:          to.Ptr("DefaultWithDumps"),
			MaximumTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-18T02:19:51.427Z"); return t }()),
			MinimumTimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-12-18T02:18:51.427Z"); return t }()),
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
