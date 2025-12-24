package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/455d20a5e76d8184f7cff960501a57e1f88986b7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/KubernetesVersions_List.json
func ExampleManagedClustersClient_ListKubernetesVersions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClustersClient().ListKubernetesVersions(ctx, "location1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KubernetesVersionListResult = armcontainerservice.KubernetesVersionListResult{
	// 	Values: []*armcontainerservice.KubernetesVersion{
	// 		{
	// 			Capabilities: &armcontainerservice.KubernetesVersionCapabilities{
	// 				SupportPlan: []*armcontainerservice.KubernetesSupportPlan{
	// 					to.Ptr(armcontainerservice.KubernetesSupportPlanKubernetesOfficial)},
	// 				},
	// 				PatchVersions: map[string]*armcontainerservice.KubernetesPatchVersion{
	// 					"1.23.12": &armcontainerservice.KubernetesPatchVersion{
	// 						Upgrades: []*string{
	// 							to.Ptr("1.23.15"),
	// 							to.Ptr("1.24.6"),
	// 							to.Ptr("1.24.9")},
	// 						},
	// 						"1.23.15": &armcontainerservice.KubernetesPatchVersion{
	// 							Upgrades: []*string{
	// 								to.Ptr("1.24.6"),
	// 								to.Ptr("1.24.9")},
	// 							},
	// 						},
	// 						Version: to.Ptr("1.23"),
	// 					},
	// 					{
	// 						Capabilities: &armcontainerservice.KubernetesVersionCapabilities{
	// 							SupportPlan: []*armcontainerservice.KubernetesSupportPlan{
	// 								to.Ptr(armcontainerservice.KubernetesSupportPlanKubernetesOfficial)},
	// 							},
	// 							IsDefault: to.Ptr(true),
	// 							PatchVersions: map[string]*armcontainerservice.KubernetesPatchVersion{
	// 								"1.24.6": &armcontainerservice.KubernetesPatchVersion{
	// 									Upgrades: []*string{
	// 										to.Ptr("1.24.9"),
	// 										to.Ptr("1.25.4"),
	// 										to.Ptr("1.25.5")},
	// 									},
	// 									"1.24.9": &armcontainerservice.KubernetesPatchVersion{
	// 										Upgrades: []*string{
	// 											to.Ptr("1.25.4"),
	// 											to.Ptr("1.25.5")},
	// 										},
	// 									},
	// 									Version: to.Ptr("1.24"),
	// 								},
	// 								{
	// 									Capabilities: &armcontainerservice.KubernetesVersionCapabilities{
	// 										SupportPlan: []*armcontainerservice.KubernetesSupportPlan{
	// 											to.Ptr(armcontainerservice.KubernetesSupportPlanKubernetesOfficial)},
	// 										},
	// 										PatchVersions: map[string]*armcontainerservice.KubernetesPatchVersion{
	// 											"1.25.4": &armcontainerservice.KubernetesPatchVersion{
	// 												Upgrades: []*string{
	// 													to.Ptr("1.25.5"),
	// 													to.Ptr("1.26.0")},
	// 												},
	// 												"1.25.5": &armcontainerservice.KubernetesPatchVersion{
	// 													Upgrades: []*string{
	// 														to.Ptr("1.26.0")},
	// 													},
	// 												},
	// 												Version: to.Ptr("1.25"),
	// 											},
	// 											{
	// 												Capabilities: &armcontainerservice.KubernetesVersionCapabilities{
	// 													SupportPlan: []*armcontainerservice.KubernetesSupportPlan{
	// 														to.Ptr(armcontainerservice.KubernetesSupportPlanKubernetesOfficial)},
	// 													},
	// 													IsPreview: to.Ptr(true),
	// 													PatchVersions: map[string]*armcontainerservice.KubernetesPatchVersion{
	// 														"1.26.0": &armcontainerservice.KubernetesPatchVersion{
	// 															Upgrades: []*string{
	// 															},
	// 														},
	// 													},
	// 													Version: to.Ptr("1.26"),
	// 											}},
	// 										}
}
