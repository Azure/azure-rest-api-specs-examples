package armdeviceprovisioningservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceprovisioningservices/armdeviceprovisioningservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSListKeys.json
func ExampleIotDpsResourceClient_NewListKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceprovisioningservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIotDpsResourceClient().NewListKeysPager("myFirstProvisioningService", "myResourceGroup", nil)
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
		// page.SharedAccessSignatureAuthorizationRuleListResult = armdeviceprovisioningservices.SharedAccessSignatureAuthorizationRuleListResult{
		// 	Value: []*armdeviceprovisioningservices.SharedAccessSignatureAuthorizationRuleAccessRightsDescription{
		// 		{
		// 			KeyName: to.Ptr("key1"),
		// 			PrimaryKey: to.Ptr("#####################################"),
		// 			Rights: to.Ptr(armdeviceprovisioningservices.AccessRightsDescriptionServiceConfig),
		// 			SecondaryKey: to.Ptr("###################################"),
		// 		},
		// 		{
		// 			KeyName: to.Ptr("key2"),
		// 			PrimaryKey: to.Ptr("#######################################"),
		// 			Rights: to.Ptr(armdeviceprovisioningservices.AccessRightsDescriptionServiceConfig),
		// 			SecondaryKey: to.Ptr("####################################="),
		// 	}},
		// }
	}
}
