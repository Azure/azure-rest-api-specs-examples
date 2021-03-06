package armdatabox_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsPatch.json
func ExampleJobsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<job-name>",
		armdatabox.JobResourceUpdateParameter{
			Properties: &armdatabox.UpdateJobProperties{
				Details: &armdatabox.UpdateJobDetails{
					ContactDetails: &armdatabox.ContactDetails{
						ContactName: to.StringPtr("<contact-name>"),
						EmailList: []*string{
							to.StringPtr("testing@microsoft.com")},
						Phone:          to.StringPtr("<phone>"),
						PhoneExtension: to.StringPtr("<phone-extension>"),
					},
					ShippingAddress: &armdatabox.ShippingAddress{
						AddressType:     armdatabox.AddressTypeCommercial.ToPtr(),
						City:            to.StringPtr("<city>"),
						CompanyName:     to.StringPtr("<company-name>"),
						Country:         to.StringPtr("<country>"),
						PostalCode:      to.StringPtr("<postal-code>"),
						StateOrProvince: to.StringPtr("<state-or-province>"),
						StreetAddress1:  to.StringPtr("<street-address1>"),
						StreetAddress2:  to.StringPtr("<street-address2>"),
					},
				},
			},
		},
		&armdatabox.JobsClientBeginUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.JobsClientUpdateResult)
}
