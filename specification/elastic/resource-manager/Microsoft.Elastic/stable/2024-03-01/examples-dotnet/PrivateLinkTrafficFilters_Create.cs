using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Elastic.Models;
using Azure.ResourceManager.Elastic;

// Generated from example definition: specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/PrivateLinkTrafficFilters_Create.json
// this example is just showing the usage of "createAndAssociatePLFilter_Create" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
string privateEndpointGuid = "fdb54d3b-e85e-4d08-8958-0d2f7g523df9";
string privateEndpointName = "myPrivateEndpoint";
await elasticMonitor.CreateAndAssociatePrivateLinkFilterAsync(WaitUntil.Completed, privateEndpointGuid: privateEndpointGuid, privateEndpointName: privateEndpointName);

Console.WriteLine("Succeeded");
