package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: 2025-03-01-preview/IacProfile_List.json
func ExampleIacProfilesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("a0a37f63-7183-4e86-9ac7-ce8036a3ed31", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIacProfilesClient().NewListPager(nil)
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
		// page = armdevhub.IacProfilesClientListResponse{
		// 	IacProfileListResult: armdevhub.IacProfileListResult{
		// 		NextLink: to.Ptr("https://management.azure.com:443/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DevHub/iacProfiles?api-version=2018-05-05-preview&timeRange=1&$skipToken="),
		// 		Value: []*armdevhub.IacProfile{
		// 			{
		// 				Name: to.Ptr("profilename"),
		// 				Type: to.Ptr("Microsoft.DevHub/iacProfiles"),
		// 				ID: to.Ptr("/subscriptions/a0a37f63-7183-4e86-9ac7-ce8036a3ed31/resourceGroups/resourceGroup1/providers/Microsoft.DevHub/iacProfiles/profilename"),
		// 				Location: to.Ptr(""),
		// 				Properties: &armdevhub.IacProfileProperties{
		// 					GithubProfile: &armdevhub.IacGitHubProfile{
		// 						AuthStatus: to.Ptr(armdevhub.AuthorizationStatusAuthorized),
		// 						BranchName: to.Ptr("aks-devhub-ybvmw"),
		// 						PrStatus: to.Ptr(armdevhub.PullRequestStatusSubmitted),
		// 						PullNumber: to.Ptr[int32](8),
		// 						RepositoryMainBranch: to.Ptr("main"),
		// 						RepositoryName: to.Ptr("localtest"),
		// 						RepositoryOwner: to.Ptr("owner"),
		// 					},
		// 					Stages: []*armdevhub.StageProperties{
		// 						{
		// 							Dependencies: []*string{
		// 							},
		// 							GitEnvironment: to.Ptr("Terraform"),
		// 							StageName: to.Ptr("dev"),
		// 						},
		// 						{
		// 							Dependencies: []*string{
		// 								to.Ptr("dev"),
		// 							},
		// 							GitEnvironment: to.Ptr("Terraform"),
		// 							StageName: to.Ptr("qa"),
		// 						},
		// 						{
		// 							Dependencies: []*string{
		// 								to.Ptr("qa"),
		// 							},
		// 							GitEnvironment: to.Ptr("Terraform"),
		// 							StageName: to.Ptr("prod"),
		// 						},
		// 					},
		// 					Templates: []*armdevhub.IacTemplateProperties{
		// 						{
		// 							InstanceName: to.Ptr("contoso"),
		// 							InstanceStage: to.Ptr("dev"),
		// 							SourceResourceID: to.Ptr("/subscriptions/xxxx/resourceGroups/xxxx"),
		// 							TemplateDetails: []*armdevhub.IacTemplateDetails{
		// 								{
		// 									Count: to.Ptr[int32](1),
		// 									NamingConvention: to.Ptr("$sitid-hci"),
		// 									ProductName: to.Ptr("HCI"),
		// 								},
		// 								{
		// 									Count: to.Ptr[int32](1),
		// 									NamingConvention: to.Ptr("$sitid-aks"),
		// 									ProductName: to.Ptr("AKSarc"),
		// 								},
		// 							},
		// 							TemplateName: to.Ptr("base"),
		// 						},
		// 					},
		// 					TerraformProfile: &armdevhub.TerraformProfile{
		// 						StorageAccountName: to.Ptr("hybridiac"),
		// 						StorageAccountResourceGroup: to.Ptr("hybrid-iac"),
		// 						StorageAccountSubscription: to.Ptr("subscription"),
		// 						StorageContainerName: to.Ptr("hybridiac"),
		// 					},
		// 				},
		// 				SystemData: &armdevhub.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-16T09:56:23.933651400Z"); return t}()),
		// 					CreatedBy: to.Ptr(""),
		// 					CreatedByType: to.Ptr(armdevhub.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-16T14:32:00.783994200Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("aks-dev-hub-devhub"),
		// 					LastModifiedByType: to.Ptr(armdevhub.CreatedByTypeApplication),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
