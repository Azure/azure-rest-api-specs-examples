package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ImagesGetUploadUrlForData.json
func ExampleImagesClient_GetUploadURLForData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImagesClient().GetUploadURLForData(ctx, "TestHubRG", "sdkTestHub", armcustomerinsights.GetImageUploadURLInput{
		EntityType:     to.Ptr("Profile"),
		EntityTypeName: to.Ptr("Contact"),
		RelativePath:   to.Ptr("images/profile1.png"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ImageDefinition = armcustomerinsights.ImageDefinition{
	// 	ContentURL: to.Ptr("https://ucidfdbl2001img.blob.core.windows.net/images/3ac3a97a5e3246ffb41812f60fd9f83c/EntityTypeImage/Profile/Contact/images/profile1.png?sv=2015-04-05&sr=b&sig=yIt7DGVRTyNl15%2BPc0kO%2FDITJ2cExnBPVvvh6p86qdc%3D&se=2017-01-06T01%3A56%3A43Z&sp=cw"),
	// 	ImageExists: to.Ptr(false),
	// 	RelativePath: to.Ptr("images/profile1.png"),
	// }
}
