package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/DCIOperations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationListResult = armcustomerinsights.OperationListResult{
		// 	Value: []*armcustomerinsights.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read Azure Customer Insights Hub"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Hubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/write"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Azure Customer Insights Hub"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Hubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/delete"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Delete Azure Customer Insights Hub"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Hubs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/views/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read Azure Customer Insights App View"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights App Views"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/views/write"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Azure Customer Insights App View"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights App Views"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/views/delete"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Delete Azure Customer Insights App View"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights App Views"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/connectors/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read Azure Customer Insights Connector"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Connectors"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/connectors/write"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Azure Customer Insights Connector"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Connectors"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/connectors/delete"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Delete Azure Customer Insights Connector"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Connectors"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/connectors/mappings/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read Azure Customer Insights Connector Mapping"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Connector Mappings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/connectors/mappings/write"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Azure Customer Insights Connector Mapping"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Connector Mappings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/connectors/mappings/delete"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Delete Azure Customer Insights App View"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Connector Mappings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/interactions/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read Azure Customer Insights Interaction"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Interactions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/interactions/write"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Azure Customer Insights Interaction"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Interactions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/profiles/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read Azure Customer Insights Profile"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Profiles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/profiles/write"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Azure Customer Insights Profile"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Profiles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/kpi/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read Azure Customer Insights Key Performance Indicator"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Create or Update Customer Insights Key Performance Indicators"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/kpi/write"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Azure Customer Insights Key Performance Indicator"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Create or Update Customer Insights Key Performance Indicators"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/kpi/delete"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Delete Azure Customer Insights Key Performance Indicator"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Create or Update Customer Insights Key Performance Indicators"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/roleAssignments/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read Azure Customer Insights Rbac Assignment"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Rbac Assignments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/roleAssignments/write"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Azure Customer Insights Rbac Assignment"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Rbac Assignments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/roleAssignments/delete"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Delete Azure Customer Insights Rbac Assignment"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Rbac Assignments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/authorizationPolicies/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read Azure Customer Insights Shared Access Signature Policy"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Shared Access Signature Policies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/authorizationPolicies/write"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Create or Update Azure Customer Insights Shared Access Signature Policy"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Shared Access Signature Policies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/authorizationPolicies/delete"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Delete Azure Customer Insights Shared Access Signature Policy"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Shared Access Signature Policies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/authorizationPolicies/regeneratePrimaryKey/action"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Regenerate Azure Customer Insights Shared Access Signature Policy primary key"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Shared Access Signature Policies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/authorizationPolicies/regenerateSecondaryKey/action"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Regenerate Azure Customer Insights Shared Access Signature Policy secondary key"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Azure Customer Insights Shared Access Signature Policies"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read resource metric definitions"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Microsoft Azure Customer Insights Metric Definition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Read resource log definitions"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Microsoft Azure Customer Insights Log Definition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Diagnostic setting read"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Microsoft Azure Customer Insights Diagnostic Settings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.CustomerInsights/hubs/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armcustomerinsights.OperationDisplay{
		// 				Operation: to.Ptr("Diagnostic setting write"),
		// 				Provider: to.Ptr("Microsoft Azure Customer Insights"),
		// 				Resource: to.Ptr("Microsoft Azure Customer Insights Diagnostic Settings"),
		// 			},
		// 	}},
		// }
	}
}
