using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualNetworks_List.json
// this example is just showing the usage of "VirtualNetworks_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabResource created on azure
// for more information of creating DevTestLabResource, please refer to the document of DevTestLabResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
ResourceIdentifier devTestLabResourceId = DevTestLabResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName);
DevTestLabResource devTestLab = client.GetDevTestLabResource(devTestLabResourceId);

// get the collection of this DevTestLabVirtualNetworkResource
DevTestLabVirtualNetworkCollection collection = devTestLab.GetDevTestLabVirtualNetworks();

// invoke the operation and iterate over the result
await foreach (DevTestLabVirtualNetworkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DevTestLabVirtualNetworkData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
