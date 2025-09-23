using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-10-01/examples/ExpressRouteConnectionGet.json
// this example is just showing the usage of "ExpressRouteConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteConnectionResource created on azure
// for more information of creating ExpressRouteConnectionResource, please refer to the document of ExpressRouteConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "resourceGroupName";
string expressRouteGatewayName = "expressRouteGatewayName";
string connectionName = "connectionName";
ResourceIdentifier expressRouteConnectionResourceId = ExpressRouteConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, expressRouteGatewayName, connectionName);
ExpressRouteConnectionResource expressRouteConnection = client.GetExpressRouteConnectionResource(expressRouteConnectionResourceId);

// invoke the operation
ExpressRouteConnectionResource result = await expressRouteConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExpressRouteConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
