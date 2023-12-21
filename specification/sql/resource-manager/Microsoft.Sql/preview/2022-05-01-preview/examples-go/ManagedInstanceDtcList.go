package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/ManagedInstanceDtcList.json
func ExampleManagedInstanceDtcsClient_NewListByManagedInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedInstanceDtcsClient().NewListByManagedInstancePager("testrg", "testinstance", nil)
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
		// page.ManagedInstanceDtcListResult = armsql.ManagedInstanceDtcListResult{
		// 	Value: []*armsql.ManagedInstanceDtc{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/dtc"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance/dtc/current"),
		// 			Properties: &armsql.ManagedInstanceDtcProperties{
		// 				DtcEnabled: to.Ptr(true),
		// 				DtcHostNameDNSSuffix: to.Ptr("dtcHostNameSuffixExample.com"),
		// 				ExternalDNSSuffixSearchList: []*string{
		// 					to.Ptr("dns.example1.com"),
		// 					to.Ptr("dns.example2.com")},
		// 					ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
		// 					SecuritySettings: &armsql.ManagedInstanceDtcSecuritySettings{
		// 						SnaLu6Point2TransactionsEnabled: to.Ptr(false),
		// 						TransactionManagerCommunicationSettings: &armsql.ManagedInstanceDtcTransactionManagerCommunicationSettings{
		// 							AllowInboundEnabled: to.Ptr(false),
		// 							AllowOutboundEnabled: to.Ptr(true),
		// 							Authentication: to.Ptr("NoAuth"),
		// 						},
		// 						XaTransactionsDefaultTimeout: to.Ptr[int32](1000),
		// 						XaTransactionsEnabled: to.Ptr(false),
		// 						XaTransactionsMaximumTimeout: to.Ptr[int32](3000),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
