package armdevhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/34f5146bc945549d087d38a8a593c0a5f475ad7f/specification/developerhub/resource-manager/Microsoft.DevHub/preview/2024-05-01-preview/examples/IacProfile_Get.json
func ExampleIacProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIacProfilesClient().Get(ctx, "resourceGroup1", "iacprofile", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IacProfile = armdevhub.IacProfile{
	// 	Name: to.Ptr("iacprofile"),
	// 	Type: to.Ptr("Micfosoft.DevHub/iacProfiles"),
	// 	ID: to.Ptr("/subscriptions/a0a37f63-7183-4e86-9ac7-ce8036a3ed31/resourceGroups/resourceGroup1/providers/Microsoft.DevHub/iacProfiles/iacprofile"),
	// 	SystemData: &armdevhub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-16T09:56:23.933Z"); return t}()),
	// 		CreatedBy: to.Ptr(""),
	// 		CreatedByType: to.Ptr(armdevhub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-16T14:35:15.206Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("aks-dev-hub-devhub"),
	// 		LastModifiedByType: to.Ptr(armdevhub.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("location1"),
	// 	Tags: map[string]*string{
	// 		"appname": to.Ptr("testapp"),
	// 	},
	// 	Properties: &armdevhub.IacProfileProperties{
	// 		GithubProfile: &armdevhub.IacGitHubProfile{
	// 			AuthStatus: to.Ptr(armdevhub.AuthorizationStatusAuthorized),
	// 			BranchName: to.Ptr("aks-devhub-ybvmw"),
	// 			PrStatus: to.Ptr(armdevhub.PullRequestStatusSubmitted),
	// 			PullNumber: to.Ptr[int32](8),
	// 			RepositoryMainBranch: to.Ptr("main"),
	// 			RepositoryName: to.Ptr("localtest"),
	// 			RepositoryOwner: to.Ptr("owner"),
	// 		},
	// 		Stages: []*armdevhub.StageProperties{
	// 			{
	// 				Dependencies: []*string{
	// 				},
	// 				GitEnvironment: to.Ptr("Terraform"),
	// 				StageName: to.Ptr("dev"),
	// 			},
	// 			{
	// 				Dependencies: []*string{
	// 					to.Ptr("dev")},
	// 					GitEnvironment: to.Ptr("Terraform"),
	// 					StageName: to.Ptr("qa"),
	// 				},
	// 				{
	// 					Dependencies: []*string{
	// 						to.Ptr("qa")},
	// 						GitEnvironment: to.Ptr("Terraform"),
	// 						StageName: to.Ptr("prod"),
	// 				}},
	// 				Templates: []*armdevhub.IacTemplateProperties{
	// 					{
	// 						InstanceName: to.Ptr("contoso"),
	// 						InstanceStage: to.Ptr("dev"),
	// 						SourceResourceID: to.Ptr("/subscriptions/xxxx/resourceGroups/xxxx"),
	// 						TemplateDetails: []*armdevhub.IacTemplateDetails{
	// 							{
	// 								Count: to.Ptr[int32](1),
	// 								NamingConvention: to.Ptr("$sitid-hci"),
	// 								ProductName: to.Ptr("HCI"),
	// 							},
	// 							{
	// 								Count: to.Ptr[int32](1),
	// 								NamingConvention: to.Ptr("$sitid-aks"),
	// 								ProductName: to.Ptr("AKSarc"),
	// 						}},
	// 						TemplateName: to.Ptr("base"),
	// 				}},
	// 				TerraformProfile: &armdevhub.TerraformProfile{
	// 					StorageAccountName: to.Ptr("hybridiac"),
	// 					StorageAccountResourceGroup: to.Ptr("hybrid-iac"),
	// 					StorageAccountSubscription: to.Ptr("subscriptions"),
	// 					StorageContainerName: to.Ptr("hybridiac"),
	// 				},
	// 			},
	// 		}
}
