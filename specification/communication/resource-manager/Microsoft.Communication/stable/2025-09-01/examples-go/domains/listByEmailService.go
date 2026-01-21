package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7832c9e47b8998a1c994b98345eea24dbc2ac5b8/specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/domains/listByEmailService.json
func ExampleDomainsClient_NewListByEmailServiceResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDomainsClient().NewListByEmailServiceResourcePager("MyResourceGroup", "MyEmailServiceResource", nil)
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
		// page.DomainResourceList = armcommunication.DomainResourceList{
		// 	Value: []*armcommunication.DomainResource{
		// 		{
		// 			Name: to.Ptr("mydomain.com"),
		// 			Type: to.Ptr("Microsoft.Communication/EmailServices/Domains"),
		// 			ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/EmailServices/MyEmailServiceResource/Domains/mydomain.com"),
		// 			Location: to.Ptr("Global"),
		// 			Properties: &armcommunication.DomainProperties{
		// 				DataLocation: to.Ptr("United States"),
		// 				DomainManagement: to.Ptr(armcommunication.DomainManagementCustomerManaged),
		// 				FromSenderDomain: to.Ptr("mydomain.com"),
		// 				MailFromSenderDomain: to.Ptr("mydomain.com"),
		// 				ProvisioningState: to.Ptr(armcommunication.DomainsProvisioningStateSucceeded),
		// 				VerificationRecords: &armcommunication.DomainPropertiesVerificationRecords{
		// 					Domain: &armcommunication.DNSRecord{
		// 						Name: to.Ptr("recordName"),
		// 						Type: to.Ptr("TXT"),
		// 						TTL: to.Ptr[int32](3600),
		// 						Value: to.Ptr("recordValue"),
		// 					},
		// 					SPF: &armcommunication.DNSRecord{
		// 						Name: to.Ptr("recordName"),
		// 						Type: to.Ptr("TXT"),
		// 						TTL: to.Ptr[int32](3600),
		// 						Value: to.Ptr("recordValue"),
		// 					},
		// 				},
		// 				VerificationStates: &armcommunication.DomainPropertiesVerificationStates{
		// 					Domain: &armcommunication.VerificationStatusRecord{
		// 						ErrorCode: to.Ptr(""),
		// 						Status: to.Ptr(armcommunication.VerificationStatusVerified),
		// 					},
		// 					SPF: &armcommunication.VerificationStatusRecord{
		// 						ErrorCode: to.Ptr(""),
		// 						Status: to.Ptr(armcommunication.VerificationStatusNotStarted),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
