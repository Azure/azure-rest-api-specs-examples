package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/ListSubscriptionSqlManagedInstance.json
func ExampleSQLManagedInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLManagedInstancesClient().NewListPager(nil)
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
		// page.SQLManagedInstanceListResult = armazurearcdata.SQLManagedInstanceListResult{
		// 	Value: []*armazurearcdata.SQLManagedInstance{
		// 		{
		// 			Name: to.Ptr("sqlManagedInstances1"),
		// 			Type: to.Ptr("Microsoft.AzureArcData/sqlManagedInstances"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/SqlManagedInstances/sqlManagedInstance1"),
		// 			SystemData: &armazurearcdata.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			ExtendedLocation: &armazurearcdata.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
		// 				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armazurearcdata.SQLManagedInstanceProperties{
		// 				ActiveDirectoryInformation: &armazurearcdata.ActiveDirectoryInformation{
		// 					KeytabInformation: &armazurearcdata.KeytabInformation{
		// 					},
		// 				},
		// 				Admin: to.Ptr("Admin user"),
		// 				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				ClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
		// 				EndTime: to.Ptr("Instance end time"),
		// 				ExtensionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
		// 				K8SRaw: &armazurearcdata.SQLManagedInstanceK8SRaw{
		// 					AdditionalProperties: map[string]any{
		// 						"additionalProperty": float64(1234),
		// 					},
		// 					Spec: &armazurearcdata.SQLManagedInstanceK8SSpec{
		// 						Replicas: to.Ptr[int32](1),
		// 						Scheduling: &armazurearcdata.K8SScheduling{
		// 							Default: &armazurearcdata.K8SSchedulingOptions{
		// 								Resources: &armazurearcdata.K8SResourceRequirements{
		// 									Limits: map[string]*string{
		// 										"additionalProperty": to.Ptr("additionalValue"),
		// 										"cpu": to.Ptr("1"),
		// 										"memory": to.Ptr("8Gi"),
		// 									},
		// 									Requests: map[string]*string{
		// 										"additionalProperty": to.Ptr("additionalValue"),
		// 										"cpu": to.Ptr("1"),
		// 										"memory": to.Ptr("8Gi"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 				LicenseType: to.Ptr(armazurearcdata.ArcSQLManagedInstanceLicenseTypeLicenseIncluded),
		// 				StartTime: to.Ptr("Instance start time"),
		// 			},
		// 			SKU: &armazurearcdata.SQLManagedInstanceSKU{
		// 				Name: to.Ptr("vCore"),
		// 				Dev: to.Ptr(true),
		// 				Tier: to.Ptr(armazurearcdata.SQLManagedInstanceSKUTierGeneralPurpose),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sqlManagedInstances2"),
		// 			Type: to.Ptr("Microsoft.AzureArcData/sqlManagedInstances"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/SqlManagedInstances/sqlManagedInstance2"),
		// 			SystemData: &armazurearcdata.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			ExtendedLocation: &armazurearcdata.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
		// 				Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
		// 			},
		// 			Properties: &armazurearcdata.SQLManagedInstanceProperties{
		// 				ActiveDirectoryInformation: &armazurearcdata.ActiveDirectoryInformation{
		// 					KeytabInformation: &armazurearcdata.KeytabInformation{
		// 					},
		// 				},
		// 				Admin: to.Ptr("Admin user"),
		// 				BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
		// 					Username: to.Ptr("username"),
		// 				},
		// 				ClusterID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s"),
		// 				EndTime: to.Ptr("Instance end time"),
		// 				ExtensionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/connectedk8s/providers/Microsoft.KubernetesConfiguration/extensions/extension"),
		// 				K8SRaw: &armazurearcdata.SQLManagedInstanceK8SRaw{
		// 					AdditionalProperties: map[string]any{
		// 						"additionalProperty": float64(1234),
		// 					},
		// 					Spec: &armazurearcdata.SQLManagedInstanceK8SSpec{
		// 						Replicas: to.Ptr[int32](1),
		// 						Scheduling: &armazurearcdata.K8SScheduling{
		// 							Default: &armazurearcdata.K8SSchedulingOptions{
		// 								Resources: &armazurearcdata.K8SResourceRequirements{
		// 									Limits: map[string]*string{
		// 										"additionalProperty": to.Ptr("additionalValue"),
		// 										"cpu": to.Ptr("1"),
		// 										"memory": to.Ptr("8Gi"),
		// 									},
		// 									Requests: map[string]*string{
		// 										"additionalProperty": to.Ptr("additionalValue"),
		// 										"cpu": to.Ptr("1"),
		// 										"memory": to.Ptr("8Gi"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 				LicenseType: to.Ptr(armazurearcdata.ArcSQLManagedInstanceLicenseTypeLicenseIncluded),
		// 				StartTime: to.Ptr("Instance start time"),
		// 			},
		// 			SKU: &armazurearcdata.SQLManagedInstanceSKU{
		// 				Name: to.Ptr("vCore"),
		// 				Dev: to.Ptr(true),
		// 				Tier: to.Ptr(armazurearcdata.SQLManagedInstanceSKUTierGeneralPurpose),
		// 			},
		// 	}},
		// }
	}
}
