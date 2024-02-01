package armspringappdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/springappdiscovery/armspringappdiscovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootservers_Get_MinimumSet_Gen.json
func ExampleSpringbootserversClient_Get_springbootserversGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSpringbootserversClient().Get(ctx, "rgspringbootservers", "hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj", "zkarbqnwnxeozvjrkpdqmgnwedwgtwcmmyqwaijkn", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SpringbootserversModel = armspringappdiscovery.SpringbootserversModel{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.OffAzureSpringBoot/springbootsites/testsite1/springbootservers/server1"),
	// }
}
