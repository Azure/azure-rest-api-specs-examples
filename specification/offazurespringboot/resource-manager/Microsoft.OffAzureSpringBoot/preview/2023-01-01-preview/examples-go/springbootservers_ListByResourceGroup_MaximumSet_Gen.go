package armspringappdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/springappdiscovery/armspringappdiscovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootservers_ListByResourceGroup_MaximumSet_Gen.json
func ExampleSpringbootserversClient_NewListByResourceGroupPager_springbootserversListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armspringappdiscovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSpringbootserversClient().NewListByResourceGroupPager("rgspringbootservers", "hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SpringbootserversListResult = armspringappdiscovery.SpringbootserversListResult{
		// 	Value: []*armspringappdiscovery.SpringbootserversModel{
		// 		{
		// 			Name: to.Ptr("licrobomiqalwnrti"),
		// 			Type: to.Ptr("wlmmljjoqbqnhbotk"),
		// 			ID: to.Ptr("xsfjeazznmryzgjsuvymj"),
		// 			SystemData: &armspringappdiscovery.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:00:31.448Z"); return t}()),
		// 				CreatedBy: to.Ptr("wnkpkjmpusyo"),
		// 				CreatedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-02T09:00:31.448Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("pbpaajxtbivcftbnhqtngytbx"),
		// 				LastModifiedByType: to.Ptr(armspringappdiscovery.CreatedByTypeUser),
		// 			},
		// 			Properties: &armspringappdiscovery.SpringbootserversProperties{
		// 				Errors: []*armspringappdiscovery.Error{
		// 				},
		// 				FqdnAndIPAddressList: []*string{
		// 				},
		// 				MachineArmID: to.Ptr("fvfkiapbqsprnbzczdfmuryknrna"),
		// 				Port: to.Ptr[int32](10),
		// 				ProvisioningState: to.Ptr(armspringappdiscovery.ProvisioningStateSucceeded),
		// 				Server: to.Ptr("thhuxocfyqpeluqcgnypi"),
		// 				SpringBootApps: to.Ptr[int32](17),
		// 				TotalApps: to.Ptr[int32](5),
		// 			},
		// 	}},
		// }
	}
}
