using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/L3Networks_Patch.json
// this example is just showing the usage of "L3Networks_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this L3NetworkResource created on azure
// for more information of creating L3NetworkResource, please refer to the document of L3NetworkResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string l3NetworkName = "l3NetworkName";
ResourceIdentifier l3NetworkResourceId = L3NetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3NetworkName);
L3NetworkResource l3Network = client.GetL3NetworkResource(l3NetworkResourceId);

// invoke the operation
L3NetworkPatch patch = new L3NetworkPatch()
{
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2",
    },
};
L3NetworkResource result = await l3Network.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
L3NetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
