package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: 2025-03-01-preview/AppliancesListClusterUserCredential.json
func ExampleAppliancesClient_ListClusterUserCredential() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppliancesClient().ListClusterUserCredential(ctx, "testresourcegroup", "appliance01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armresourceconnector.AppliancesClientListClusterUserCredentialResponse{
	// 	ApplianceListCredentialResults: &armresourceconnector.ApplianceListCredentialResults{
	// 		HybridConnectionConfig: &armresourceconnector.HybridConnectionConfig{
	// 			ExpirationTime: to.Ptr[int64](123456789),
	// 			HybridConnectionName: to.Ptr("provider/type/bc36ffcf318d5bedfc05ba8b0628dba"),
	// 			Token: to.Ptr("mockSecretOtherprovider/type/bc36ffcf318d5bedfc05ba91c157ReceiverToken"),
	// 			Relay: to.Ptr("relayName"),
	// 		},
	// 		Kubeconfigs: []*armresourceconnector.ApplianceCredentialKubeconfig{
	// 			{
	// 				Name: to.Ptr(armresourceconnector.AccessProfileType("kubeconfigName1")),
	// 				Value: to.Ptr("xxxxxxxxxxxxx"),
	// 			},
	// 		},
	// 	},
	// }
}
