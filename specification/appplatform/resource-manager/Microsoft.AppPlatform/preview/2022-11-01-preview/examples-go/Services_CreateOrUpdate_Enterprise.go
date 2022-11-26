package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/Services_CreateOrUpdate_Enterprise.json
func ExampleServicesClient_BeginCreateOrUpdate_servicesCreateOrUpdateEnterprise() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappplatform.NewServicesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", armappplatform.ServiceResource{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armappplatform.ClusterResourceProperties{
			MarketplaceResource: &armappplatform.MarketplaceResource{
				Plan:      to.Ptr("tanzu-asc-ent-mtr"),
				Product:   to.Ptr("azure-spring-cloud-vmware-tanzu-2"),
				Publisher: to.Ptr("vmware-inc"),
			},
		},
		SKU: &armappplatform.SKU{
			Name: to.Ptr("E0"),
			Tier: to.Ptr("Enterprise"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
