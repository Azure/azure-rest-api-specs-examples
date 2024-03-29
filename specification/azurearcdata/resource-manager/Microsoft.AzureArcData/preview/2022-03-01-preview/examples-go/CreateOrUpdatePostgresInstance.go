package armazurearcdata_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdatePostgresInstance.json
func ExamplePostgresInstancesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurearcdata.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPostgresInstancesClient().BeginCreate(ctx, "testrg", "testpostgresInstance", armazurearcdata.PostgresInstance{
		Location: to.Ptr("eastus"),
		ExtendedLocation: &armazurearcdata.ExtendedLocation{
			Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
			Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
		},
		Properties: &armazurearcdata.PostgresInstanceProperties{
			Admin: to.Ptr("admin"),
			BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
				Password: to.Ptr("********"),
				Username: to.Ptr("username"),
			},
			DataControllerID: to.Ptr("dataControllerId"),
			K8SRaw: map[string]any{
				"apiVersion": "apiVersion",
				"kind":       "postgresql-12",
				"metadata": map[string]any{
					"name":              "pg1",
					"creationTimestamp": "2020-08-25T14:55:10Z",
					"generation":        float64(1),
					"namespace":         "test",
					"resourceVersion":   "527780",
					"selfLink":          "/apis/arcdata.microsoft.com/v1alpha1/namespaces/test/postgresql-12s/pg1",
					"uid":               "1111aaaa-ffff-ffff-ffff-99999aaaaaaa",
				},
				"spec": map[string]any{
					"backups": map[string]any{
						"deltaMinutes": float64(3),
						"fullMinutes":  float64(10),
						"tiers": []any{
							map[string]any{
								"retention": map[string]any{
									"maximums": []any{
										"6",
										"512MB",
									},
									"minimums": []any{
										"3",
									},
								},
								"storage": map[string]any{
									"volumeSize": "1Gi",
								},
							},
						},
					},
					"engine": map[string]any{
						"extensions": []any{
							map[string]any{
								"name": "citus",
							},
						},
					},
					"scale": map[string]any{
						"shards": float64(3),
					},
					"scheduling": map[string]any{
						"default": map[string]any{
							"resources": map[string]any{
								"requests": map[string]any{
									"memory": "256Mi",
								},
							},
						},
					},
					"service": map[string]any{
						"type": "NodePort",
					},
					"storage": map[string]any{
						"data": map[string]any{
							"className": "local-storage",
							"size":      "5Gi",
						},
						"logs": map[string]any{
							"className": "local-storage",
							"size":      "5Gi",
						},
					},
				},
				"status": map[string]any{
					"externalEndpoint": nil,
					"readyPods":        "4/4",
					"state":            "Ready",
				},
			},
		},
		SKU: &armazurearcdata.PostgresInstanceSKU{
			Name: to.Ptr("default"),
			Dev:  to.Ptr(true),
			Tier: to.Ptr("Hyperscale"),
		},
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
	// res.PostgresInstance = armazurearcdata.PostgresInstance{
	// 	Name: to.Ptr("testpostgresInstance"),
	// 	Type: to.Ptr("Microsoft.AzureArcData/PostgresInstance"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.AzureArcData/PostgresInstance/testpostgresInstance"),
	// 	SystemData: &armazurearcdata.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurearcdata.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("northeurope"),
	// 	Tags: map[string]*string{
	// 		"mytag": to.Ptr("myval"),
	// 	},
	// 	ExtendedLocation: &armazurearcdata.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation"),
	// 		Type: to.Ptr(armazurearcdata.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armazurearcdata.PostgresInstanceProperties{
	// 		BasicLoginInformation: &armazurearcdata.BasicLoginInformation{
	// 			Username: to.Ptr("username"),
	// 		},
	// 	},
	// 	SKU: &armazurearcdata.PostgresInstanceSKU{
	// 		Name: to.Ptr("default"),
	// 		Dev: to.Ptr(true),
	// 		Tier: to.Ptr("Hyperscale"),
	// 	},
	// }
}
