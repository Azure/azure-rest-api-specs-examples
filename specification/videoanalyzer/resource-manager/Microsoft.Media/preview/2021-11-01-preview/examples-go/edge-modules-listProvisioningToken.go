package armvideoanalyzer_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/edge-modules-listProvisioningToken.json
func ExampleEdgeModulesClient_ListProvisioningToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewEdgeModulesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListProvisioningToken(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<edge-module-name>",
		armvideoanalyzer.ListProvisioningTokenInput{
			ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-23T11:04:49.0526841-08:00"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
