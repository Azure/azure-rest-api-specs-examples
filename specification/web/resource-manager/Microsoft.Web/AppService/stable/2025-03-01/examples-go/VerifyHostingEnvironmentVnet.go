package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/VerifyHostingEnvironmentVnet.json
func ExampleWebSiteManagementClient_VerifyHostingEnvironmentVnet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebSiteManagementClient().VerifyHostingEnvironmentVnet(ctx, armappservice.VnetParameters{
		Properties: &armappservice.VnetParametersProperties{
			VnetName:          to.Ptr("vNet123"),
			VnetResourceGroup: to.Ptr("vNet123rg"),
			VnetSubnetName:    to.Ptr("vNet123SubNet"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VnetValidationFailureDetails = armappservice.VnetValidationFailureDetails{
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/verifyHostingEnvironmentVnet"),
	// 	Properties: &armappservice.VnetValidationFailureDetailsProperties{
	// 		Failed: to.Ptr(false),
	// 		FailedTests: []*armappservice.VnetValidationTestFailure{
	// 		},
	// 	},
	// }
}
