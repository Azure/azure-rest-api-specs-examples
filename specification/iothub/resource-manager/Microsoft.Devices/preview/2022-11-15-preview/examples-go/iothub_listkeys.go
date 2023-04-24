package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/iothub_listkeys.json
func ExampleResourceClient_NewListKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceClient().NewListKeysPager("myResourceGroup", "testHub", nil)
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
		// page.SharedAccessSignatureAuthorizationRuleListResult = armiothub.SharedAccessSignatureAuthorizationRuleListResult{
		// 	Value: []*armiothub.SharedAccessSignatureAuthorizationRule{
		// 		{
		// 			KeyName: to.Ptr("iothubowner"),
		// 			PrimaryKey: to.Ptr("<primaryKey>"),
		// 			Rights: to.Ptr(armiothub.AccessRightsRegistryWriteServiceConnectDeviceConnect),
		// 			SecondaryKey: to.Ptr("<secondaryKey>"),
		// 		},
		// 		{
		// 			KeyName: to.Ptr("service"),
		// 			PrimaryKey: to.Ptr("<primaryKey>"),
		// 			Rights: to.Ptr(armiothub.AccessRightsServiceConnect),
		// 			SecondaryKey: to.Ptr("<secondaryKey>"),
		// 		},
		// 		{
		// 			KeyName: to.Ptr("device"),
		// 			PrimaryKey: to.Ptr("<primaryKey>"),
		// 			Rights: to.Ptr(armiothub.AccessRightsDeviceConnect),
		// 			SecondaryKey: to.Ptr("<secondaryKey>"),
		// 		},
		// 		{
		// 			KeyName: to.Ptr("registryRead"),
		// 			PrimaryKey: to.Ptr("<primaryKey>"),
		// 			Rights: to.Ptr(armiothub.AccessRightsRegistryRead),
		// 			SecondaryKey: to.Ptr("<secondaryKey>"),
		// 		},
		// 		{
		// 			KeyName: to.Ptr("registryReadWrite"),
		// 			PrimaryKey: to.Ptr("<primaryKey>"),
		// 			Rights: to.Ptr(armiothub.AccessRightsRegistryWrite),
		// 			SecondaryKey: to.Ptr("<secondaryKey>"),
		// 	}},
		// }
	}
}
