package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/listAutomationAccountsBySubscription.json
func ExampleAccountClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountClient().NewListPager(nil)
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
		// page.AccountListResult = armautomation.AccountListResult{
		// 	Value: []*armautomation.Account{
		// 		{
		// 			Name: to.Ptr("JPEDDeployDSC1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/JPEDeploy1/providers/Microsoft.Automation/automationAccounts/JPEDDeployDSC1"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-04-28T23:48:25.143Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.060Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("jpeDemoAutomation1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/jpeDemo1/providers/Microsoft.Automation/automationAccounts/jpeDemoAutomation1"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-04-25T02:04:10.223Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.060Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ASEAutomationAccount1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/ASERG1/providers/Microsoft.Automation/automationAccounts/ASEAutomationAccount1"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-04-12T05:19:19.480Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:15.640Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AAEU2DSCDemo"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/eus2Demo1/providers/Microsoft.Automation/automationAccounts/AAEU2DSCDemo"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-08-04T14:44:02.397Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AAEU2DSCDemo2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/eus2Demo1/providers/Microsoft.Automation/automationAccounts/AAEU2DSCDemo2"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-08-04T15:03:45.977Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AAsnoverDemo1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/aadscdemo1/providers/Microsoft.Automation/automationAccounts/AAsnoverDemo1"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-01-29T02:29:13.180Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("automationaccdelete"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/oaastest/providers/Microsoft.Automation/automationAccounts/automationaccdelete"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-29T20:30:49.970Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mytest1212"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/aadscdemo2/providers/Microsoft.Automation/automationAccounts/mytest1212"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-12T20:25:36.340Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("deleteacc"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/CIDRG/providers/Microsoft.Automation/automationAccounts/deleteacc"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-12T22:13:39.790Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("deleteme"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/test/providers/Microsoft.Automation/automationAccounts/deleteme"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-12T21:56:10.267Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("deleteme3"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/tst/providers/Microsoft.Automation/automationAccounts/deleteme3"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-12T22:00:51.333Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Eus2Account1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/OaaSCSXASVNLMD6CUTP2UKUNHMCSLLJRVOSRAS2HOBKX4B3A3UBNLZWZEA-East-US/providers/Microsoft.Automation/automationAccounts/Eus2Account1"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-19T19:07:43.200Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("EUS2DDeployDSC1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/EUS2Deploy1/providers/Microsoft.Automation/automationAccounts/EUS2DDeployDSC1"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-04-28T23:50:56.160Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("eusAccount2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/OaaSCSXASVNLMD6CUTP2UKUNHMCSLLJRVOSRAS2HOBKX4B3A3UBNLZWZEA-East-US/providers/Microsoft.Automation/automationAccounts/eusAccount2"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-05-19T19:12:19.853Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myProdAutomation1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/myProdAutomation1"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-16T21:31:06.333Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-16T21:31:06.333Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myProdDevAutomation"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/myProdDevAutomation"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-27T21:11:16.710Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myProdPublicAutomation"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/myProdPublicAutomation"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-11-18T19:49:08.893Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myTestaccount"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myTestaccount"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-24T00:47:04.227Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kjohn-sandbox-eus"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/kjohn-sandbox-eus"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T17:29:18.493Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-19T17:29:18.493Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kjohn-sandbox-eus-proddev"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/kjohn-sandbox-eus-proddev"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-06T02:33:10.290Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kjohn-sandbox-eus-prodtest"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/kjohn-rg/providers/Microsoft.Automation/automationAccounts/kjohn-sandbox-eus-prodtest"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-10T23:40:13.103Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kjohn-sandbox-eus-prodtest2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/kjohn-rg/providers/Microsoft.Automation/automationAccounts/kjohn-sandbox-eus-prodtest2"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-11T00:20:50.463Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("LinuxPatchingOpsEUS-AA2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/LinuxPatchingOpsEUS-AA2"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-24T03:17:00.043Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("LinuxTestNewAA"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/LinuxTestNewAA"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-23T18:50:54.887Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("psrdfeAccount2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/OaasCSsubid-east-us/providers/Microsoft.Automation/automationAccounts/psrdfeAccount2"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-05-05T00:26:49.020Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-account-one"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/myrg/providers/Microsoft.Automation/automationAccounts/my-account-one"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-10-27T17:54:31.007Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myku-no-vms"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/myku-no-vms"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-30T01:12:00.853Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myku-win-vms"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/myku-win-vms"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-30T01:17:07.613Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SVCPrnAcctTest1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/aadscdemo2/providers/Microsoft.Automation/automationAccounts/SVCPrnAcctTest1"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-28T20:12:48.163Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testfgbhfghfgh"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/tesdbgfvfhgjghjgh/providers/Microsoft.Automation/automationAccounts/testfgbhfghfgh"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-12T20:48:59.300Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("test-linux-linuxopsworkspace"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-eus/providers/Microsoft.Automation/automationAccounts/test-linux-linuxopsworkspace"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-23T18:59:56.990Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-dsc-test-1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my-dev/providers/Microsoft.Automation/automationAccounts/my-dsc-test-1"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-11-10T00:21:05.133Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-test-automation-1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my-dev/providers/Microsoft.Automation/automationAccounts/my-test-automation-1"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-11-09T19:57:50.043Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAccount"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/myresourcegroupeus/providers/Microsoft.Automation/automationAccounts/myAccount"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T19:10:30.453Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T19:10:30.453Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAccount123"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAccount123"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-29T00:32:32.520Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAccountasfads"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAccountasfads"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:21:03.270Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:21:03.270Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAccountEUS"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/myResourceGroupEUS/providers/Microsoft.Automation/automationAccounts/myAccountEUS"),
		// 			Location: to.Ptr("eastus2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-29T23:13:38.873Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.400Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount1"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:22:33.260Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:22:33.260Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount11"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount11"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:10:24.523Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-26T02:11:12.027Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount2"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:24:20.310Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:24:20.310Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount3"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount3"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:24:43.967Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-24T23:24:43.967Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount4"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount4"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:04:56.900Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:04:56.900Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount6"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount6"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:10:44.567Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:10:44.567Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAutomationAccount7"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount7"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:19:17.943Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-25T02:19:17.943Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CSSCase1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/WEURG1/providers/Microsoft.Automation/automationAccounts/CSSCase1"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-08-18T05:53:58.910Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:15.373Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("deleteme"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my/providers/Microsoft.Automation/automationAccounts/deleteme"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-12T21:48:47.980Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:15.373Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("LinuxPatchingOpsWEU-AA2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/mms-weu/providers/Microsoft.Automation/automationAccounts/LinuxPatchingOpsWEU-AA2"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-24T02:27:35.713Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:15.373Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("LinuxPatchJobs"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my-dev/providers/Microsoft.Automation/automationAccounts/LinuxPatchJobs"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-14T22:02:28.223Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:15.373Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MyCoolAAC1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/IgnniteRG/providers/Microsoft.Automation/automationAccounts/MyCoolAAC1"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-22T05:59:22.443Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:15.373Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MyCoolACT1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/IgnniteRG/providers/Microsoft.Automation/automationAccounts/MyCoolACT1"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-22T06:05:18.500Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:15.373Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MyTestmyTest-WEU"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/OaaSCSXASVNLMD6CUTP2UKUNHMCSLLJRVOSRAS2HOBKX4B3A3UBNLZWZEA-West-Europe/providers/Microsoft.Automation/automationAccounts/MyTestmyTest-WEU"),
		// 			Location: to.Ptr("West Europe"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2014-11-12T02:48:51.473Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:15.373Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("WEDDeployDSC1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/WEDeploy1/providers/Microsoft.Automation/automationAccounts/WEDDeployDSC1"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-04-28T23:20:01.730Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-13T08:43:47.360Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SEADDeployDSC1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/SEADeploy1/providers/Microsoft.Automation/automationAccounts/SEADDeployDSC1"),
		// 			Location: to.Ptr("southeamyia"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-04-28T19:06:39.100Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.683Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("seaDemoAutomation1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/seaDemo1/providers/Microsoft.Automation/automationAccounts/seaDemoAutomation1"),
		// 			Location: to.Ptr("southeamyia"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-04-25T01:31:32.150Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:16.683Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AAsnoverDemo2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/aadscdemo1/providers/Microsoft.Automation/automationAccounts/AAsnoverDemo2"),
		// 			Location: to.Ptr("South Central US"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-01-29T02:30:05.840Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:17.107Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SCUSDDeployDSC1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/SCUSDeploy1/providers/Microsoft.Automation/automationAccounts/SCUSDDeployDSC1"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-06-04T23:06:44.897Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:17.107Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("scusposthydtest1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/eus2Demo1/providers/Microsoft.Automation/automationAccounts/scusposthydtest1"),
		// 			Location: to.Ptr("South Central US"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-06-04T16:44:18.823Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:17.107Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("scustestaccount1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/OaaSCSXASVNLMD6CUTP2UKUNHMCSLLJRVOSRAS2HOBKX4B3A3UBNLZWZEA-South-Central-US/providers/Microsoft.Automation/automationAccounts/scustestaccount1"),
		// 			Location: to.Ptr("South Central US"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-05-30T01:54:57.313Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:17.107Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("AAUKSmyTest1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/UKSRG1/providers/Microsoft.Automation/automationAccounts/AAUKSmyTest1"),
		// 			Location: to.Ptr("uksouth"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-15T02:31:03.190Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-15T02:31:03.190Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Tip-WCUS-AutomationAccount"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/Gaurav_machines/providers/Microsoft.Automation/automationAccounts/Tip-WCUS-AutomationAccount"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-08T19:53:36.293Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-08T19:53:36.293Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("aa-my"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/myRG/providers/Microsoft.Automation/automationAccounts/aa-my"),
		// 			Location: to.Ptr("northeurope"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-09-21T18:59:56.260Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:15.457Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CIDAccout1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/CIDRG/providers/Microsoft.Automation/automationAccounts/CIDAccout1"),
		// 			Location: to.Ptr("Central India"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-02-24T20:04:58.867Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:19.003Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("aaspntest"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my/providers/Microsoft.Automation/automationAccounts/aaspntest"),
		// 			Location: to.Ptr("northcentralus"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-06T17:18:51.880Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:19.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dsccomposite"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my/providers/Microsoft.Automation/automationAccounts/dsccomposite"),
		// 			Location: to.Ptr("northcentralus"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-13T17:37:55.163Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T17:19:15.090Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dsclinux"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my/providers/Microsoft.Automation/automationAccounts/dsclinux"),
		// 			Location: to.Ptr("northcentralus"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T21:50:05.493Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T21:50:05.493Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mydsc"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my/providers/Microsoft.Automation/automationAccounts/mydsc"),
		// 			Location: to.Ptr("northcentralus"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-11-01T17:28:36.197Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-09T21:35:19.370Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("mydsc2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/my/providers/Microsoft.Automation/automationAccounts/mydsc2"),
		// 			Location: to.Ptr("northcentralus"),
		// 			Properties: &armautomation.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-11-01T18:50:06.063Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-11T01:33:13.113Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
