package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/EnvironmentTypes_Put.json
func ExampleEnvironmentTypesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnvironmentTypesClient().CreateOrUpdate(ctx, "rg1", "Contoso", "DevTest", armdevcenter.EnvironmentType{
		Properties: &armdevcenter.EnvironmentTypeProperties{
			DisplayName: to.Ptr("Dev"),
		},
		Tags: map[string]*string{
			"Owner": to.Ptr("superuser"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EnvironmentType = armdevcenter.EnvironmentType{
	// 	Name: to.Ptr("DevTest"),
	// 	Type: to.Ptr("Microsoft.DevCenter/devcenters/environmenttypes"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/environmentTypes/DevTest"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 		CreatedBy: to.Ptr("User1@contoso.com"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User1@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Properties: &armdevcenter.EnvironmentTypeProperties{
	// 		DisplayName: to.Ptr("Dev"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Owner": to.Ptr("superuser"),
	// 		"hidden-title": to.Ptr("Dev"),
	// 	},
	// }
}
