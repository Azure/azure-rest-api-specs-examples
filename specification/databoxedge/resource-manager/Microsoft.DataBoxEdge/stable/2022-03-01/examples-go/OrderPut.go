package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/OrderPut.json
func ExampleOrdersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataboxedge.NewOrdersClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testedgedevice",
		"GroupForEdgeAutomation",
		armdataboxedge.Order{
			Properties: &armdataboxedge.OrderProperties{
				ContactInformation: &armdataboxedge.ContactDetails{
					CompanyName:   to.Ptr("Microsoft"),
					ContactPerson: to.Ptr("John Mcclane"),
					EmailList: []*string{
						to.Ptr("john@microsoft.com")},
					Phone: to.Ptr("(800) 426-9400"),
				},
				ShippingAddress: &armdataboxedge.Address{
					AddressLine1: to.Ptr("Microsoft Corporation"),
					AddressLine2: to.Ptr("One Microsoft Way"),
					AddressLine3: to.Ptr("Redmond"),
					City:         to.Ptr("WA"),
					Country:      to.Ptr("USA"),
					PostalCode:   to.Ptr("98052"),
					State:        to.Ptr("WA"),
				},
			},
		},
		nil)
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
