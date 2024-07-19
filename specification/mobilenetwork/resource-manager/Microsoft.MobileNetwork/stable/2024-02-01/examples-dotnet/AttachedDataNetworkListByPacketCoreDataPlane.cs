using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/AttachedDataNetworkListByPacketCoreDataPlane.json
// this example is just showing the usage of "AttachedDataNetworks_ListByPacketCoreDataPlane" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PacketCoreDataPlaneResource created on azure
// for more information of creating PacketCoreDataPlaneResource, please refer to the document of PacketCoreDataPlaneResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string packetCoreControlPlaneName = "TestPacketCoreCP";
string packetCoreDataPlaneName = "TestPacketCoreDP";
ResourceIdentifier packetCoreDataPlaneResourceId = PacketCoreDataPlaneResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName);
PacketCoreDataPlaneResource packetCoreDataPlane = client.GetPacketCoreDataPlaneResource(packetCoreDataPlaneResourceId);

// get the collection of this MobileAttachedDataNetworkResource
MobileAttachedDataNetworkCollection collection = packetCoreDataPlane.GetMobileAttachedDataNetworks();

// invoke the operation and iterate over the result
await foreach (MobileAttachedDataNetworkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MobileAttachedDataNetworkData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
