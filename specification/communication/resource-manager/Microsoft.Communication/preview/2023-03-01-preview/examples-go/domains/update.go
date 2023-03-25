package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ddcde53422ff186f5b69fb32338ecac3d11c3bca/specification/communication/resource-manager/Microsoft.Communication/preview/2023-03-01-preview/examples/domains/update.json
func ExampleDomainsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDomainsClient().BeginUpdate(ctx, "MyResourceGroup", "MyEmailServiceResource", "mydomain.com", armcommunication.UpdateDomainRequestParameters{
		Properties: &armcommunication.UpdateDomainProperties{
			UserEngagementTracking: to.Ptr(armcommunication.UserEngagementTrackingEnabled),
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
	// res.DomainResource = armcommunication.DomainResource{
	// 	Name: to.Ptr("mydomain.com"),
	// 	Type: to.Ptr("Microsoft.Communication/EmailServices/Domains"),
	// 	ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/EmailServices/MyEmailServiceResource/Domains/mydomain.com"),
	// 	Location: to.Ptr("Global"),
	// 	Properties: &armcommunication.DomainProperties{
	// 		DataLocation: to.Ptr("United States"),
	// 		DomainManagement: to.Ptr(armcommunication.DomainManagementCustomerManaged),
	// 		FromSenderDomain: to.Ptr("mydomain.com"),
	// 		MailFromSenderDomain: to.Ptr("mydomain.com"),
	// 		ProvisioningState: to.Ptr(armcommunication.DomainsProvisioningStateSucceeded),
	// 		UserEngagementTracking: to.Ptr(armcommunication.UserEngagementTrackingEnabled),
	// 		VerificationRecords: &armcommunication.DomainPropertiesVerificationRecords{
	// 			Domain: &armcommunication.DNSRecord{
	// 				Name: to.Ptr("recordName"),
	// 				Type: to.Ptr("TXT"),
	// 				TTL: to.Ptr[int32](3600),
	// 				Value: to.Ptr("recordValue"),
	// 			},
	// 			SPF: &armcommunication.DNSRecord{
	// 				Name: to.Ptr("recordName"),
	// 				Type: to.Ptr("TXT"),
	// 				TTL: to.Ptr[int32](3600),
	// 				Value: to.Ptr("recordValue"),
	// 			},
	// 		},
	// 		VerificationStates: &armcommunication.DomainPropertiesVerificationStates{
	// 			Domain: &armcommunication.VerificationStatusRecord{
	// 				ErrorCode: to.Ptr(""),
	// 				Status: to.Ptr(armcommunication.VerificationStatusVerified),
	// 			},
	// 			SPF: &armcommunication.VerificationStatusRecord{
	// 				ErrorCode: to.Ptr(""),
	// 				Status: to.Ptr(armcommunication.VerificationStatusNotStarted),
	// 			},
	// 		},
	// 	},
	// }
}
