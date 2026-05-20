package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: 2025-03-01-preview/IacProfile_ScaleTemplate.json
func ExampleIacProfilesClient_Scale() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("a0a37f63-7183-4e86-9ac7-ce8036a3ed31", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIacProfilesClient().Scale(ctx, "resourceGroup1", "iacprofile", armdevhub.ScaleTemplateRequest{
		ScaleRequirement: []*armdevhub.ScaleProperty{
			{
				NumberOfStore: to.Ptr[int32](10),
				Region:        to.Ptr("useast"),
				Stage:         to.Ptr("dev"),
			},
		},
		TemplateName: to.Ptr("base"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevhub.IacProfilesClientScaleResponse{
	// 	PrLinkResponse: armdevhub.PrLinkResponse{
	// 		PrLink: to.Ptr("xxxx"),
	// 	},
	// }
}
