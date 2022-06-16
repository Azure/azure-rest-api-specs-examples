package armservicelinker_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker"
)

// x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2021-11-01-preview/examples/PutLink.json
func ExampleLinkerClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicelinker.NewLinkerClient(cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-uri>",
		"<linker-name>",
		armservicelinker.LinkerResource{
			Properties: &armservicelinker.LinkerProperties{
				AuthInfo: &armservicelinker.SecretAuthInfo{
					AuthType: armservicelinker.AuthType("secret").ToPtr(),
					Name:     to.StringPtr("<name>"),
					Secret:   to.StringPtr("<secret>"),
				},
				TargetID: to.StringPtr("<target-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LinkerClientCreateOrUpdateResult)
}
