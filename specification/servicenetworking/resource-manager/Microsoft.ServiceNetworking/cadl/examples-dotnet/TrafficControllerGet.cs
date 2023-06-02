using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceNetworking;

// Generated from example definition: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/TrafficControllerGet.json
// this example is just showing the usage of "TrafficControllerInterface_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this TrafficControllerResource
TrafficControllerCollection collection = resourceGroupResource.GetTrafficControllers();

// invoke the operation
string trafficControllerName = "tc1";
bool result = await collection.ExistsAsync(trafficControllerName);

Console.WriteLine($"Succeeded: {result}");
