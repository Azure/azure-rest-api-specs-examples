package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdatePortalRevision.json
func ExamplePortalRevisionClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPortalRevisionClient().BeginUpdate(ctx, "rg1", "apimService1", "20201112101010", "*", armapimanagement.PortalRevisionContract{
		Properties: &armapimanagement.PortalRevisionContractProperties{
			Description: to.Ptr("portal revision update"),
			IsCurrent:   to.Ptr(true),
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
	// res.PortalRevisionContract = armapimanagement.PortalRevisionContract{
	// 	Name: to.Ptr("20201112101010"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/portalRevisions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/namedValues/testprop2"),
	// 	Properties: &armapimanagement.PortalRevisionContractProperties{
	// 		Description: to.Ptr("portal revision update"),
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-13T22:47:13.397Z"); return t}()),
	// 		IsCurrent: to.Ptr(true),
	// 		Status: to.Ptr(armapimanagement.PortalRevisionStatusCompleted),
	// 		UpdatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-13T23:29:25.34Z"); return t}()),
	// 	},
	// }
}
