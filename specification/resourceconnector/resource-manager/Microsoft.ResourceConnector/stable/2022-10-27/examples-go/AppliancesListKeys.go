package armresourceconnector_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourceconnector/armresourceconnector"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListKeys.json
func ExampleAppliancesClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourceconnector.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAppliancesClient().ListKeys(ctx, "testresourcegroup", "appliance01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplianceListKeysResults = armresourceconnector.ApplianceListKeysResults{
	// 	ArtifactProfiles: map[string]*armresourceconnector.ArtifactProfile{
	// 		"LogsArtifactType": &armresourceconnector.ArtifactProfile{
	// 			Endpoint: to.Ptr("https://<storage-account-name>.blob.core.windows.net/<container-name>?<SAS-token>"),
	// 		},
	// 	},
	// 	Kubeconfigs: []*armresourceconnector.ApplianceCredentialKubeconfig{
	// 		{
	// 			Name: to.Ptr(armresourceconnector.AccessProfileType("kubeconfigName1")),
	// 			Value: to.Ptr("xxxxxxxxxxxxx"),
	// 	}},
	// 	SSHKeys: map[string]*armresourceconnector.SSHKey{
	// 		"LogsKey": &armresourceconnector.SSHKey{
	// 			Certificate: to.Ptr("<Generated Certificate>"),
	// 			CreationTimeStamp: to.Ptr[int64](1660946559),
	// 			ExpirationTimeStamp: to.Ptr[int64](1724119358),
	// 			PrivateKey: to.Ptr("<Generated Private Key>"),
	// 		},
	// 		"ManagementCAKey": &armresourceconnector.SSHKey{
	// 			PublicKey: to.Ptr("<Generated Public Key>"),
	// 		},
	// 		"SSHCustomerUser": &armresourceconnector.SSHKey{
	// 			PrivateKey: to.Ptr("xxxxxxxx"),
	// 			PublicKey: to.Ptr("xxxxxxxx"),
	// 		},
	// 		"ScopedAccessKey": &armresourceconnector.SSHKey{
	// 			Certificate: to.Ptr("<Generated Certificate>"),
	// 			CreationTimeStamp: to.Ptr[int64](1660946559),
	// 			ExpirationTimeStamp: to.Ptr[int64](1724119358),
	// 			PrivateKey: to.Ptr("<Generated Private Key>"),
	// 		},
	// 	},
	// }
}
