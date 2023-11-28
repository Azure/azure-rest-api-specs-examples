package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getAllConnectionTypes_Next100.json
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
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.870Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.883Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT28"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT28"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.070Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.087Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT29"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT29"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.290Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.307Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT3"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT3"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.477Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.493Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT30"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT30"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.523Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.540Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT31"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT31"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.727Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.743Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT32"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT32"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.930Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:21.947Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT33"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT33"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.133Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.150Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT34"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT34"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.353Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT35"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT35"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.570Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.587Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT36"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT36"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.773Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.790Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT37"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT37"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.977Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:22.993Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT38"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT38"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.197Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.213Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT39"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT39"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.417Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.447Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT4"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT4"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.680Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.697Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT40"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT40"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.633Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.650Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT41"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT41"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.853Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:23.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT42"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT42"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.057Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.070Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT43"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT43"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.260Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.273Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT44"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT44"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.477Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.493Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT45"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT45"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.680Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.697Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT46"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT46"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.883Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:24.900Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT47"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT47"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.103Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.120Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT48"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT48"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.307Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.320Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT49"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT49"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.523Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.540Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT5"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT5"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.900Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.917Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT50"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT50"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.727Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.743Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT51"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT51"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.930Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:25.947Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT52"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT52"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.133Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT53"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT53"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.353Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT54"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT54"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.570Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.587Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT55"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT55"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.773Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.790Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT56"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT56"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.977Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:26.993Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT57"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT57"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.197Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.213Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT58"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT58"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.477Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.493Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT59"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT59"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.680Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.697Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT6"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT6"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.133Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT60"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT60"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.917Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:27.947Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT61"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT61"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.133Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.150Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT62"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT62"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.337Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.353Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT63"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT63"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.587Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.587Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT64"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT64"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.790Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.807Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT65"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT65"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:28.993Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.010Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT66"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT66"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.197Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.213Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT67"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT67"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.417Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.463Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT68"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT68"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.650Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.697Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT69"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT69"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.883Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:29.900Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT7"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT7"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.400Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.447Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT70"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT70"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.087Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.120Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT71"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT71"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.320Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.353Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT72"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT72"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.557Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.570Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT73"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT73"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.773Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.790Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT74"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT74"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.980Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:30.980Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT75"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT75"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.180Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.200Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT76"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT76"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.383Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT77"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT77"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.587Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.620Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT78"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT78"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.837Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:31.853Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT79"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT79"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.057Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.057Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT8"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT8"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.633Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.680Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT80"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT80"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.260Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.277Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT81"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT81"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.493Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.510Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT82"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT82"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.730Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.743Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT83"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT83"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.930Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:32.947Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT84"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT84"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.133Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.197Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT85"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT85"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.400Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.417Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT86"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT86"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.603Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.620Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT87"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT87"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.820Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:33.837Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT88"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT88"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.027Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.040Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT89"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT89"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.230Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.243Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT9"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT9"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.883Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:16.917Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT90"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT90"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.430Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.480Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT91"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT91"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.667Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.713Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT92"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT92"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.900Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:34.947Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT93"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT93"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.133Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.150Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT94"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT94"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.370Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.417Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT95"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT95"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.603Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.620Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT96"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT96"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.823Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:35.837Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT97"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT97"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.027Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.040Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT98"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT98"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.243Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.290Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT99"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT99"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.493Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.527Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
