package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v11"
)

// Generated from example definition: 2018-06-01/Factories_ListByResourceGroup.json
func ExampleFactoriesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFactoriesClient().NewListByResourceGroupPager("exampleResourceGroup", nil)
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
		// page = armdatafactory.FactoriesClientListByResourceGroupResponse{
		// 	FactoryListResponse: armdatafactory.FactoryListResponse{
		// 		Value: []*armdatafactory.Factory{
		// 			{
		// 				Name: to.Ptr("exampleFactoryName-linked"),
		// 				Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 				ETag: to.Ptr("\"00008a02-0000-0000-0000-5b237f270000\""),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName-linked"),
		// 				Identity: &armdatafactory.FactoryIdentity{
		// 					Type: to.Ptr(armdatafactory.FactoryIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("10743799-44d2-42fe-8c4d-5bc5c51c0684"),
		// 					TenantID: to.Ptr("12345678-1234-1234-1234-123456789abc"),
		// 				},
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armdatafactory.FactoryProperties{
		// 					CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-15T08:56:07.1828318Z"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					Version: to.Ptr("2017-09-01-preview"),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("FactoryToUpgrade"),
		// 				Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 				ETag: to.Ptr("\"00003d04-0000-0000-0000-5b28962f0000\""),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/factorytoupgrade"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armdatafactory.FactoryProperties{
		// 					CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:35:35.7133828Z"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					Version: to.Ptr("2018-06-01"),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("exampleFactoryName"),
		// 				Type: to.Ptr("Microsoft.DataFactory/factories"),
		// 				ETag: to.Ptr("\"00004004-0000-0000-0000-5b28979e0000\""),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armdatafactory.FactoryProperties{
		// 					CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-19T05:41:50.0041314Z"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					RepoConfiguration: &armdatafactory.FactoryVSTSConfiguration{
		// 						Type: to.Ptr("FactoryVSTSConfiguration"),
		// 						AccountName: to.Ptr("ADF"),
		// 						CollaborationBranch: to.Ptr("master"),
		// 						LastCommitID: to.Ptr(""),
		// 						ProjectName: to.Ptr("project"),
		// 						RepositoryName: to.Ptr("repo"),
		// 						RootFolder: to.Ptr("/"),
		// 						TenantID: to.Ptr(""),
		// 					},
		// 					Version: to.Ptr("2018-06-01"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"exampleTag": to.Ptr("exampleValue"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
