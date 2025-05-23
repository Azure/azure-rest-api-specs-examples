using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NewRelicObservability.Models;
using Azure.ResourceManager.NewRelicObservability;

// Generated from example definition: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-03-01/examples/MonitoredSubscriptions_Delete.json
// this example is just showing the usage of "MonitoredSubscriptions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NewRelicMonitoredSubscriptionResource created on azure
// for more information of creating NewRelicMonitoredSubscriptionResource, please refer to the document of NewRelicMonitoredSubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
MonitoredSubscriptionConfigurationName configurationName = MonitoredSubscriptionConfigurationName.Default;
ResourceIdentifier newRelicMonitoredSubscriptionResourceId = NewRelicMonitoredSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, configurationName);
NewRelicMonitoredSubscriptionResource newRelicMonitoredSubscription = client.GetNewRelicMonitoredSubscriptionResource(newRelicMonitoredSubscriptionResourceId);

// invoke the operation
await newRelicMonitoredSubscription.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
