package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/ListSkusByOffer.json
func ExampleSKUsClient_NewListByOfferPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListByOfferPager("test-rg", "myCluster", "publisher1", "offer1", &armazurestackhci.SKUsClientListByOfferOptions{Expand: nil})
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
		// page.SKUList = armazurestackhci.SKUList{
		// 	Value: []*armazurestackhci.SKU{
		// 		{
		// 			Name: to.Ptr("sku1"),
		// 			Type: to.Ptr("Microsoft.AzureStackHCI/clusters/publishers/offers/skus"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster/publishers/publisher1/offers/offer1/skus/sku1"),
		// 			Properties: &armazurestackhci.SKUProperties{
		// 				Content: to.Ptr("{\"id\":\"canonical.ubuntuserver1404lts-arm-14.04.201808140\",\"displayName\":\"Ubuntu Server 14.04 LTS\",\"publisherId\":\"Canonical\",\"publisherName\":\"Canonical\",\"type\":\"VirtualMachine\",\"version\":\"14.04.201808140\",\"properties\":{\"description\":\"Ubuntu Server 14.04.5 LTS amd64. Ubuntu Server is the world's most popular Linux for cloud environments. Updates and patches for Ubuntu 14.04 LTS will be available until 2019-04-17.  Ubuntu Server is the perfect virtual machine (VM) platform for all workloads from web applications to NoSQL databases and Hadoop. For more information see <a href='http://partners.ubuntu.com/microsoft' target='_blank'>Ubuntu on Azure</a> and <a href='http://juju.ubuntu.com' target='_blank'>using Juju to deploy your workloads</a>.<p><h3 class='msportalfx-text-header'>Legal Terms</h3></p><p>By clicking the Create button, I acknowledge that I am getting this software from Canonical and that the <a href='http://www.ubuntu.com/project/about-ubuntu/licensing' target='_blank'>legal terms</a> of Canonical apply to it. Microsoft does not provide rights for third-party software. Also see the <a href='http://www.ubuntu.com/aboutus/privacypolicy' target='_blank'>privacy statement</a> from Canonical.</p>\"},\"extendedProperties\":{\"osType\":\"Linux\",\"offer\":\"UbuntuServer\",\"offerVersion\":\"1.0.52\",\"sku\":\"14.04.5-LTS\",\"galleryItemIdentity\":\"Canonical.UbuntuServer1404LTS-ARM.1.0.52\"},\"links\":[{\"name\":[],\"uri\":[]},{\"name\":[],\"uri\":[]},{\"name\":[],\"uri\":[]},{\"name\":[],\"uri\":[]}],\"iconUris\":{\"medium\":\"https://azstmktdfwcu001.blob.core.windows.net/icons/e5da743bb86d4d429320a75bfa5b96b8/Medium.png\",\"wide\":\"https://azstmktdfwcu001.blob.core.windows.net/icons/e5da743bb86d4d429320a75bfa5b96b8/Wide.png\",\"large\":\"https://azstmktdfwcu001.blob.core.windows.net/icons/e5da743bb86d4d429320a75bfa5b96b8/Large.png\",\"small\":\"https://azstmktdfwcu001.blob.core.windows.net/icons/e5da743bb86d4d429320a75bfa5b96b8/Small.png\"},\"payloadLength\":32212288276,\"compatibility\":{\"isCompatible\":true,\"message\":\"None\",\"description\":\"None\",\"issues\":[]}}"),
		// 				ContentVersion: to.Ptr("2018-01-01"),
		// 				OfferID: to.Ptr("offer1"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				PublisherID: to.Ptr("publisher1"),
		// 				SKUMappings: []*armazurestackhci.SKUMappings{
		// 					{
		// 						CatalogPlanID: to.Ptr("microsoftsqlserver.sql2019-ubuntu2004enterprise-arm"),
		// 						MarketplaceSKUID: to.Ptr("enterprise"),
		// 						MarketplaceSKUVersions: []*string{
		// 							to.Ptr("15.0.220208")},
		// 					}},
		// 				},
		// 		}},
		// 	}
	}
}
