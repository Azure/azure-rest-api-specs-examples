package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Access_ClusterList.json
func ExampleAccessClient_ListClusters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListClusters(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"pageSize":  to.Ptr("10"),
			"pageToken": to.Ptr("asc4fts4ft"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessListClusterSuccessResponse = armconfluent.AccessListClusterSuccessResponse{
	// 	Data: []*armconfluent.ClusterRecord{
	// 		{
	// 			ID: to.Ptr("dlz-f3a90de"),
	// 			Kind: to.Ptr("Cluster"),
	// 			Metadata: &armconfluent.MetadataEntity{
	// 				CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-abc123/cloud-cluster=lkc-12345"),
	// 				Self: to.Ptr("https://api.confluent.cloud/cmk/v2/clusters/lkc-12345"),
	// 				UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 			},
	// 			Spec: &armconfluent.ClusterSpecEntity{
	// 				APIEndpoint: to.Ptr("https://pkac-00000.us-west-2.aws.confluent.cloud"),
	// 				Availability: to.Ptr("SINGLE_ZONE"),
	// 				Byok: &armconfluent.ClusterByokEntity{
	// 					ID: to.Ptr("cck-00000"),
	// 					Related: to.Ptr("https://api.confluent.cloud/byok/v1/keys/cck-00000"),
	// 					ResourceName: to.Ptr("https://api.confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/key=cck-00000"),
	// 				},
	// 				Cloud: to.Ptr("GCP"),
	// 				Config: &armconfluent.ClusterConfigEntity{
	// 					Kind: to.Ptr("Basic"),
	// 				},
	// 				DisplayName: to.Ptr("ProdKafkaCluster"),
	// 				Environment: &armconfluent.ClusterEnvironmentEntity{
	// 					Environment: to.Ptr("string"),
	// 					ID: to.Ptr("env-00000"),
	// 					Related: to.Ptr("https://api.confluent.cloud/v2/environments/env-00000"),
	// 					ResourceName: to.Ptr("https://api.confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-00000"),
	// 				},
	// 				HTTPEndpoint: to.Ptr("https://lkc-00000-00000.us-central1.gcp.glb.confluent.cloud"),
	// 				KafkaBootstrapEndpoint: to.Ptr("lkc-00000-00000.us-central1.gcp.glb.confluent.cloud:9092"),
	// 				Network: &armconfluent.ClusterNetworkEntity{
	// 					Environment: to.Ptr("string"),
	// 					ID: to.Ptr("n-00000"),
	// 					Related: to.Ptr("https://api.confluent.cloud/networking/v1/networks/n-00000"),
	// 					ResourceName: to.Ptr("https://api.confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-abc123/network=n-00000"),
	// 				},
	// 				Region: to.Ptr("us-east4"),
	// 			},
	// 			Status: &armconfluent.ClusterStatusEntity{
	// 				Cku: to.Ptr[int32](2),
	// 				Phase: to.Ptr("PROVISIONED"),
	// 			},
	// 	}},
	// 	Kind: to.Ptr("ClusterList"),
	// 	Metadata: &armconfluent.ListMetadata{
	// 		First: to.Ptr("https://api.confluent.cloud/cmk/v2/clusters"),
	// 		Last: to.Ptr("https://api.confluent.cloud/cmk/v2/clusters?page_token=bcAOehAY8F16YD84Z1wT"),
	// 		Next: to.Ptr("https://api.confluent.cloud/cmk/v2/clusters?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 		Prev: to.Ptr("https://api.confluent.cloud/cmk/v2/clusters?page_token=YIXRY97wWYmwzrax4dld"),
	// 		TotalSize: to.Ptr[int32](123),
	// 	},
	// }
}
