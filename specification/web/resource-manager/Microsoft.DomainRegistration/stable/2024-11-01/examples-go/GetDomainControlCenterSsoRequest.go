package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.DomainRegistration/stable/2024-11-01/examples/GetDomainControlCenterSsoRequest.json
func ExampleDomainsClient_GetControlCenterSsoRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDomainsClient().GetControlCenterSsoRequest(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DomainControlCenterSsoRequest = armappservice.DomainControlCenterSsoRequest{
	// 	PostParameterKey: to.Ptr("saml"),
	// 	PostParameterValue: to.Ptr("1242sds|k1340"),
	// 	URL: to.Ptr("https://Reseller.provider.com/?xRequestId=sdsdsds&app=dcc"),
	// }
}
