package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsPatch.json
func ExampleJobsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatabox.NewJobsClient("fa68082f-8ff7-4a25-95c7-ce9da541242f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"SdkRg5154",
		"SdkJob952",
		armdatabox.JobResourceUpdateParameter{
			Properties: &armdatabox.UpdateJobProperties{
				Details: &armdatabox.UpdateJobDetails{
					ContactDetails: &armdatabox.ContactDetails{
						ContactName: to.Ptr("Update Job"),
						EmailList: []*string{
							to.Ptr("testing@microsoft.com")},
						Phone:          to.Ptr("1234567890"),
						PhoneExtension: to.Ptr("1234"),
					},
					ShippingAddress: &armdatabox.ShippingAddress{
						AddressType:     to.Ptr(armdatabox.AddressTypeCommercial),
						City:            to.Ptr("San Francisco"),
						CompanyName:     to.Ptr("Microsoft"),
						Country:         to.Ptr("US"),
						PostalCode:      to.Ptr("94107"),
						StateOrProvince: to.Ptr("CA"),
						StreetAddress1:  to.Ptr("16 TOWNSEND ST"),
						StreetAddress2:  to.Ptr("Unit 1"),
					},
				},
			},
		},
		&armdatabox.JobsClientBeginUpdateOptions{IfMatch: nil})
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
