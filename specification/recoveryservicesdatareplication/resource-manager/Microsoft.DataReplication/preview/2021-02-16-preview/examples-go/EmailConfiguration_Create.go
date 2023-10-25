package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/EmailConfiguration_Create.json
func ExampleEmailConfigurationClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEmailConfigurationClient().Create(ctx, "rgrecoveryservicesdatareplication", "4", "0", armrecoveryservicesdatareplication.EmailConfigurationModel{
		Properties: &armrecoveryservicesdatareplication.EmailConfigurationModelProperties{
			CustomEmailAddresses: []*string{
				to.Ptr("ketvbducyailcny")},
			Locale:       to.Ptr("vpnjxjvdqtebnucyxiyrjiko"),
			SendToOwners: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EmailConfigurationModel = armrecoveryservicesdatareplication.EmailConfigurationModel{
	// 	Name: to.Ptr("ywjplnjzaeu"),
	// 	Type: to.Ptr("bkaq"),
	// 	ID: to.Ptr("bvbfy"),
	// 	Properties: &armrecoveryservicesdatareplication.EmailConfigurationModelProperties{
	// 		CustomEmailAddresses: []*string{
	// 			to.Ptr("ketvbducyailcny")},
	// 			Locale: to.Ptr("vpnjxjvdqtebnucyxiyrjiko"),
	// 			SendToOwners: to.Ptr(true),
	// 		},
	// 		SystemData: &armrecoveryservicesdatareplication.EmailConfigurationModelSystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.022Z"); return t}()),
	// 			CreatedBy: to.Ptr("ewufpudzcjrljhmmzhfnxoqdqwnya"),
	// 			CreatedByType: to.Ptr("zioqm"),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:53.022Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("rx"),
	// 			LastModifiedByType: to.Ptr("tqbvuqoakaaqij"),
	// 		},
	// 	}
}
