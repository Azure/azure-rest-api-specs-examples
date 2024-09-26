using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Elastic;

// Generated from example definition: specification/elastic/resource-manager/Microsoft.Elastic/preview/2024-06-15-preview/examples/MonitoredSubscriptions_Delete.json
// this example is just showing the usage of "MonitoredSubscriptions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MonitoredSubscriptionPropertyResource created on azure
// for more information of creating MonitoredSubscriptionPropertyResource, please refer to the document of MonitoredSubscriptionPropertyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
string configurationName = "default";
ResourceIdentifier monitoredSubscriptionPropertyResourceId = MonitoredSubscriptionPropertyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, configurationName);
MonitoredSubscriptionPropertyResource monitoredSubscriptionProperty = client.GetMonitoredSubscriptionPropertyResource(monitoredSubscriptionPropertyResourceId);

// invoke the operation
await monitoredSubscriptionProperty.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
