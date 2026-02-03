using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ExpressRouteCrossConnectionUpdateTags.json
// this example is just showing the usage of "ExpressRouteCrossConnections_UpdateTags" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteCrossConnectionResource created on azure
// for more information of creating ExpressRouteCrossConnectionResource, please refer to the document of ExpressRouteCrossConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "CrossConnection-SiliconValley";
string crossConnectionName = "<circuitServiceKey>";
ResourceIdentifier expressRouteCrossConnectionResourceId = ExpressRouteCrossConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, crossConnectionName);
ExpressRouteCrossConnectionResource expressRouteCrossConnection = client.GetExpressRouteCrossConnectionResource(expressRouteCrossConnectionResourceId);

// invoke the operation
NetworkTagsObject crossConnectionParameters = new NetworkTagsObject
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
ExpressRouteCrossConnectionResource result = await expressRouteCrossConnection.UpdateAsync(crossConnectionParameters);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExpressRouteCrossConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
