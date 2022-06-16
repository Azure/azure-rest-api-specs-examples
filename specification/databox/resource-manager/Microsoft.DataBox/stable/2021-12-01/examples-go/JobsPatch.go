package armdatabox_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databox/resource-manager/Microsoft.DataBox/stable/2021-12-01/examples/JobsPatch.json
func ExampleJobsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<job-name>",
		armdatabox.JobResourceUpdateParameter{
			Properties: &armdatabox.UpdateJobProperties{
				Details: &armdatabox.UpdateJobDetails{
					ContactDetails: &armdatabox.ContactDetails{
						ContactName: to.Ptr("<contact-name>"),
						EmailList: []*string{
							to.Ptr("testing@microsoft.com")},
						Phone:          to.Ptr("<phone>"),
						PhoneExtension: to.Ptr("<phone-extension>"),
					},
					ShippingAddress: &armdatabox.ShippingAddress{
						AddressType:     to.Ptr(armdatabox.AddressTypeCommercial),
						City:            to.Ptr("<city>"),
						CompanyName:     to.Ptr("<company-name>"),
						Country:         to.Ptr("<country>"),
						PostalCode:      to.Ptr("<postal-code>"),
						StateOrProvince: to.Ptr("<state-or-province>"),
						StreetAddress1:  to.Ptr("<street-address1>"),
						StreetAddress2:  to.Ptr("<street-address2>"),
					},
				},
			},
		},
		&armdatabox.JobsClientBeginUpdateOptions{IfMatch: nil,
			ResumeToken: "",
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
