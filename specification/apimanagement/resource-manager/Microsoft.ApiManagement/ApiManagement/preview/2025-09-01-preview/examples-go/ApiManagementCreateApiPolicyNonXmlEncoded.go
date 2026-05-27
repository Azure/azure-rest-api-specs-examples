package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateApiPolicyNonXmlEncoded.json
func ExampleAPIPolicyClient_CreateOrUpdate_apiManagementCreateApiPolicyNonXmlEncoded() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIPolicyClient().CreateOrUpdate(ctx, "rg1", "apimService1", "5600b57e7e8880006a040001", armapimanagement.PolicyIDNamePolicy, armapimanagement.PolicyContract{
		Properties: &armapimanagement.PolicyContractProperties{
			Format: to.Ptr(armapimanagement.PolicyContentFormatRawxml),
			Value:  to.Ptr("<policies>\r\n     <inbound>\r\n     <base />\r\n  <set-header name=\"newvalue\" exists-action=\"override\">\r\n   <value>\"@(context.Request.Headers.FirstOrDefault(h => h.Ke==\"Via\"))\" </value>\r\n    </set-header>\r\n  </inbound>\r\n      </policies>"),
		},
	}, &armapimanagement.APIPolicyClientCreateOrUpdateOptions{
		IfMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.APIPolicyClientCreateOrUpdateResponse{
	// 	PolicyContract: armapimanagement.PolicyContract{
	// 		Name: to.Ptr("policy"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/apis/operations/policies"),
	// 		ID: to.Ptr("/subscriptions/4c1a3bc6-89f9-46fe-a175-5d8984b25095/resourcegroups/Api-DF-West-US/providers/Microsoft.ApiManagement/service/samirmsiservice2/apis/echo-api/operations/create-resource/policies/policy"),
	// 		Properties: &armapimanagement.PolicyContractProperties{
	// 			Value: to.Ptr("<policies>\r\n  <inbound>\r\n    <base />\r\n    <set-header name=\"newvalue\" exists-action=\"override\">\r\n      <value>\"@(context.Request.Headers.FirstOrDefault(h =&gt; h.Ke==\"Via\"))\" </value>\r\n    </set-header>\r\n  </inbound>\r\n</policies>"),
	// 		},
	// 	},
	// }
}
