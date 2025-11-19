using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Elastic;

// Generated from example definition: 2025-06-01/MonitoredSubscriptions_Delete.json
// this example is just showing the usage of "MonitoredSubscriptionProperties_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticMonitoredSubscriptionResource created on azure
// for more information of creating ElasticMonitoredSubscriptionResource, please refer to the document of ElasticMonitoredSubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
string configurationName = "default";
ResourceIdentifier elasticMonitoredSubscriptionResourceId = ElasticMonitoredSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, configurationName);
ElasticMonitoredSubscriptionResource elasticMonitoredSubscription = client.GetElasticMonitoredSubscriptionResource(elasticMonitoredSubscriptionResourceId);

// invoke the operation
await elasticMonitoredSubscription.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
