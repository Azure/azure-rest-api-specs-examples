package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2021-02-01/examples/ListTopLevelDomainAgreements.json
func ExampleTopLevelDomainsClient_ListAgreements() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewTopLevelDomainsClient("<subscription-id>", cred, nil)
	pager := client.ListAgreements("<name>",
		armappservice.TopLevelDomainAgreementOption{
			ForTransfer:    to.BoolPtr(false),
			IncludePrivacy: to.BoolPtr(true),
		},
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}
