package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/examples/ResourceGuardCRUD/ListBackupSecurityPINRequests.json
func ExampleResourceGuardsClient_NewGetBackupSecurityPINRequestsObjectsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceGuardsClient().NewGetBackupSecurityPINRequestsObjectsPager("SampleResourceGroup", "swaggerExample", nil)
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
		// page.DppBaseResourceList = armdataprotection.DppBaseResourceList{
		// 	Value: []*armdataprotection.DppBaseResource{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.DataProtection/resourceGuards/getBackupSecurityPINRequests"),
		// 			ID: to.Ptr("subscriotions/0b352192-dcac-4cc7-992e-a96190ccc68c/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/resourceGuards/swaggerExample/getBackupSecurityPINRequests/default"),
		// 	}},
		// }
	}
}
