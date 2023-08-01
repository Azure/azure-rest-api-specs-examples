using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/ExpressRoutePortDelete.json
// this example is just showing the usage of "ExpressRoutePorts_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRoutePortResource created on azure
// for more information of creating ExpressRoutePortResource, please refer to the document of ExpressRoutePortResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string expressRoutePortName = "portName";
ResourceIdentifier expressRoutePortResourceId = ExpressRoutePortResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, expressRoutePortName);
ExpressRoutePortResource expressRoutePort = client.GetExpressRoutePortResource(expressRoutePortResourceId);

// invoke the operation
await expressRoutePort.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
