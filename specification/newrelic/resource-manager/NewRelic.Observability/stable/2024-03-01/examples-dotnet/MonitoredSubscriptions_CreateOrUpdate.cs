using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NewRelicObservability.Models;
using Azure.ResourceManager.NewRelicObservability;

// Generated from example definition: specification/newrelic/resource-manager/NewRelic.Observability/stable/2024-03-01/examples/MonitoredSubscriptions_CreateOrUpdate.json
// this example is just showing the usage of "MonitoredSubscriptions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NewRelicMonitorResource created on azure
// for more information of creating NewRelicMonitorResource, please refer to the document of NewRelicMonitorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
ResourceIdentifier newRelicMonitorResourceId = NewRelicMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
NewRelicMonitorResource newRelicMonitorResource = client.GetNewRelicMonitorResource(newRelicMonitorResourceId);

// get the collection of this NewRelicMonitoredSubscriptionResource
NewRelicMonitoredSubscriptionCollection collection = newRelicMonitorResource.GetNewRelicMonitoredSubscriptions();

// invoke the operation
MonitoredSubscriptionConfigurationName configurationName = MonitoredSubscriptionConfigurationName.Default;
NewRelicMonitoredSubscriptionData data = new NewRelicMonitoredSubscriptionData();
ArmOperation<NewRelicMonitoredSubscriptionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, configurationName, data);
NewRelicMonitoredSubscriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NewRelicMonitoredSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
