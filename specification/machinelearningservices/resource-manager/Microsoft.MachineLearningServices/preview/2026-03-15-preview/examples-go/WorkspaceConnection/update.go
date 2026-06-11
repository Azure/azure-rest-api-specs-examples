package armmachinelearning_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
)

// Generated from example definition: 2026-03-15-preview/WorkspaceConnection/update.json
func ExampleWorkspaceConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmachinelearning.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceConnectionsClient().Update(ctx, "test-rg", "workspace-1", "connection-1", armmachinelearning.WorkspaceConnectionUpdateParameter{
		Properties: &armmachinelearning.AccessKeyAuthTypeWorkspaceConnectionProperties{
			AuthType: to.Ptr(armmachinelearning.ConnectionAuthTypeAccessKey),
			Category: to.Ptr(armmachinelearning.ConnectionCategoryADLSGen2),
			Credentials: &armmachinelearning.WorkspaceConnectionAccessKey{
				AccessKeyID:     to.Ptr("some_string"),
				SecretAccessKey: to.Ptr("some_string"),
			},
			ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z"); return t }()),
			Metadata:   map[string]*string{},
			Target:     to.Ptr("some_string"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmachinelearning.WorkspaceConnectionsClientUpdateResponse{
	// 	WorkspaceConnectionPropertiesV2BasicResource: armmachinelearning.WorkspaceConnectionPropertiesV2BasicResource{
	// 		Name: to.Ptr("connection-1"),
	// 		Type: to.Ptr("Microsoft.MachineLearningServices/workspaces/connections"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup-1/providers/Microsoft.MachineLearningServices/workspaces/workspace-1/connections/connection-1"),
	// 		Properties: &armmachinelearning.AccessKeyAuthTypeWorkspaceConnectionProperties{
	// 			AuthType: to.Ptr(armmachinelearning.ConnectionAuthTypeAccessKey),
	// 			Category: to.Ptr(armmachinelearning.ConnectionCategoryADLSGen2),
	// 			ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z"); return t}()),
	// 			Metadata: map[string]*string{
	// 			},
	// 			Target: to.Ptr("some_string"),
	// 		},
	// 		SystemData: &armmachinelearning.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z"); return t}()),
	// 			CreatedBy: to.Ptr("some_string"),
	// 			CreatedByType: to.Ptr(armmachinelearning.CreatedByTypeManagedIdentity),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T00:00:00Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("some_string"),
	// 			LastModifiedByType: to.Ptr(armmachinelearning.CreatedByTypeApplication),
	// 		},
	// 	},
	// }
}
