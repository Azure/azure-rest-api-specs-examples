package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c1cea38fb7e5cec9afe223a2ed15cbe2fbeecbdb/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/examples/OpenShiftVersions_List.json
func ExampleOpenShiftVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOpenShiftVersionsClient().NewListPager("location", nil)
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
		// page.OpenShiftVersionList = armredhatopenshift.OpenShiftVersionList{
		// 	Value: []*armredhatopenshift.OpenShiftVersion{
		// 		{
		// 			ID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armredhatopenshift.OpenShiftVersionProperties{
		// 				Version: to.Ptr("4.10.20"),
		// 			},
		// 	}},
		// }
	}
}
