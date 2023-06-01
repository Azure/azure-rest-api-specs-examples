using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceNetworking;

// Generated from example definition: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/FrontendGet.json
// this example is just showing the usage of "FrontendsInterface_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrafficControllerResource created on azure
// for more information of creating TrafficControllerResource, please refer to the document of TrafficControllerResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string trafficControllerName = "tc1";
ResourceIdentifier trafficControllerResourceId = TrafficControllerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, trafficControllerName);
TrafficControllerResource trafficController = client.GetTrafficControllerResource(trafficControllerResourceId);

// get the collection of this FrontendResource
FrontendCollection collection = trafficController.GetFrontends();

// invoke the operation
string frontendName = "fe1";
bool result = await collection.ExistsAsync(frontendName);

Console.WriteLine($"Succeeded: {result}");
