package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/EmailConfiguration_List.json
func ExampleEmailConfigurationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEmailConfigurationClient().NewListPager("rgrecoveryservicesdatareplication", "4", nil)
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
		// page.EmailConfigurationModelCollection = armrecoveryservicesdatareplication.EmailConfigurationModelCollection{
		// 	Value: []*armrecoveryservicesdatareplication.EmailConfigurationModel{
		// 		{
		// 			Name: to.Ptr("ywjplnjzaeu"),
		// 			Type: to.Ptr("bkaq"),
		// 			ID: to.Ptr("bvbfy"),
		// 			Properties: &armrecoveryservicesdatareplication.EmailConfigurationModelProperties{
		// 				CustomEmailAddresses: []*string{
		// 					to.Ptr("ketvbducyailcny")},
		// 					Locale: to.Ptr("vpnjxjvdqtebnucyxiyrjiko"),
		// 					SendToOwners: to.Ptr(true),
		// 				},
		// 				SystemData: &armrecoveryservicesdatareplication.EmailConfigurationModelSystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.022Z"); return t}()),
		// 					CreatedBy: to.Ptr("ewufpudzcjrljhmmzhfnxoqdqwnya"),
		// 					CreatedByType: to.Ptr("zioqm"),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.022Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("rx"),
		// 					LastModifiedByType: to.Ptr("tqbvuqoakaaqij"),
		// 				},
		// 		}},
		// 	}
	}
}
