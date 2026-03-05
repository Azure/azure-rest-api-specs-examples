package armdomainregistration_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainregistration/armdomainregistration"
)

// Generated from example definition: 2024-11-01/CreateAppServiceDomain.json
func ExampleDomainsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdomainregistration.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDomainsClient().BeginCreateOrUpdate(ctx, "testrg123", "example.com", armdomainregistration.Domain{
		Location: to.Ptr("global"),
		Properties: &armdomainregistration.DomainProperties{
			AuthCode:  to.Ptr("exampleAuthCode"),
			AutoRenew: to.Ptr(true),
			Consent: &armdomainregistration.DomainPurchaseConsent{
				AgreedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-10T19:30:53Z"); return t }()),
				AgreedBy: to.Ptr("192.0.2.1"),
				AgreementKeys: []*string{
					to.Ptr("agreementKey1"),
				},
			},
			ContactAdmin: &armdomainregistration.Contact{
				AddressMailing: &armdomainregistration.Address{
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
			ContactBilling: &armdomainregistration.Contact{
				AddressMailing: &armdomainregistration.Address{
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
			ContactRegistrant: &armdomainregistration.Contact{
				AddressMailing: &armdomainregistration.Address{
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
			ContactTech: &armdomainregistration.Contact{
				AddressMailing: &armdomainregistration.Address{
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
			DNSType: to.Ptr(armdomainregistration.DNSTypeDefaultDomainRegistrarDNS),
			Privacy: to.Ptr(false),
		},
		Tags: map[string]*string{},
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
	// res = armdomainregistration.DomainsClientCreateOrUpdateResponse{
	// 	Domain: &armdomainregistration.Domain{
	// 		Name: to.Ptr("example.com"),
	// 		Type: to.Ptr("Microsoft.DomainRegistration/domains"),
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.DomainRegistration/domains/example.com"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armdomainregistration.DomainProperties{
	// 			AuthCode: to.Ptr("exampleAuthCode"),
	// 			AutoRenew: to.Ptr(true),
	// 			Consent: &armdomainregistration.DomainPurchaseConsent{
	// 				AgreedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-10T19:30:53Z"); return t}()),
	// 				AgreedBy: to.Ptr("192.0.2.1"),
	// 				AgreementKeys: []*string{
	// 					to.Ptr("agreementKey1"),
	// 				},
	// 			},
	// 			ContactAdmin: &armdomainregistration.Contact{
	// 				AddressMailing: &armdomainregistration.Address{
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
	// 			ContactBilling: &armdomainregistration.Contact{
	// 				AddressMailing: &armdomainregistration.Address{
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
	// 			ContactRegistrant: &armdomainregistration.Contact{
	// 				AddressMailing: &armdomainregistration.Address{
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
	// 			ContactTech: &armdomainregistration.Contact{
	// 				AddressMailing: &armdomainregistration.Address{
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
	// 			CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-10T19:30:53Z"); return t}()),
	// 			DNSType: to.Ptr(armdomainregistration.DNSTypeDefaultDomainRegistrarDNS),
	// 			DomainNotRenewableReasons: []*armdomainregistration.ResourceNotRenewableReason{
	// 				to.Ptr(armdomainregistration.ResourceNotRenewableReasonExpirationNotInRenewalTimeRange),
	// 			},
	// 			ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-10T19:30:53Z"); return t}()),
	// 			ManagedHostNames: []*armdomainregistration.HostName{
	// 			},
	// 			NameServers: []*string{
	// 				to.Ptr("ns01.ote.domaincontrol.com"),
	// 				to.Ptr("ns02.ote.domaincontrol.com"),
	// 			},
	// 			Privacy: to.Ptr(false),
	// 			ProvisioningState: to.Ptr(armdomainregistration.ProvisioningStateSucceeded),
	// 			ReadyForDNSRecordManagement: to.Ptr(true),
	// 			RegistrationStatus: to.Ptr(armdomainregistration.DomainStatusActive),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
