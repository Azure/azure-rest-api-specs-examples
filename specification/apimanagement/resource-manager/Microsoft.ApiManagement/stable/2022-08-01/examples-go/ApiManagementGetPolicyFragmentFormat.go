package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetPolicyFragmentFormat.json
func ExamplePolicyFragmentClient_Get_apiManagementGetPolicyFragmentFormat() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPolicyFragmentClient().Get(ctx, "rg1", "apimService1", "policyFragment1", &armapimanagement.PolicyFragmentClientGetOptions{Format: to.Ptr(armapimanagement.PolicyFragmentContentFormatRawxml)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PolicyFragmentContract = armapimanagement.PolicyFragmentContract{
	// 	Name: to.Ptr("policyFragment1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/policyFragments"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/policyFragments/policyFragment1"),
	// 	Properties: &armapimanagement.PolicyFragmentContractProperties{
	// 		Format: to.Ptr(armapimanagement.PolicyFragmentContentFormatRawxml),
	// 		Description: to.Ptr("A policy fragment example"),
	// 		Value: to.Ptr("<fragment><set-header name=\"correlationid\" exists-action=\"skip\">\r\n			<value>@{\n                var guidBinary = new byte[16];\n                Array.Copy(Guid.NewGuid().ToByteArray(), 0, guidBinary, 0, 10);\n                long time = DateTime.Now.Ticks;\n                byte[] bytes = new byte[6];\n                unchecked\n                {\n                       bytes[5] = (byte)(time >> 40);\n                       bytes[4] = (byte)(time >> 32);\n                       bytes[3] = (byte)(time >> 24);\n                       bytes[2] = (byte)(time >> 16);\n                       bytes[1] = (byte)(time >> 8);\n                       bytes[0] = (byte)(time);\n                }\n                Array.Copy(bytes, 0, guidBinary, 10, 6);\n                return new Guid(guidBinary).ToString();\n            }\n            </value>\r\n		</set-header></fragment>"),
	// 	},
	// }
}
