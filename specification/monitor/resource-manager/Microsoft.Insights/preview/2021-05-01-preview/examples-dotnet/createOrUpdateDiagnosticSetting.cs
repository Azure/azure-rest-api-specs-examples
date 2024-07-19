using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/examples/createOrUpdateDiagnosticSetting.json
// this example is just showing the usage of "DiagnosticSettings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this DiagnosticSettingResource
string resourceUri = "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", resourceUri));
DiagnosticSettingCollection collection = client.GetDiagnosticSettings(scopeId);

// invoke the operation
string name = "mysetting";
DiagnosticSettingData data = new DiagnosticSettingData()
{
    StorageAccountId = new ResourceIdentifier("/subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/apptest/providers/Microsoft.Storage/storageAccounts/appteststorage1"),
    EventHubAuthorizationRuleId = new ResourceIdentifier("/subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourceGroups/montest/providers/microsoft.eventhub/namespaces/mynamespace/authorizationrules/myrule"),
    EventHubName = "myeventhub",
    Metrics =
    {
    new MetricSettings(true)
    {
    Category = "WorkflowMetrics",
    RetentionPolicy = new RetentionPolicy(false,0),
    }
    },
    Logs =
    {
    new LogSettings(true)
    {
    CategoryGroup = "allLogs",
    RetentionPolicy = new RetentionPolicy(false,0),
    }
    },
    WorkspaceId = new ResourceIdentifier(""),
    MarketplacePartnerId = new ResourceIdentifier("/subscriptions/abcdeabc-1234-1234-ab12-123a1234567a/resourceGroups/test-rg/providers/Microsoft.Datadog/monitors/dd1"),
    LogAnalyticsDestinationType = "Dedicated",
};
ArmOperation<DiagnosticSettingResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
DiagnosticSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiagnosticSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
