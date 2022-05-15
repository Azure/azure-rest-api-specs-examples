Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvideoanalyzer%2Farmvideoanalyzer%2Fv0.4.2/sdk/resourcemanager/videoanalyzer/armvideoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```go
package armvideoanalyzer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-create.json
func ExamplePipelineTopologiesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-topology-name>",
		armvideoanalyzer.PipelineTopology{
			Kind: to.Ptr(armvideoanalyzer.KindLive),
			Properties: &armvideoanalyzer.PipelineTopologyProperties{
				Description: to.Ptr("<description>"),
				Parameters: []*armvideoanalyzer.ParameterDeclaration{
					{
						Name:        to.Ptr("<name>"),
						Type:        to.Ptr(armvideoanalyzer.ParameterTypeString),
						Description: to.Ptr("<description>"),
						Default:     to.Ptr("<default>"),
					},
					{
						Name:        to.Ptr("<name>"),
						Type:        to.Ptr(armvideoanalyzer.ParameterTypeSecretString),
						Description: to.Ptr("<description>"),
						Default:     to.Ptr("<default>"),
					}},
				Sinks: []armvideoanalyzer.SinkNodeBaseClassification{
					&armvideoanalyzer.VideoSink{
						Name: to.Ptr("<name>"),
						Type: to.Ptr("<type>"),
						Inputs: []*armvideoanalyzer.NodeInput{
							{
								NodeName: to.Ptr("<node-name>"),
							}},
						VideoCreationProperties: &armvideoanalyzer.VideoCreationProperties{
							Description:   to.Ptr("<description>"),
							SegmentLength: to.Ptr("<segment-length>"),
							Title:         to.Ptr("<title>"),
						},
						VideoName: to.Ptr("<video-name>"),
						VideoPublishingOptions: &armvideoanalyzer.VideoPublishingOptions{
							DisableArchive:        to.Ptr("<disable-archive>"),
							DisableRtspPublishing: to.Ptr("<disable-rtsp-publishing>"),
						},
					}},
				Sources: []armvideoanalyzer.SourceNodeBaseClassification{
					&armvideoanalyzer.RtspSource{
						Name: to.Ptr("<name>"),
						Type: to.Ptr("<type>"),
						Endpoint: &armvideoanalyzer.UnsecuredEndpoint{
							Type: to.Ptr("<type>"),
							Credentials: &armvideoanalyzer.UsernamePasswordCredentials{
								Type:     to.Ptr("<type>"),
								Password: to.Ptr("<password>"),
								Username: to.Ptr("<username>"),
							},
							URL: to.Ptr("<url>"),
						},
						Transport: to.Ptr(armvideoanalyzer.RtspTransportHTTP),
					}},
			},
			SKU: &armvideoanalyzer.SKU{
				Name: to.Ptr(armvideoanalyzer.SKUNameLiveS1),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
