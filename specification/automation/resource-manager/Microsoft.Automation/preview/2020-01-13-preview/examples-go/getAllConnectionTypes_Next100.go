package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getAllConnectionTypes_Next100.json
func ExampleConnectionTypeClient_NewListByAutomationAccountPager_getConnectionTypesNext100() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectionTypeClient().NewListByAutomationAccountPager("rg", "myAutomationAccount25", nil)
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
		// page.ConnectionTypeListResult = armautomation.ConnectionTypeListResult{
		// 	Value: []*armautomation.ConnectionType{
		// 		{
		// 			Name: to.Ptr("myCT27"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT27"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.87+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.883+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT28"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT28"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.07+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.087+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT29"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT29"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.29+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.307+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT3"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT3"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.477+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.493+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT30"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT30"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.523+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.54+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT31"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT31"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.727+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.743+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT32"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT32"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.93+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.947+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT33"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT33"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.133+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.15+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT34"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT34"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.353+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.37+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT35"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT35"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.57+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.587+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT36"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT36"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.773+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.79+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT37"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT37"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.977+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.993+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT38"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT38"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.197+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.213+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT39"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT39"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.417+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.447+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT4"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT4"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.68+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT40"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT40"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.633+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.65+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT41"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT41"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.853+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.87+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT42"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT42"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.057+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.07+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT43"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT43"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.26+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.273+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT44"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT44"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.477+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.493+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT45"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT45"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.68+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT46"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT46"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.883+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.9+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT47"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT47"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.103+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.12+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT48"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT48"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.307+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.32+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT49"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT49"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.523+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.54+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT5"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT5"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.9+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.917+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT50"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT50"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.727+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.743+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT51"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT51"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.93+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.947+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT52"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT52"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.133+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.167+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT53"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT53"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.353+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.37+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT54"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT54"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.57+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.587+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT55"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT55"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.773+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.79+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT56"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT56"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.977+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.993+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT57"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT57"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.197+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.213+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT58"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT58"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.477+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.493+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT59"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT59"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.68+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT6"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT6"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.133+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.167+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT60"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT60"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.917+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.947+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT61"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT61"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.133+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.15+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT62"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT62"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.337+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.353+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT63"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT63"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.587+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.587+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT64"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT64"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.79+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.807+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT65"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT65"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.993+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.01+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT66"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT66"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.197+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.213+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT67"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT67"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.417+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.463+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT68"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT68"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.65+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT69"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT69"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.883+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.9+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT7"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT7"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.4+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.447+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT70"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT70"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.087+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.12+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT71"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT71"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.32+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.353+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT72"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT72"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.557+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.57+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT73"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT73"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.773+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.79+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT74"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT74"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.98+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.98+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT75"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT75"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.18+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.2+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT76"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT76"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.383+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.4+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT77"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT77"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.587+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.62+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT78"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT78"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.837+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.853+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT79"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT79"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.057+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.057+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT8"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT8"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.633+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.68+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT80"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT80"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.26+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.277+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT81"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT81"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.493+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.51+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT82"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT82"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.73+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.743+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT83"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT83"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.93+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.947+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT84"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT84"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.133+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.197+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT85"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT85"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.4+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.417+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT86"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT86"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.603+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.62+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT87"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT87"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.82+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.837+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT88"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT88"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.027+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.04+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT89"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT89"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.23+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.243+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT9"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT9"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.883+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.917+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT90"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT90"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.43+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.48+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT91"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT91"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.667+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.713+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT92"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT92"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.9+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.947+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT93"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT93"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.133+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.15+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT94"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT94"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.37+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.417+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT95"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT95"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.603+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.62+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT96"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT96"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.823+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.837+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT97"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT97"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.027+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.04+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT98"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT98"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.243+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.29+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT99"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT99"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.493+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.527+00:00"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
