package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"
)

// x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/CheckQuotaAvailability.json
func ExampleResourceClient_CheckQuotaAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetapp.NewResourceClient("<subscription-id>", cred, nil)
	res, err := client.CheckQuotaAvailability(ctx,
		"<location>",
		armnetapp.QuotaAvailabilityRequest{
			Name:          to.StringPtr("<name>"),
			Type:          armnetapp.CheckQuotaNameResourceTypes("Microsoft.NetApp/netAppAccounts").ToPtr(),
			ResourceGroup: to.StringPtr("<resource-group>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ResourceClientCheckQuotaAvailabilityResult)
}
