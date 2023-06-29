using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkFabrics_Update_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkFabrics_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricResource created on azure
// for more information of creating NetworkFabricResource, please refer to the document of NetworkFabricResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string networkFabricName = "FabricName";
ResourceIdentifier networkFabricResourceId = NetworkFabricResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkFabricName);
NetworkFabricResource networkFabric = client.GetNetworkFabricResource(networkFabricResourceId);

// invoke the operation
NetworkFabricPatch patch = new NetworkFabricPatch()
{
    Tags =
    {
    ["key1758"] = "",
    },
    Annotation = "annotationValue",
    TerminalServerConfiguration = new TerminalServerPatchableProperties()
    {
        Username = "username",
        Password = "xxxxxxx",
        SerialNumber = "234567",
    },
};
ArmOperation<NetworkFabricResource> lro = await networkFabric.UpdateAsync(WaitUntil.Completed, patch);
NetworkFabricResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
