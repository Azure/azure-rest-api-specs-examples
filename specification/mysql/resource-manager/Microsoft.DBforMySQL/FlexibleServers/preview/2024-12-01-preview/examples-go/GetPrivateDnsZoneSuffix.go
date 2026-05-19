package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: 2024-12-01-preview/GetPrivateDnsZoneSuffix.json
func ExampleGetPrivateDNSZoneSuffixClient_Execute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGetPrivateDNSZoneSuffixClient().Execute(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmysqlflexibleservers.GetPrivateDNSZoneSuffixClientExecuteResponse{
	// 	GetPrivateDNSZoneSuffixResponse: &armmysqlflexibleservers.GetPrivateDNSZoneSuffixResponse{
	// 		PrivateDNSZoneSuffix: to.Ptr("suffix-example"),
	// 	},
	// }
}
