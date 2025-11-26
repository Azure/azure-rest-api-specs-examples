package armappservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-01/examples/CreateAppServiceDomain.json
func ExampleDomainsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDomainsClient().BeginCreateOrUpdate(ctx, "testrg123", "example.com", armappservice.Domain{
		Location: to.Ptr("global"),
		Tags:     map[string]*string{},
		Properties: &armappservice.DomainProperties{
			AuthCode:  to.Ptr("exampleAuthCode"),
			AutoRenew: to.Ptr(true),
			Consent: &armappservice.DomainPurchaseConsent{
				AgreedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-10T19:30:53.000Z"); return t }()),
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
	// res.Domain = armappservice.Domain{
	// 	Name: to.Ptr("example.com"),
	// 	Type: to.Ptr("Microsoft.DomainRegistration/domains"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.DomainRegistration/domains/example.com"),
	// 	Location: to.Ptr("global"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armappservice.DomainProperties{
	// 		AuthCode: to.Ptr("exampleAuthCode"),
	// 		AutoRenew: to.Ptr(true),
	// 		Consent: &armappservice.DomainPurchaseConsent{
	// 			AgreedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-10T19:30:53.000Z"); return t}()),
	// 			AgreedBy: to.Ptr("192.0.2.1"),
	// 			AgreementKeys: []*string{
	// 				to.Ptr("agreementKey1")},
	// 			},
	// 			ContactAdmin: &armappservice.Contact{
	// 				AddressMailing: &armappservice.Address{
	// 					Address1: to.Ptr("3400 State St"),
	// 					City: to.Ptr("Chicago"),
	// 					Country: to.Ptr("United States"),
	// 					PostalCode: to.Ptr("67098"),
	// 					State: to.Ptr("IL"),
	// 				},
	// 				Email: to.Ptr("admin@email.com"),
	// 				Fax: to.Ptr("1-245-534-2242"),
	// 				JobTitle: to.Ptr("Admin"),
	// 				NameFirst: to.Ptr("John"),
	// 				NameLast: to.Ptr("Doe"),
	// 				NameMiddle: to.Ptr(""),
	// 				Organization: to.Ptr("Microsoft Inc."),
	// 				Phone: to.Ptr("1-245-534-2242"),
	// 			},
	// 			ContactBilling: &armappservice.Contact{
	// 				AddressMailing: &armappservice.Address{
	// 					Address1: to.Ptr("3400 State St"),
	// 					City: to.Ptr("Chicago"),
	// 					Country: to.Ptr("United States"),
	// 					PostalCode: to.Ptr("67098"),
	// 					State: to.Ptr("IL"),
	// 				},
	// 				Email: to.Ptr("billing@email.com"),
	// 				Fax: to.Ptr("1-245-534-2242"),
	// 				JobTitle: to.Ptr("Billing"),
	// 				NameFirst: to.Ptr("John"),
	// 				NameLast: to.Ptr("Doe"),
	// 				NameMiddle: to.Ptr(""),
	// 				Organization: to.Ptr("Microsoft Inc."),
	// 				Phone: to.Ptr("1-245-534-2242"),
	// 			},
	// 			ContactRegistrant: &armappservice.Contact{
	// 				AddressMailing: &armappservice.Address{
	// 					Address1: to.Ptr("3400 State St"),
	// 					City: to.Ptr("Chicago"),
	// 					Country: to.Ptr("United States"),
	// 					PostalCode: to.Ptr("67098"),
	// 					State: to.Ptr("IL"),
	// 				},
	// 				Email: to.Ptr("registrant@email.com"),
	// 				Fax: to.Ptr("1-245-534-2242"),
	// 				JobTitle: to.Ptr("Registrant"),
	// 				NameFirst: to.Ptr("John"),
	// 				NameLast: to.Ptr("Doe"),
	// 				NameMiddle: to.Ptr(""),
	// 				Organization: to.Ptr("Microsoft Inc."),
	// 				Phone: to.Ptr("1-245-534-2242"),
	// 			},
	// 			ContactTech: &armappservice.Contact{
	// 				AddressMailing: &armappservice.Address{
	// 					Address1: to.Ptr("3400 State St"),
	// 					City: to.Ptr("Chicago"),
	// 					Country: to.Ptr("United States"),
	// 					PostalCode: to.Ptr("67098"),
	// 					State: to.Ptr("IL"),
	// 				},
	// 				Email: to.Ptr("tech@email.com"),
	// 				Fax: to.Ptr("1-245-534-2242"),
	// 				JobTitle: to.Ptr("Tech"),
	// 				NameFirst: to.Ptr("John"),
	// 				NameLast: to.Ptr("Doe"),
	// 				NameMiddle: to.Ptr(""),
	// 				Organization: to.Ptr("Microsoft Inc."),
	// 				Phone: to.Ptr("1-245-534-2242"),
	// 			},
	// 			CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-10T19:30:53.000Z"); return t}()),
	// 			DNSType: to.Ptr(armappservice.DNSTypeDefaultDomainRegistrarDNS),
	// 			DomainNotRenewableReasons: []*armappservice.ResourceNotRenewableReason{
	// 				to.Ptr(armappservice.ResourceNotRenewableReasonExpirationNotInRenewalTimeRange)},
	// 				ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-10T19:30:53.000Z"); return t}()),
	// 				ManagedHostNames: []*armappservice.HostName{
	// 				},
	// 				NameServers: []*string{
	// 					to.Ptr("ns01.ote.domaincontrol.com"),
	// 					to.Ptr("ns02.ote.domaincontrol.com")},
	// 					Privacy: to.Ptr(false),
	// 					ProvisioningState: to.Ptr(armappservice.ProvisioningStateSucceeded),
	// 					ReadyForDNSRecordManagement: to.Ptr(true),
	// 					RegistrationStatus: to.Ptr(armappservice.DomainStatusActive),
	// 				},
	// 			}
}
