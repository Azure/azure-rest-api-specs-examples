using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/AttachedDataNetworkUpdateTags.json
// this example is just showing the usage of "AttachedDataNetworks_UpdateTags" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MobileAttachedDataNetworkResource created on azure
// for more information of creating MobileAttachedDataNetworkResource, please refer to the document of MobileAttachedDataNetworkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string packetCoreControlPlaneName = "TestPacketCoreCP";
string packetCoreDataPlaneName = "TestPacketCoreDP";
string attachedDataNetworkName = "TestAttachedDataNetwork";
ResourceIdentifier mobileAttachedDataNetworkResourceId = MobileAttachedDataNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, packetCoreControlPlaneName, packetCoreDataPlaneName, attachedDataNetworkName);
MobileAttachedDataNetworkResource mobileAttachedDataNetwork = client.GetMobileAttachedDataNetworkResource(mobileAttachedDataNetworkResourceId);

// invoke the operation
MobileNetworkTagsPatch patch = new MobileNetworkTagsPatch()
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    },
};
MobileAttachedDataNetworkResource result = await mobileAttachedDataNetwork.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MobileAttachedDataNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
