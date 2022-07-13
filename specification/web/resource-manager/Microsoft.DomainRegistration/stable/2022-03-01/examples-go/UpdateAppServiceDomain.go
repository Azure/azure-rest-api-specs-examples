package armappservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/UpdateAppServiceDomain.json
func ExampleDomainsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappservice.NewDomainsClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"testrg123",
		"example.com",
		armappservice.DomainPatchResource{
			Properties: &armappservice.DomainPatchResourceProperties{
				AuthCode:  to.Ptr("exampleAuthCode"),
				AutoRenew: to.Ptr(true),
				Consent: &armappservice.DomainPurchaseConsent{
					AgreedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-10T19:30:53Z"); return t }()),
					AgreedBy: to.Ptr("192.0.2.1"),
					AgreementKeys: []*string{
						to.Ptr("agreementKey1")},
				},
				ContactAdmin: &armappservice.Contact{
					AddressMailing: &armappservice.Address{
						Address1:   to.Ptr("3400 State St"),
						City:       to.Ptr("Chicago"),
						Country:    to.Ptr("United States"),
						PostalCode: to.Ptr("67098"),
						State:      to.Ptr("IL"),
					},
					Email:        to.Ptr("admin@email.com"),
					Fax:          to.Ptr("1-245-534-2242"),
					JobTitle:     to.Ptr("Admin"),
					NameFirst:    to.Ptr("John"),
					NameLast:     to.Ptr("Doe"),
					NameMiddle:   to.Ptr(""),
					Organization: to.Ptr("Microsoft Inc."),
					Phone:        to.Ptr("1-245-534-2242"),
				},
				ContactBilling: &armappservice.Contact{
					AddressMailing: &armappservice.Address{
						Address1:   to.Ptr("3400 State St"),
						City:       to.Ptr("Chicago"),
						Country:    to.Ptr("United States"),
						PostalCode: to.Ptr("67098"),
						State:      to.Ptr("IL"),
					},
					Email:        to.Ptr("billing@email.com"),
					Fax:          to.Ptr("1-245-534-2242"),
					JobTitle:     to.Ptr("Billing"),
					NameFirst:    to.Ptr("John"),
					NameLast:     to.Ptr("Doe"),
					NameMiddle:   to.Ptr(""),
					Organization: to.Ptr("Microsoft Inc."),
					Phone:        to.Ptr("1-245-534-2242"),
				},
				ContactRegistrant: &armappservice.Contact{
					AddressMailing: &armappservice.Address{
						Address1:   to.Ptr("3400 State St"),
						City:       to.Ptr("Chicago"),
						Country:    to.Ptr("United States"),
						PostalCode: to.Ptr("67098"),
						State:      to.Ptr("IL"),
					},
					Email:        to.Ptr("registrant@email.com"),
					Fax:          to.Ptr("1-245-534-2242"),
					JobTitle:     to.Ptr("Registrant"),
					NameFirst:    to.Ptr("John"),
					NameLast:     to.Ptr("Doe"),
					NameMiddle:   to.Ptr(""),
					Organization: to.Ptr("Microsoft Inc."),
					Phone:        to.Ptr("1-245-534-2242"),
				},
				ContactTech: &armappservice.Contact{
					AddressMailing: &armappservice.Address{
						Address1:   to.Ptr("3400 State St"),
						City:       to.Ptr("Chicago"),
						Country:    to.Ptr("United States"),
						PostalCode: to.Ptr("67098"),
						State:      to.Ptr("IL"),
					},
					Email:        to.Ptr("tech@email.com"),
					Fax:          to.Ptr("1-245-534-2242"),
					JobTitle:     to.Ptr("Tech"),
					NameFirst:    to.Ptr("John"),
					NameLast:     to.Ptr("Doe"),
					NameMiddle:   to.Ptr(""),
					Organization: to.Ptr("Microsoft Inc."),
					Phone:        to.Ptr("1-245-534-2242"),
				},
				DNSType: to.Ptr(armappservice.DNSTypeDefaultDomainRegistrarDNS),
				Privacy: to.Ptr(false),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
