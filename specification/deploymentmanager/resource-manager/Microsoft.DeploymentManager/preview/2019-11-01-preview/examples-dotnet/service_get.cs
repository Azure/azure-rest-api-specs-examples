using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeploymentManager;

// Generated from example definition: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/service_get.json
// this example is just showing the usage of "Services_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceTopologyResource created on azure
// for more information of creating ServiceTopologyResource, please refer to the document of ServiceTopologyResource
string subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
string resourceGroupName = "myResourceGroup";
string serviceTopologyName = "myTopology";
ResourceIdentifier serviceTopologyResourceId = ServiceTopologyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceTopologyName);
ServiceTopologyResource serviceTopologyResource = client.GetServiceTopologyResource(serviceTopologyResourceId);

// get the collection of this ServiceResource
ServiceResourceCollection collection = serviceTopologyResource.GetServiceResources();

// invoke the operation
string serviceName = "myService";
bool result = await collection.ExistsAsync(serviceName);

Console.WriteLine($"Succeeded: {result}");
