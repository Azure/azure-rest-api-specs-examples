using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2023-01-01/examples/postTestNotificationsAtActionGroupResourceLevel.json
// this example is just showing the usage of "ActionGroups_CreateNotificationsAtActionGroupResourceLevel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ActionGroupResource created on azure
// for more information of creating ActionGroupResource, please refer to the document of ActionGroupResource
string subscriptionId = "11111111-1111-1111-1111-111111111111";
string resourceGroupName = "TestRgName";
string actionGroupName = "TestAgName";
ResourceIdentifier actionGroupResourceId = ActionGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, actionGroupName);
ActionGroupResource actionGroup = client.GetActionGroupResource(actionGroupResourceId);

// invoke the operation
NotificationContent content = new NotificationContent("budget")
{
    EmailReceivers =
    {
    new MonitorEmailReceiver("John Doe's email","johndoe@email.com")
    {
    UseCommonAlertSchema = false,
    },new MonitorEmailReceiver("Jane Smith's email","janesmith@email.com")
    {
    UseCommonAlertSchema = true,
    }
    },
    SmsReceivers =
    {
    new MonitorSmsReceiver("John Doe's mobile","1","1234567890"),new MonitorSmsReceiver("Jane Smith's mobile","1","0987654321")
    },
    WebhookReceivers =
    {
    new MonitorWebhookReceiver("Sample webhook 1",new Uri("http://www.example.com/webhook1"))
    {
    UseCommonAlertSchema = true,
    },new MonitorWebhookReceiver("Sample webhook 2",new Uri("http://www.example.com/webhook2"))
    {
    UseCommonAlertSchema = true,
    UseAadAuth = true,
    ObjectId = "d3bb868c-fe44-452c-aa26-769a6538c808",
    IdentifierUri = new Uri("http://someidentifier/d7811ba3-7996-4a93-99b6-6b2f3f355f8a"),
    TenantId = Guid.Parse("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
    }
    },
    ItsmReceivers =
    {
    new MonitorItsmReceiver("Sample itsm","5def922a-3ed4-49c1-b9fd-05ec533819a3|55dfd1f8-7e59-4f89-bf56-4c82f5ace23c","a3b9076c-ce8e-434e-85b4-aff10cb3c8f1","{\"PayloadRevision\":0,\"WorkItemType\":\"Incident\",\"UseTemplate\":false,\"WorkItemData\":\"{}\",\"CreateOneWIPerCI\":false}",new AzureLocation("westcentralus"))
    },
    AzureAppPushReceivers =
    {
    new MonitorAzureAppPushReceiver("Sample azureAppPush","johndoe@email.com")
    },
    AutomationRunbookReceivers =
    {
    new MonitorAutomationRunbookReceiver(new ResourceIdentifier("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest"),"Sample runbook",new ResourceIdentifier("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest/webhooks/Alert1510184037084"),false)
    {
    Name = "testRunbook",
    ServiceUri = new Uri("http://test.me"),
    UseCommonAlertSchema = true,
    }
    },
    VoiceReceivers =
    {
    new MonitorVoiceReceiver("Sample voice","1","1234567890")
    },
    LogicAppReceivers =
    {
    new MonitorLogicAppReceiver("Sample logicApp",new ResourceIdentifier("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/LogicApp/providers/Microsoft.Logic/workflows/testLogicApp"),new Uri("https://prod-27.northcentralus.logic.azure.com/workflows/68e572e818e5457ba898763b7db90877/triggers/manual/paths/invoke/azns/test?api-version=2016-10-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=Abpsb72UYJxPPvmDo937uzofupO5r_vIeWEx7KVHo7w"))
    {
    UseCommonAlertSchema = false,
    }
    },
    AzureFunctionReceivers =
    {
    new MonitorAzureFunctionReceiver("Sample azureFunction",new ResourceIdentifier("/subscriptions/5def922a-3ed4-49c1-b9fd-05ec533819a3/resourceGroups/aznsTest/providers/Microsoft.Web/sites/testFunctionApp"),"HttpTriggerCSharp1",new Uri("http://test.me"))
    {
    UseCommonAlertSchema = true,
    }
    },
    ArmRoleReceivers =
    {
    new MonitorArmRoleReceiver("ArmRole-Common","11111111-1111-1111-1111-111111111111")
    {
    UseCommonAlertSchema = true,
    },new MonitorArmRoleReceiver("ArmRole-nonCommon","11111111-1111-1111-1111-111111111111")
    {
    UseCommonAlertSchema = false,
    }
    },
    EventHubReceivers =
    {
    new MonitorEventHubReceiver("Sample eventHub","testEventHubNameSpace","testEventHub","187f412d-1758-44d9-b052-169e2564721d")
    {
    TenantId = Guid.Parse("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
    }
    },
};
ArmOperation<NotificationStatus> lro = await actionGroup.CreateNotificationsAsync(WaitUntil.Completed, content);
NotificationStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
