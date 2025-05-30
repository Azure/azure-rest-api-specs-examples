package armservicelinker_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad60d7f8eba124edc6999677c55aba2184e303b0/specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/PutDryrun.json
func ExampleLinkersClient_BeginCreateDryrun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicelinker.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLinkersClient().BeginCreateDryrun(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app", "dryrunName", armservicelinker.DryrunResource{
		Properties: &armservicelinker.DryrunProperties{
			Parameters: &armservicelinker.CreateOrUpdateDryrunParameters{
				ActionName: to.Ptr(armservicelinker.DryrunActionNameCreateOrUpdate),
				AuthInfo: &armservicelinker.SecretAuthInfo{
					AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
					Name:     to.Ptr("name"),
					SecretInfo: &armservicelinker.ValueSecretInfo{
						SecretType: to.Ptr(armservicelinker.SecretTypeRawValue),
						Value:      to.Ptr("secret"),
					},
				},
				TargetService: &armservicelinker.AzureResource{
					Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
					ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
				},
			},
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
	// res.DryrunResource = armservicelinker.DryrunResource{
	// 	Name: to.Ptr("dryrunName"),
	// 	Type: to.Ptr("Microsoft.ServiceLinker/dryruns"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app/providers/Microsoft.ServiceLinker/dryruns/dryrunName"),
	// 	Properties: &armservicelinker.DryrunProperties{
	// 		OperationPreviews: []*armservicelinker.DryrunOperationPreview{
	// 			{
	// 				Name: to.Ptr("configFirewallRule"),
	// 				Description: to.Ptr("Config firewall rule for target service to allow source service access"),
	// 				Action: to.Ptr("Microsoft.DocumentDb/databaseAccounts/write"),
	// 				OperationType: to.Ptr(armservicelinker.DryrunPreviewOperationTypeConfigNetwork),
	// 				Scope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc"),
	// 		}},
	// 		Parameters: &armservicelinker.CreateOrUpdateDryrunParameters{
	// 			ActionName: to.Ptr(armservicelinker.DryrunActionNameCreateOrUpdate),
	// 			AuthInfo: &armservicelinker.SecretAuthInfo{
	// 				AuthType: to.Ptr(armservicelinker.AuthTypeSecret),
	// 				Name: to.Ptr("name"),
	// 			},
	// 			TargetService: &armservicelinker.AzureResource{
	// 				Type: to.Ptr(armservicelinker.TargetServiceTypeAzureResource),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db"),
	// 			},
	// 		},
	// 		PrerequisiteResults: []armservicelinker.DryrunPrerequisiteResultClassification{
	// 			&armservicelinker.BasicErrorDryrunPrerequisiteResult{
	// 				Type: to.Ptr(armservicelinker.DryrunPrerequisiteResultTypeBasicError),
	// 				Code: to.Ptr("ResourceNotFound"),
	// 				Message: to.Ptr("Target resource is not found"),
	// 			},
	// 			&armservicelinker.PermissionsMissingDryrunPrerequisiteResult{
	// 				Type: to.Ptr(armservicelinker.DryrunPrerequisiteResultTypePermissionsMissing),
	// 				Permissions: []*string{
	// 					to.Ptr("Microsoft.DocumentDb/databaseAccounts/write")},
	// 					Scope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc"),
	// 			}},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	}
}
