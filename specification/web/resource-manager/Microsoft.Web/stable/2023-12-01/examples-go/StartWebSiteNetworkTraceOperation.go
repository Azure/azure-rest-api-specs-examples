package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/StartWebSiteNetworkTraceOperation.json
func ExampleWebAppsClient_BeginStartWebSiteNetworkTraceOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWebAppsClient().BeginStartWebSiteNetworkTraceOperation(ctx, "testrg123", "SampleApp", &armappservice.WebAppsClientBeginStartWebSiteNetworkTraceOperationOptions{DurationInSeconds: to.Ptr[int32](60),
		MaxFrameLength: nil,
		SasURL:         nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkTraceArray = []*armappservice.NetworkTrace{
	// 	{
	// 		Path: to.Ptr("D:\\home\\LogFiles\\networktrace\\10.0.0.1_2018_02_01T00_00_00.zip"),
	// 		Message: to.Ptr("Trace file has been saved as D:\\home\\LogFiles\\networktrace\\10.0.0.1_2018_02_01T00_00_00.zip"),
	// 		Status: to.Ptr("Succeeded"),
	// }}
}
