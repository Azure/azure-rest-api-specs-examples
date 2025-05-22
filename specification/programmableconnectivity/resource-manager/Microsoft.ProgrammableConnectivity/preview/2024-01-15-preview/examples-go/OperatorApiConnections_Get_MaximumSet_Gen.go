package armprogrammableconnectivity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/programmableconnectivity/armprogrammableconnectivity"
)

// Generated from example definition: 2024-01-15-preview/OperatorApiConnections_Get_MaximumSet_Gen.json
func ExampleOperatorAPIConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprogrammableconnectivity.NewClientFactory("B976474B-99FA-4C25-A3BD-8B05C3C3D07A", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperatorAPIConnectionsClient().Get(ctx, "rgopenapi", "uetzqjrwqtkwgcirdqy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armprogrammableconnectivity.OperatorAPIConnectionsClientGetResponse{
	// 	OperatorAPIConnection: &armprogrammableconnectivity.OperatorAPIConnection{
	// 		Properties: &armprogrammableconnectivity.OperatorAPIConnectionProperties{
	// 			OperatorAPIPlanID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/providers/Microsoft.ProgrammableConnectivity/operatorApiPlans/livmzrh"),
	// 			SaasProperties: &armprogrammableconnectivity.SaasProperties{
	// 				SaasSubscriptionID: to.Ptr("mgyusmqt"),
	// 				SaasResourceID: to.Ptr("pekejefyvfviabimdrmno"),
	// 			},
	// 			ConfiguredApplication: &armprogrammableconnectivity.ApplicationProperties{
	// 				Name: to.Ptr("idzqqen"),
	// 				ApplicationDescription: to.Ptr("gjlwegnqvffvsc"),
	// 				ApplicationType: to.Ptr("f"),
	// 				LegalName: to.Ptr("ar"),
	// 				OrganizationDescription: to.Ptr("fcueqzlxxr"),
	// 				TaxNumber: to.Ptr("ngzv"),
	// 				PrivacyContactEmailAddress: to.Ptr("l"),
	// 			},
	// 			AppID: to.Ptr("czgrhbvgr"),
	// 			GatewayID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/example-rg/providers/Microsoft.ProgrammableConnectivity/gateways/cdvcixxcdhjqw"),
	// 			AccountType: to.Ptr(armprogrammableconnectivity.AccountTypeAzureManaged),
	// 			OperatorName: to.Ptr("tjisxmigbgtejdnelabhc"),
	// 			CamaraAPIName: to.Ptr("lnktvqbqcdzmkubxwblwvgwifp"),
	// 			ProvisioningState: to.Ptr(armprogrammableconnectivity.ProvisioningStateSucceeded),
	// 			Status: &armprogrammableconnectivity.Status{
	// 				State: to.Ptr("rvez"),
	// 				Reason: to.Ptr("fpteanxqzqixfmymib"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key5536": to.Ptr("bjhvpzsmtalqxmjjbsfdizhg"),
	// 		},
	// 		Location: to.Ptr("dwvzfkjoepbmksygazllqryyinn"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/example-rg/providers/Microsoft.ProgrammableConnectivity/operatorApiConnections/uecwablqeufseigocrwf"),
	// 		Name: to.Ptr("zsilgtpflhroamaglfbywbn"),
	// 		Type: to.Ptr("stxhhdjwawmqtep"),
	// 		SystemData: &armprogrammableconnectivity.SystemData{
	// 			CreatedBy: to.Ptr("kuprrapuolhnvju"),
	// 			CreatedByType: to.Ptr(armprogrammableconnectivity.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T16:41:38.838Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("lsmrhxnvkpmrxncylgqpkr"),
	// 			LastModifiedByType: to.Ptr(armprogrammableconnectivity.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T16:41:38.838Z"); return t}()),
	// 		},
	// 	},
	// }
}
