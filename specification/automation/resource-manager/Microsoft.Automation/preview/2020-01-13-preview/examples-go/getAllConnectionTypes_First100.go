package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getAllConnectionTypes_First100.json
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
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2014-03-18T23:02:40.613Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2014-03-18T23:02:40.740Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureClassicCertificate"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/AzureClassicCertificate"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-04-25T23:54:02.650Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-04-25T23:54:03.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureServicePrincipal"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/AzureServicePrincipal"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-16T23:08:41.853Z"); return t}()),
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
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-16T23:08:42.407Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:14.430Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:14.510Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT0"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT0"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:14.837Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:14.853Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT1"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT1"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.040Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.057Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT10"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT10"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.120Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.150Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT100"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT100"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.730Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.760Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT101"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT101"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.963Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:36.993Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT102"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT102"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.197Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.213Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT103"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT103"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.417Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.430Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT104"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT104"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.633Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.650Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT105"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT105"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.837Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:37.883Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT106"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT106"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.073Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.087Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT107"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT107"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.277Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.290Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT108"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT108"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.493Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.510Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT109"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT109"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.697Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.713Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT11"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT11"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.337Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.383Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT110"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT110"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.930Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:38.947Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT111"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT111"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.150Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT112"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT112"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.353Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT113"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT113"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.557Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.573Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT114"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT114"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.777Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.790Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT115"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT115"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.980Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:39.993Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT116"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT116"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.180Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.197Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT117"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT117"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.417Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.430Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT118"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT118"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.633Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.650Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT119"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT119"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.837Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:40.853Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT12"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT12"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.587Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.603Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT120"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT120"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.040Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.057Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT121"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT121"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.260Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.277Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT122"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT122"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.480Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.557Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT123"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT123"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.743Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.760Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT124"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT124"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.947Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:41.963Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT125"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT125"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.180Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.197Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT126"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT126"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.383Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT127"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT127"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.603Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.620Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT128"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT128"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.823Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:42.837Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT129"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT129"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.027Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.040Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT13"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT13"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.790Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.807Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT130"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT130"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.243Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.260Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT131"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT131"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.447Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.463Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT132"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT132"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.650Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.667Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT133"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT133"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.870Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:43.887Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT134"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT134"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.073Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.087Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT135"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT135"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.290Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.307Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT136"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT136"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.493Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.510Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT137"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT137"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.760Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:44.807Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT138"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT138"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.010Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.027Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT139"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT139"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.213Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.230Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT14"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT14"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:17.993Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.010Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT140"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT140"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.433Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.463Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT141"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT141"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.650Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.667Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT142"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT142"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.883Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:45.917Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT143"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT143"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.103Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.133Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT144"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT144"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.340Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.353Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT145"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT145"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.540Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.557Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT146"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT146"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.760Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.777Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT147"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT147"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.963Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:46.980Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT148"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT148"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.197Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.213Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT149"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT149"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.417Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.430Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT15"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT15"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.213Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.227Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT150"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT150"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.637Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.650Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT151"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT151"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.853Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:47.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT152"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT152"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.103Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT153"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT153"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.370Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT154"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT154"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.590Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.667Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT155"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT155"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.887Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:48.930Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT156"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT156"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.120Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.167Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT157"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT157"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.370Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT158"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT158"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.603Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.637Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT159"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT159"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.840Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:49.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT16"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT16"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.417Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.430Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT160"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT160"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.090Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.103Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT161"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT161"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.290Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.307Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT162"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT162"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.493Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.510Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT163"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT163"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.697Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.713Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT164"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT164"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.917Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:50.933Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT165"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT165"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.120Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.137Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT166"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT166"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.340Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.353Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT167"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT167"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.540Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.557Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT168"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT168"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.743Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.760Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT169"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT169"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.980Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:51.980Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT17"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT17"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.620Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.633Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT170"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT170"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.213Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.230Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT171"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT171"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.417Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.433Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT172"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT172"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.637Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.650Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT173"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT173"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.840Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:52.870Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT174"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT174"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:53.057Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:53.073Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT175"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT175"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:53.260Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:53.290Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT18"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT18"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.820Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:18.837Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT19"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT19"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.040Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.057Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT2"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT2"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.260Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:15.273Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT20"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT20"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.243Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.260Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT21"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT21"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.463Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.477Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT22"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT22"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.667Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.680Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT23"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT23"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.883Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:19.900Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT24"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT24"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.103Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.120Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT25"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT25"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.400Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.417Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myCT26"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount25/connectionTypes/myCT26"),
		// 			Properties: &armautomation.ConnectionTypeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.650Z"); return t}()),
		// 				FieldDefinitions: map[string]*armautomation.FieldDefinition{
		// 					"myBoolField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("bool"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringField": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(false),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 					"myStringFieldEncrypted": &armautomation.FieldDefinition{
		// 						Type: to.Ptr("string"),
		// 						IsEncrypted: to.Ptr(true),
		// 						IsOptional: to.Ptr(false),
		// 					},
		// 				},
		// 				IsGlobal: to.Ptr(false),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-27T06:25:20.667Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
