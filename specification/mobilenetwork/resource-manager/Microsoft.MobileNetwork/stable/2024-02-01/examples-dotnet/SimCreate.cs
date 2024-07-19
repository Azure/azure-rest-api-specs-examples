using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimCreate.json
// this example is just showing the usage of "Sims_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MobileNetworkSimGroupResource created on azure
// for more information of creating MobileNetworkSimGroupResource, please refer to the document of MobileNetworkSimGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string simGroupName = "testSimGroup";
ResourceIdentifier mobileNetworkSimGroupResourceId = MobileNetworkSimGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, simGroupName);
MobileNetworkSimGroupResource mobileNetworkSimGroup = client.GetMobileNetworkSimGroupResource(mobileNetworkSimGroupResourceId);

// get the collection of this MobileNetworkSimResource
MobileNetworkSimCollection collection = mobileNetworkSimGroup.GetMobileNetworkSims();

// invoke the operation
string simName = "testSim";
MobileNetworkSimData data = new MobileNetworkSimData("00000")
{
    IntegratedCircuitCardIdentifier = "8900000000000000000",
    DeviceType = "Video camera",
    SimPolicyId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
    StaticIPConfiguration =
    {
    new SimStaticIPProperties()
    {
    AttachedDataNetworkId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
    SliceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
    StaticIPIPv4Address = "2.4.0.1",
    }
    },
    AuthenticationKey = "00000000000000000000000000000000",
    OperatorKeyCode = "00000000000000000000000000000000",
};
ArmOperation<MobileNetworkSimResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, simName, data);
MobileNetworkSimResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MobileNetworkSimData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
