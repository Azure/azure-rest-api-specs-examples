package armmanagementpartner_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementpartner/armmanagementpartner"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementpartner/resource-manager/Microsoft.ManagementPartner/preview/2018-02-01/examples/GetPartnerDetailsNoPartnerId.json
func ExamplePartnersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementpartner.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPartnersClient().Get(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PartnerResponse = armmanagementpartner.PartnerResponse{
	// 	Name: to.Ptr("123456"),
	// 	Type: to.Ptr("Microsoft.ManagementPartner/partner"),
	// 	Etag: to.Ptr[int32](3),
	// 	ID: to.Ptr("/providers/microsoft.managementpartner/partners"),
	// 	Properties: &armmanagementpartner.PartnerProperties{
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-20T01:23:40.528Z"); return t}()),
	// 		ObjectID: to.Ptr("aa67f786-0552-423e-8849-244ed12bf581"),
	// 		PartnerID: to.Ptr("123456"),
	// 		PartnerName: to.Ptr("Test_jefl"),
	// 		State: to.Ptr(armmanagementpartner.ManagementPartnerStateActive),
	// 		TenantID: to.Ptr("1b1121dd-6900-412a-af73-e8d44f81e1c1"),
	// 		UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-20T01:52:57.912Z"); return t}()),
	// 		Version: to.Ptr[int32](3),
	// 	},
	// }
}
