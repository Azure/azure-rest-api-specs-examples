package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getAllConnectionTypes_First100.json
func ExampleConnectionTypeClient_NewListByAutomationAccountPager_getConnectionTypesFirst100() {
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
		// 			Name: to.Ptr("Azure"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/Azure"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2014-03-18T23:02:40.613+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"AutomationCertificateName": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("System.String"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"SubscriptionID": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("System.String"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2014-03-18T23:02:40.74+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureClassicCertificate"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/AzureClassicCertificate"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-04-25T23:54:02.65+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"CertificateAssetName": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("System.String"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"SubscriptionId": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("System.String"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"SubscriptionName": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("System.String"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-04-25T23:54:03.167+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureServicePrincipal"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/AzureServicePrincipal"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-16T23:08:41.853+00:00"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"ApplicationId": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("System.String"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"CertificateThumbprint": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("System.String"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"SubscriptionId": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("System.String"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"TenantId": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("System.String"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(true),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-16T23:08:42.407+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:14.43+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:14.51+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT0"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT0"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:14.837+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:14.853+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT1"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT1"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.04+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.057+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT10"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT10"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.12+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.15+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT100"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT100"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.73+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.76+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT101"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT101"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.963+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.993+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT102"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT102"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.197+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.213+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT103"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT103"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.417+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.43+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT104"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT104"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.633+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.65+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT105"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT105"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.837+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.883+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT106"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT106"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.073+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.087+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT107"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT107"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.277+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.29+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT108"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT108"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.493+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.51+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT109"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT109"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.697+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.713+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT11"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT11"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.337+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.383+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT110"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT110"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.93+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.947+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT111"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT111"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.15+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.167+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT112"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT112"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.353+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.37+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT113"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT113"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.557+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.573+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT114"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT114"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.777+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.79+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT115"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT115"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.98+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.993+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT116"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT116"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.18+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.197+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT117"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT117"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.417+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.43+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT118"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT118"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.633+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.65+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT119"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT119"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.837+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.853+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT12"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT12"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.587+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.603+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT120"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT120"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.04+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.057+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT121"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT121"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.26+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.277+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT122"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT122"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.48+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.557+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT123"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT123"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.743+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.76+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT124"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT124"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.947+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.963+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT125"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT125"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.18+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.197+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT126"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT126"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.383+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.4+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT127"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT127"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.603+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.62+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT128"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT128"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.823+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.837+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT129"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT129"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.027+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.04+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT13"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT13"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.79+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.807+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT130"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT130"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.243+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.26+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT131"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT131"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.447+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.463+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT132"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT132"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.65+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.667+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT133"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT133"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.87+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.887+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT134"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT134"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.073+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.087+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT135"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT135"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.29+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.307+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT136"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT136"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.493+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.51+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT137"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT137"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.76+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.807+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT138"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT138"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.01+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.027+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT139"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT139"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.213+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.23+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT14"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT14"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.993+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.01+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT140"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT140"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.433+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.463+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT141"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT141"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.65+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.667+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT142"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT142"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.883+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.917+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT143"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT143"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.103+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.133+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT144"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT144"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.34+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.353+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT145"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT145"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.54+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.557+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT146"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT146"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.76+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.777+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT147"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT147"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.963+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.98+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT148"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT148"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.197+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.213+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT149"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT149"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.417+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.43+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT15"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT15"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.213+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.227+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT150"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT150"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.637+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.65+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT151"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT151"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.853+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.87+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT152"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT152"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.103+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.167+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT153"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT153"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.37+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.4+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT154"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT154"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.59+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.667+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT155"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT155"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.887+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.93+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT156"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT156"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.12+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.167+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT157"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT157"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.37+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.4+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT158"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT158"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.603+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.637+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT159"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT159"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.84+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.87+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT16"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT16"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.417+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.43+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT160"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT160"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.09+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.103+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT161"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT161"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.29+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.307+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT162"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT162"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.493+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.51+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT163"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT163"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.697+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.713+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT164"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT164"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.917+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.933+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT165"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT165"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.12+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.137+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT166"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT166"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.34+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.353+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT167"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT167"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.54+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.557+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT168"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT168"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.743+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.76+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT169"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT169"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.98+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.98+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT17"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT17"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.62+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.633+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT170"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT170"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.213+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.23+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT171"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT171"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.417+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.433+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT172"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT172"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.637+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.65+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT173"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT173"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.84+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.87+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT174"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT174"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:53.057+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:53.073+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT175"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT175"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:53.26+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:53.29+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT18"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT18"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.82+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.837+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT19"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT19"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.04+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.057+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT2"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT2"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.26+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.273+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT20"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT20"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.243+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.26+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT21"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT21"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.463+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.477+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT22"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT22"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.667+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.68+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT23"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT23"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.883+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.9+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT24"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT24"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.103+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.12+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT25"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT25"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.4+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.417+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT26"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT26"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.65+00:00"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.667+00:00"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
