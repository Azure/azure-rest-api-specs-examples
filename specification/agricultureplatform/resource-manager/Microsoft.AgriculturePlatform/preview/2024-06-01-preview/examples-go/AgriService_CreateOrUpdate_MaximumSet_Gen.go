package armagricultureplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agricultureplatform/armagricultureplatform"
)

// Generated from example definition: 2024-06-01-preview/AgriService_CreateOrUpdate_MaximumSet_Gen.json
func ExampleAgriServiceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armagricultureplatform.NewClientFactory("83D293F5-DEFD-4D48-B120-1DC713BE338A", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgriServiceClient().BeginCreateOrUpdate(ctx, "rgopenapi", "abc123", armagricultureplatform.AgriServiceResource{
		Properties: &armagricultureplatform.AgriServiceResourceProperties{
			Config:                         &armagricultureplatform.AgriServiceConfig{},
			ManagedOnBehalfOfConfiguration: &armagricultureplatform.ManagedOnBehalfOfConfiguration{},
			DataConnectorCredentials: []*armagricultureplatform.DataConnectorCredentialMap{
				{
					Key: to.Ptr("BackendAADApplicationCredentials"),
					Value: &armagricultureplatform.DataConnectorCredentials{
						ClientID: to.Ptr("dce298a8-1eec-481a-a8f9-a3cd5a8257b2"),
					},
				},
			},
			InstalledSolutions: []*armagricultureplatform.InstalledSolutionMap{
				{
					Key: to.Ptr("bayerAgPowered.cwum"),
					Value: &armagricultureplatform.Solution{
						ApplicationName: to.Ptr("bayerAgPowered.cwum"),
					},
				},
			},
		},
		Identity: &armagricultureplatform.ManagedServiceIdentity{
			Type: to.Ptr(armagricultureplatform.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armagricultureplatform.UserAssignedIdentity{
				"key4955": {},
			},
		},
		SKU: &armagricultureplatform.SKU{
			Name:     to.Ptr("kfl"),
			Tier:     to.Ptr(armagricultureplatform.SKUTierFree),
			Size:     to.Ptr("r"),
			Family:   to.Ptr("xerdhxyjwrypvxphavgrtjphtohf"),
			Capacity: to.Ptr[int32](20),
		},
		Tags: map[string]*string{
			"key137": to.Ptr("oxwansfetzzgdwl"),
		},
		Location: to.Ptr("pkneuknooprpqirnugzwbkiie"),
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
	// res = armagricultureplatform.AgriServiceClientCreateOrUpdateResponse{
	// 	AgriServiceResource: &armagricultureplatform.AgriServiceResource{
	// 		Properties: &armagricultureplatform.AgriServiceResourceProperties{
	// 			ProvisioningState: to.Ptr(armagricultureplatform.ProvisioningStateSucceeded),
	// 			Config: &armagricultureplatform.AgriServiceConfig{
	// 				InstanceURI: to.Ptr("zwjvrpcyzxrsizerkzhwkergjrmxqe"),
	// 				Version: to.Ptr("cwhczenyyxtmslccbv"),
	// 				AppServiceResourceID: to.Ptr("icecbrltkdsresejoidqvlybwsbomotnbnmfa"),
	// 				CosmosDbResourceID: to.Ptr("ygzwcdwitjqshybczyukwhaxkomgimvqdmqsdsx"),
	// 				StorageAccountResourceID: to.Ptr("vdruhddgygpkcwngnvbstyitkzocrwnidpeowekohvisiprcmjzpe"),
	// 				KeyVaultResourceID: to.Ptr("egw"),
	// 				RedisCacheResourceID: to.Ptr("fxhznzcilbmqilgnryyazmhkssbhk"),
	// 			},
	// 			ManagedOnBehalfOfConfiguration: &armagricultureplatform.ManagedOnBehalfOfConfiguration{
	// 				MoboBrokerResources: []*armagricultureplatform.MoboBrokerResource{
	// 					{
	// 						ID: to.Ptr("bnthrkwfkfeorrzvtdxbfz"),
	// 					},
	// 				},
	// 			},
	// 			DataConnectorCredentials: []*armagricultureplatform.DataConnectorCredentialMap{
	// 				{
	// 					Key: to.Ptr("BackendAADApplicationCredentials"),
	// 					Value: &armagricultureplatform.DataConnectorCredentials{
	// 						ClientID: to.Ptr("dce298a8-1eec-481a-a8f9-a3cd5a8257b2"),
	// 					},
	// 				},
	// 			},
	// 			InstalledSolutions: []*armagricultureplatform.InstalledSolutionMap{
	// 				{
	// 					Key: to.Ptr("bayerAgPowered.cwum"),
	// 					Value: &armagricultureplatform.Solution{
	// 						ApplicationName: to.Ptr("bayerAgPowered.cwum"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Identity: &armagricultureplatform.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("16763be4-6022-406e-a950-fcd5018633ca"),
	// 			TenantID: to.Ptr("16763be4-6022-406e-a950-fcd5018633ca"),
	// 			Type: to.Ptr(armagricultureplatform.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armagricultureplatform.UserAssignedIdentity{
	// 				"key4955": &armagricultureplatform.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("16763be4-6022-406e-a950-fcd5018633ca"),
	// 					ClientID: to.Ptr("16763be4-6022-406e-a950-fcd5018633ca"),
	// 				},
	// 			},
	// 		},
	// 		SKU: &armagricultureplatform.SKU{
	// 			Name: to.Ptr("kfl"),
	// 			Tier: to.Ptr(armagricultureplatform.SKUTierFree),
	// 			Size: to.Ptr("r"),
	// 			Family: to.Ptr("xerdhxyjwrypvxphavgrtjphtohf"),
	// 			Capacity: to.Ptr[int32](20),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key137": to.Ptr("oxwansfetzzgdwl"),
	// 		},
	// 		Location: to.Ptr("pkneuknooprpqirnugzwbkiie"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789abc/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
	// 		Name: to.Ptr("mnvxvlitiwbndijhbmgiejz"),
	// 		Type: to.Ptr("fvvidjmentwsi"),
	// 		SystemData: &armagricultureplatform.SystemData{
	// 			CreatedBy: to.Ptr("gthxegufst"),
	// 			CreatedByType: to.Ptr(armagricultureplatform.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-16T11:47:57.784Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ovgqctuakdgemocstvwqmhyufe"),
	// 			LastModifiedByType: to.Ptr(armagricultureplatform.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-16T11:47:57.784Z"); return t}()),
	// 		},
	// 	},
	// }
}
