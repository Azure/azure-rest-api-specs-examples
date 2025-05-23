package armhybridconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
)

// Generated from example definition: 2024-12-01/GenerateAwsTemplate_Post.json
func ExampleGenerateAwsTemplateClient_Post() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridconnectivity.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGenerateAwsTemplateClient("5ACC4579-DB34-4C2F-8F8C-25061168F342").Post(ctx, armhybridconnectivity.GenerateAwsTemplateRequest{
		ConnectorID: to.Ptr("pnxcfjidglabnwxit"),
		SolutionTypes: []*armhybridconnectivity.SolutionTypeSettings{
			{
				SolutionType:     to.Ptr("hjyownzpfxwiufmd"),
				SolutionSettings: &armhybridconnectivity.SolutionSettings{},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridconnectivity.GenerateAwsTemplateClientPostResponse{
	// 	PostResponse: &armhybridconnectivity.PostResponse{
	// 	},
	// }
}
