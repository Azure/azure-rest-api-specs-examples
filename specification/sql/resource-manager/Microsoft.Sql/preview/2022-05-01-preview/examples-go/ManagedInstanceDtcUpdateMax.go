package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/ManagedInstanceDtcUpdateMax.json
func ExampleManagedInstanceDtcsClient_BeginCreateOrUpdate_updatesManagedInstanceDtcSettingsWithAllOptionalParametersSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstanceDtcsClient().BeginCreateOrUpdate(ctx, "testrg", "testinstance", armsql.DtcNameCurrent, armsql.ManagedInstanceDtc{
		Properties: &armsql.ManagedInstanceDtcProperties{
			DtcEnabled: to.Ptr(true),
			ExternalDNSSuffixSearchList: []*string{
				to.Ptr("dns.example1.com"),
				to.Ptr("dns.example2.com")},
			SecuritySettings: &armsql.ManagedInstanceDtcSecuritySettings{
				SnaLu6Point2TransactionsEnabled: to.Ptr(false),
				TransactionManagerCommunicationSettings: &armsql.ManagedInstanceDtcTransactionManagerCommunicationSettings{
					AllowInboundEnabled:  to.Ptr(false),
					AllowOutboundEnabled: to.Ptr(true),
					Authentication:       to.Ptr("NoAuth"),
				},
				XaTransactionsDefaultTimeout: to.Ptr[int32](1000),
				XaTransactionsEnabled:        to.Ptr(false),
				XaTransactionsMaximumTimeout: to.Ptr[int32](3000),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedInstanceDtc = armsql.ManagedInstanceDtc{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/dtc"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance/dtc/current"),
	// 	Properties: &armsql.ManagedInstanceDtcProperties{
	// 		DtcEnabled: to.Ptr(true),
	// 		DtcHostNameDNSSuffix: to.Ptr("dtcHostNameSuffixExample.com"),
	// 		ExternalDNSSuffixSearchList: []*string{
	// 			to.Ptr("dns.example1.com"),
	// 			to.Ptr("dns.example2.com")},
	// 			ProvisioningState: to.Ptr(armsql.ProvisioningStateSucceeded),
	// 			SecuritySettings: &armsql.ManagedInstanceDtcSecuritySettings{
	// 				SnaLu6Point2TransactionsEnabled: to.Ptr(false),
	// 				TransactionManagerCommunicationSettings: &armsql.ManagedInstanceDtcTransactionManagerCommunicationSettings{
	// 					AllowInboundEnabled: to.Ptr(false),
	// 					AllowOutboundEnabled: to.Ptr(true),
	// 					Authentication: to.Ptr("NoAuth"),
	// 				},
	// 				XaTransactionsDefaultTimeout: to.Ptr[int32](1000),
	// 				XaTransactionsEnabled: to.Ptr(false),
	// 				XaTransactionsMaximumTimeout: to.Ptr[int32](3000),
	// 			},
	// 		},
	// 	}
}
