package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7832c9e47b8998a1c994b98345eea24dbc2ac5b8/specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/domains/createOrUpdate.json
func ExampleDomainsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDomainsClient().BeginCreateOrUpdate(ctx, "MyResourceGroup", "MyEmailServiceResource", "mydomain.com", armcommunication.DomainResource{
		Location: to.Ptr("Global"),
		Properties: &armcommunication.DomainProperties{
			DomainManagement: to.Ptr(armcommunication.DomainManagementCustomerManaged),
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
	// 		UserEngagementTracking: to.Ptr(armcommunication.UserEngagementTrackingDisabled),
	// 		VerificationRecords: &armcommunication.DomainPropertiesVerificationRecords{
	// 			Domain: &armcommunication.DNSRecord{
	// 				Name: to.Ptr("recordName"),
	// 				Type: to.Ptr("TXT"),
	// 				TTL: to.Ptr[int32](3600),
	// 				Value: to.Ptr("recordValue"),
	// 			},
	// 		},
	// 		VerificationStates: &armcommunication.DomainPropertiesVerificationStates{
	// 			Domain: &armcommunication.VerificationStatusRecord{
	// 				ErrorCode: to.Ptr(""),
	// 				Status: to.Ptr(armcommunication.VerificationStatusNotStarted),
	// 			},
	// 		},
	// 	},
	// }
}
