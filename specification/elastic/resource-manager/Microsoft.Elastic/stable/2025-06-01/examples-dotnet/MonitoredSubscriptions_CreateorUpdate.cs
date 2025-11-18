using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Elastic;

// Generated from example definition: 2025-06-01/MonitoredSubscriptions_CreateorUpdate.json
// this example is just showing the usage of "MonitoredSubscriptionProperties_CreateorUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticMonitorResource created on azure
// for more information of creating ElasticMonitorResource, please refer to the document of ElasticMonitorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
ResourceIdentifier elasticMonitorResourceId = ElasticMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
ElasticMonitorResource elasticMonitor = client.GetElasticMonitorResource(elasticMonitorResourceId);

// get the collection of this ElasticMonitoredSubscriptionResource
ElasticMonitoredSubscriptionCollection collection = elasticMonitor.GetElasticMonitoredSubscriptions();

// invoke the operation
string configurationName = "default";
ElasticMonitoredSubscriptionData data = new ElasticMonitoredSubscriptionData();
ArmOperation<ElasticMonitoredSubscriptionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, configurationName, data);
ElasticMonitoredSubscriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticMonitoredSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
