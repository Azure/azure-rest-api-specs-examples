package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getActivityInAModule.json
func ExampleActivityClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewActivityClient().Get(ctx, "rg", "myAutomationAccount33", "OmsCompositeResources", "Add-AzureRmAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Activity = armautomation.Activity{
	// 	Name: to.Ptr("Add-AzureRmAccount"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Add-AzureRmAccount"),
	// 	Properties: &armautomation.ActivityProperties{
	// 		Description: to.Ptr("The Add-AzureRmAcccount cmdlet adds an authenticated Azure account to use for Azure Resource Manager cmdlet requests.\n\nYou can use this authenticated account only with Azure Resource Manager cmdlets. To add an authenticated account for use with Service Management cmdlets, use the Add-AzureAccount or the Import-AzurePublishSettingsFile cmdlet."),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697Z"); return t}()),
	// 		Definition: to.Ptr(""),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697Z"); return t}()),
	// 		OutputTypes: []*armautomation.ActivityOutputType{
	// 			{
	// 				Name: to.Ptr("Microsoft.Azure.Commands.Profile.Models.PSAzureProfile"),
	// 				Type: to.Ptr("Microsoft.Azure.Commands.Profile.Models.PSAzureProfile"),
	// 		}},
	// 		ParameterSets: []*armautomation.ActivityParameterSet{
	// 			{
	// 				Name: to.Ptr("SubscriptionId"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("AccessToken"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("AccountId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("ApplicationId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("CertificateThumbprint"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Credential"),
	// 						Type: to.Ptr("System.Management.Automation.PSCredential"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SubscriptionId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(true),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("ServicePrincipal"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("Credential"),
	// 						Type: to.Ptr("System.Management.Automation.PSCredential"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("ServicePrincipal"),
	// 						Type: to.Ptr("System.Management.Automation.SwitchParameter"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SubscriptionId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(true),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SubscriptionName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(true),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("SubscriptionName"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("AccessToken"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("AccountId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("ApplicationId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("CertificateThumbprint"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Credential"),
	// 						Type: to.Ptr("System.Management.Automation.PSCredential"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SubscriptionName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(true),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("User"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("Credential"),
	// 						Type: to.Ptr("System.Management.Automation.PSCredential"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("AccessToken"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("AccessToken"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("AccountId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("ServicePrincipalCertificate"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("ApplicationId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("CertificateThumbprint"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("ServicePrincipal"),
	// 						Type: to.Ptr("System.Management.Automation.SwitchParameter"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 		}},
	// 	},
	// }
}
