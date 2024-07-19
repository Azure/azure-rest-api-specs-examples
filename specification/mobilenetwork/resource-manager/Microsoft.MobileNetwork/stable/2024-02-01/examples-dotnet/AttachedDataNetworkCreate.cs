using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/AttachedDataNetworkCreate.json
// this example is just showing the usage of "AttachedDataNetworks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
string attachedDataNetworkName = "TestAttachedDataNetwork";
MobileAttachedDataNetworkData data = new MobileAttachedDataNetworkData(new AzureLocation("eastus"), new MobileNetworkInterfaceProperties()
{
    Name = "N6",
}, new string[]
{
"1.1.1.1"
})
{
    NaptConfiguration = new NaptConfiguration()
    {
        Enabled = NaptState.Enabled,
        PortRange = new MobileNetworkPortRange()
        {
            MinPort = 1024,
            MaxPort = 49999,
        },
        PortReuseHoldTime = new MobileNetworkPortReuseHoldTimes()
        {
            Tcp = 120,
            Udp = 60,
        },
        PinholeLimits = 65536,
        PinholeTimeouts = new PinholeTimeouts()
        {
            Tcp = 180,
            Udp = 30,
            Icmp = 30,
        },
    },
    UserEquipmentAddressPoolPrefix =
    {
    "2.2.0.0/16"
    },
    UserEquipmentStaticAddressPoolPrefix =
    {
    "2.4.0.0/16"
    },
};
ArmOperation<MobileAttachedDataNetworkResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, attachedDataNetworkName, data);
MobileAttachedDataNetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MobileAttachedDataNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
