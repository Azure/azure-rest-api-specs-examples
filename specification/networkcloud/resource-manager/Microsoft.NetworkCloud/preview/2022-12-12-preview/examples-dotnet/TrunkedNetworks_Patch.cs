using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/TrunkedNetworks_Patch.json
// this example is just showing the usage of "TrunkedNetworks_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrunkedNetworkResource created on azure
// for more information of creating TrunkedNetworkResource, please refer to the document of TrunkedNetworkResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string trunkedNetworkName = "trunkedNetworkName";
ResourceIdentifier trunkedNetworkResourceId = TrunkedNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, trunkedNetworkName);
TrunkedNetworkResource trunkedNetwork = client.GetTrunkedNetworkResource(trunkedNetworkResourceId);

// invoke the operation
TrunkedNetworkPatch patch = new TrunkedNetworkPatch()
{
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2",
    },
};
TrunkedNetworkResource result = await trunkedNetwork.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TrunkedNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
