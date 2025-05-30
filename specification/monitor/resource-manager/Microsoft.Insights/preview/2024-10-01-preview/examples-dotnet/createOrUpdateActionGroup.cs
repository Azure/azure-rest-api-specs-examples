using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2024-10-01-preview/examples/createOrUpdateActionGroup.json
// this example is just showing the usage of "ActionGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
string resourceGroupName = "Default-NotificationRules";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ActionGroupResource
ActionGroupCollection collection = resourceGroupResource.GetActionGroups();

// invoke the operation
string actionGroupName = "SampleActionGroup";
ActionGroupData data = new ActionGroupData(new AzureLocation("Global"))
{
    GroupShortName = "sample",
    IsEnabled = true,
    EmailReceivers = {new MonitorEmailReceiver("John Doe's email", "johndoe@email.com")
    {
    UseCommonAlertSchema = false,
    }, new MonitorEmailReceiver("Jane Smith's email", "janesmith@email.com")
    {
    UseCommonAlertSchema = true,
    }},
    SmsReceivers = { new MonitorSmsReceiver("John Doe's mobile", "1", "1234567890"), new MonitorSmsReceiver("Jane Smith's mobile", "1", "0987654321") },
    WebhookReceivers = {new MonitorWebhookReceiver("Sample webhook 1", new Uri("http://www.example.com/webhook1"))
    {
    UseCommonAlertSchema = true,
    }, new MonitorWebhookReceiver("Sample webhook 2", new Uri("http://www.example.com/webhook2"))
    {
    UseCommonAlertSchema = true,
    UseAadAuth = true,
    ObjectId = "d3bb868c-fe44-452c-aa26-769a6538c808",
    IdentifierUri = new Uri("http://someidentifier/d7811ba3-7996-4a93-99b6-6b2f3f355f8a"),
    TenantId = Guid.Parse("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
    ManagedIdentity = "30fe7a91-cd31-4edf-96ab-52883b3199cd",
    }},
    ItsmReceivers = { new MonitorItsmReceiver("Sample itsm", "5def922a-3ed4-49c1-b9fd-05ec533819a3|55dfd1f8-7e59-4f89-bf56-4c82f5ace23c", "a3b9076c-ce8e-434e-85b4-aff10cb3c8f1", "{\"PayloadRevision\":0,\"WorkItemType\":\"Incident\",\"UseTemplate\":false,\"WorkItemData\":\"{}\",\"CreateOneWIPerCI\":false}", new AzureLocation("westcentralus")) },
    AzureAppPushReceivers = { new MonitorAzureAppPushReceiver("Sample azureAppPush", "johndoe@email.com") },
    AutomationRunbookReceivers = {new MonitorAutomationRunbookReceiver(new ResourceIdentifier("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest"), "Sample runbook", new ResourceIdentifier("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest/webhooks/Alert1510184037084"), false)
    {
    Name = "testRunbook",
    ServiceUri = new Uri("<serviceUri>"),
    UseCommonAlertSchema = true,
    ManagedIdentity = "30fe7a91-cd31-4edf-96ab-52883b3199cd",
    }},
    VoiceReceivers = { new MonitorVoiceReceiver("Sample voice", "1", "1234567890") },
    LogicAppReceivers = {new MonitorLogicAppReceiver("Sample logicApp", new ResourceIdentifier("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/LogicApp/providers/Microsoft.Logic/workflows/testLogicApp"), new Uri("https://prod-27.northcentralus.logic.azure.com/workflows/68e572e818e5457ba898763b7db90877/triggers/manual/paths/invoke/azns/test?api-version=2016-10-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=Abpsb72UYJxPPvmDo937uzofupO5r_vIeWEx7KVHo7w"))
    {
    UseCommonAlertSchema = false,
    ManagedIdentity = "30fe7a91-cd31-4edf-96ab-52883b3199cd",
    }},
    AzureFunctionReceivers = {new MonitorAzureFunctionReceiver("Sample azureFunction", new ResourceIdentifier("/subscriptions/5def922a-3ed4-49c1-b9fd-05ec533819a3/resourceGroups/aznsTest/providers/Microsoft.Web/sites/testFunctionApp"), "HttpTriggerCSharp1", new Uri("http://test.me"))
    {
    UseCommonAlertSchema = true,
    ManagedIdentity = "30fe7a91-cd31-4edf-96ab-52883b3199cd",
    }},
    ArmRoleReceivers = {new MonitorArmRoleReceiver("Sample armRole", "8e3af657-a8ff-443c-a75c-2fe8c4bcb635")
    {
    UseCommonAlertSchema = true,
    }},
    EventHubReceivers = {new MonitorEventHubReceiver("Sample eventHub", "testEventHubNameSpace", "testEventHub", "187f412d-1758-44d9-b052-169e2564721d")
    {
    TenantId = Guid.Parse("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
    ManagedIdentity = "30fe7a91-cd31-4edf-96ab-52883b3199cd",
    }},
    IncidentReceivers = {new IncidentReceiver("IncidentAction", new IncidentServiceConnection("IncidentConnection", "8be638e7-1419-42d4-a059-437a5f4f4e4e"), IncidentManagementService.Icm, new Dictionary<string, string>
    {
    ["icm.automitigationenabled"] = "true",
    ["icm.correlationid"] = "${data.essentials.signalType}://${data.essentials.originAlertId}",
    ["icm.monitorid"] = "${data.essentials.alertRule}",
    ["icm.occurringlocation.environment"] = "PROD",
    ["icm.routingid"] = "${data.essentials.monitoringService}://${data.essentials.signalType}",
    ["icm.title"] = "${data.essentials.severity}:${data.essentials.monitorCondition} ${data.essentials.monitoringService}:${data.essentials.signalType} ${data.essentials.alertTargetIds}",
    ["icm.tsgid"] = "https://microsoft.com"
    })},
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/AzSecPackAutoConfigRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ThomasTestManagedIdentity_123")] = new UserAssignedIdentity(),
        [new ResourceIdentifier("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/AzSecPackAutoConfigRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ThomasTestManagedIdentity_456")] = new UserAssignedIdentity()
        },
    },
    Tags = { },
};
ArmOperation<ActionGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, actionGroupName, data);
ActionGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ActionGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
