package armcertificateregistration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/certificateregistration/armcertificateregistration"
)

// Generated from example definition: 2024-11-01/RetrieveSiteSeal.json
func ExampleAppServiceCertificateOrdersClient_RetrieveSiteSeal() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcertificateregistration.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppServiceCertificateOrdersClient().RetrieveSiteSeal(ctx, "testrg123", "SampleCertOrder", armcertificateregistration.SiteSealRequest{
		LightTheme: to.Ptr(true),
		Locale:     to.Ptr("en-us"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcertificateregistration.AppServiceCertificateOrdersClientRetrieveSiteSealResponse{
	// 	SiteSeal: &armcertificateregistration.SiteSeal{
	// 		HTML: to.Ptr("<html>SiteSeal</html>"),
	// 	},
	// }
}
