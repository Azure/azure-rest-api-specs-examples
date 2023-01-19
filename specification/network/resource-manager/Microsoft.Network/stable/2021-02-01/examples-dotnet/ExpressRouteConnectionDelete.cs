using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2021-02-01/examples/ExpressRouteConnectionDelete.json
// this example is just showing the usage of "ExpressRouteConnections_Delete" operation, for the dependent resources, they will have to be created separately.

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
await expressRouteConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
