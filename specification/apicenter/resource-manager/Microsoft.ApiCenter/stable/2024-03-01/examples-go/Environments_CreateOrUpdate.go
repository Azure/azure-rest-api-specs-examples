package armapicenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apicenter/armapicenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Environments_CreateOrUpdate.json
func ExampleEnvironmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapicenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnvironmentsClient().CreateOrUpdate(ctx, "contoso-resources", "contoso", "default", "public", armapicenter.Environment{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Environment = armapicenter.Environment{
	// 	Name: to.Ptr("public"),
	// 	Type: to.Ptr("Microsoft.ApiCenter/services/workspaces/environments"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ApiCenter/services/contoso/workspaces/default/environments/public"),
	// 	SystemData: &armapicenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T18:27:09.128Z"); return t}()),
	// 	},
	// 	Properties: &armapicenter.EnvironmentProperties{
	// 		Description: to.Ptr("The primary Azure API Management service for the European division of Contoso."),
	// 		Kind: to.Ptr(armapicenter.EnvironmentKindProduction),
	// 		Onboarding: &armapicenter.Onboarding{
	// 			DeveloperPortalURI: []*string{
	// 				to.Ptr("https://developer.contoso.com")},
	// 				Instructions: to.Ptr("Sign in or sign up in the specified developer portal to request API access. You must complete the internal privacy training for your account to be approved."),
	// 			},
	// 			Server: &armapicenter.EnvironmentServer{
	// 				Type: to.Ptr(armapicenter.EnvironmentServerTypeAzureAPIManagement),
	// 				ManagementPortalURI: []*string{
	// 					to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resources/providers/Microsoft.ApiManagement/service/contoso")},
	// 				},
	// 				Title: to.Ptr("Contoso Europe Azure API Management"),
	// 			},
	// 		}
}
