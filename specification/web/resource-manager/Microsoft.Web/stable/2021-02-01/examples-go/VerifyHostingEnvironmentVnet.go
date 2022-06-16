package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/VerifyHostingEnvironmentVnet.json
func ExampleWebSiteManagementClient_VerifyHostingEnvironmentVnet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewWebSiteManagementClient("<subscription-id>", cred, nil)
	res, err := client.VerifyHostingEnvironmentVnet(ctx,
		armappservice.VnetParameters{
			Properties: &armappservice.VnetParametersProperties{
				VnetName:          to.StringPtr("<vnet-name>"),
				VnetResourceGroup: to.StringPtr("<vnet-resource-group>"),
				VnetSubnetName:    to.StringPtr("<vnet-subnet-name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("VnetValidationFailureDetails.ID: %s\n", *res.ID)
}
