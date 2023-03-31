package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ProfilesGet.json
func ExampleProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProfilesClient().Get(ctx, "TestHubRG", "sdkTestHub", "TestProfileType396", &armcustomerinsights.ProfilesClientGetOptions{LocaleCode: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProfileResourceFormat = armcustomerinsights.ProfileResourceFormat{
	// 	Name: to.Ptr("azSdkTestHub/TestProfileType396"),
	// 	Type: to.Ptr("Microsoft.CustomerInsights/hubs/profiles"),
	// 	ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/azSdkTestHub/profiles/TestProfileType396"),
	// 	Properties: &armcustomerinsights.ProfileTypeDefinition{
	// 		LargeImage: to.Ptr("\\\\Images\\\\LargeImage"),
	// 		MediumImage: to.Ptr("\\\\Images\\\\MediumImage"),
	// 		SmallImage: to.Ptr("\\\\Images\\\\smallImage"),
	// 		APIEntitySetName: to.Ptr("TestProfileType396"),
	// 		EntityType: to.Ptr(armcustomerinsights.EntityTypesProfile),
	// 		Fields: []*armcustomerinsights.PropertyDefinition{
	// 			{
	// 				FieldName: to.Ptr("Id"),
	// 				FieldType: to.Ptr("Edm.String"),
	// 				IsArray: to.Ptr(false),
	// 				IsEnum: to.Ptr(false),
	// 				IsFlagEnum: to.Ptr(false),
	// 				IsImage: to.Ptr(false),
	// 				IsLocalizedString: to.Ptr(false),
	// 				IsName: to.Ptr(false),
	// 				IsRequired: to.Ptr(true),
	// 				PropertyID: to.Ptr("id1"),
	// 			},
	// 			{
	// 				FieldName: to.Ptr("ProfileId"),
	// 				FieldType: to.Ptr("Edm.String"),
	// 				IsArray: to.Ptr(false),
	// 				IsEnum: to.Ptr(false),
	// 				IsFlagEnum: to.Ptr(false),
	// 				IsImage: to.Ptr(false),
	// 				IsLocalizedString: to.Ptr(false),
	// 				IsName: to.Ptr(false),
	// 				IsRequired: to.Ptr(true),
	// 				PropertyID: to.Ptr("id1"),
	// 			},
	// 			{
	// 				FieldName: to.Ptr("LastName"),
	// 				FieldType: to.Ptr("Edm.String"),
	// 				IsArray: to.Ptr(false),
	// 				IsEnum: to.Ptr(false),
	// 				IsFlagEnum: to.Ptr(false),
	// 				IsImage: to.Ptr(false),
	// 				IsLocalizedString: to.Ptr(false),
	// 				IsName: to.Ptr(false),
	// 				IsRequired: to.Ptr(true),
	// 				PropertyID: to.Ptr("id1"),
	// 			},
	// 			{
	// 				FieldName: to.Ptr("TestProfileType396"),
	// 				FieldType: to.Ptr("Edm.String"),
	// 				IsArray: to.Ptr(false),
	// 				IsEnum: to.Ptr(false),
	// 				IsFlagEnum: to.Ptr(false),
	// 				IsImage: to.Ptr(false),
	// 				IsLocalizedString: to.Ptr(false),
	// 				IsName: to.Ptr(false),
	// 				IsRequired: to.Ptr(true),
	// 				PropertyID: to.Ptr("id1"),
	// 			},
	// 			{
	// 				FieldName: to.Ptr("SavingAccountBalance"),
	// 				FieldType: to.Ptr("Edm.Int32"),
	// 				IsArray: to.Ptr(false),
	// 				IsEnum: to.Ptr(false),
	// 				IsFlagEnum: to.Ptr(false),
	// 				IsImage: to.Ptr(false),
	// 				IsLocalizedString: to.Ptr(false),
	// 				IsName: to.Ptr(false),
	// 				IsRequired: to.Ptr(true),
	// 				PropertyID: to.Ptr("id1"),
	// 		}},
	// 		InstancesCount: to.Ptr[int32](0),
	// 		ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
	// 		SchemaItemTypeLink: to.Ptr("SchemaItemTypeLink"),
	// 		TenantID: to.Ptr("azsdktesthub"),
	// 		TypeName: to.Ptr("TestProfileType396"),
	// 		StrongIDs: []*armcustomerinsights.StrongID{
	// 			{
	// 				KeyPropertyNames: []*string{
	// 					to.Ptr("Id"),
	// 					to.Ptr("savingAccountBalance")},
	// 					StrongIDName: to.Ptr("Id"),
	// 				},
	// 				{
	// 					KeyPropertyNames: []*string{
	// 						to.Ptr("ProfileId"),
	// 						to.Ptr("LastName")},
	// 						StrongIDName: to.Ptr("ProfileId"),
	// 					},
	// 					{
	// 						KeyPropertyNames: []*string{
	// 							to.Ptr("ProfileId")},
	// 							StrongIDName: to.Ptr("ProfileId"),
	// 					}},
	// 				},
	// 			}
}
