using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimPolicyCreate.json
// this example is just showing the usage of "SimPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MobileNetworkResource created on azure
// for more information of creating MobileNetworkResource, please refer to the document of MobileNetworkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string mobileNetworkName = "testMobileNetwork";
ResourceIdentifier mobileNetworkResourceId = MobileNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, mobileNetworkName);
MobileNetworkResource mobileNetwork = client.GetMobileNetworkResource(mobileNetworkResourceId);

// get the collection of this MobileNetworkSimPolicyResource
MobileNetworkSimPolicyCollection collection = mobileNetwork.GetMobileNetworkSimPolicies();

// invoke the operation
string simPolicyName = "testPolicy";
MobileNetworkSimPolicyData data = new MobileNetworkSimPolicyData(new AzureLocation("eastus"), new Ambr("500 Mbps", "1 Gbps"), new WritableSubResource()
{
    Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
}, new MobileNetworkSliceConfiguration[]
{
new MobileNetworkSliceConfiguration(new WritableSubResource()
{
Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
},new WritableSubResource()
{
Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork"),
},new DataNetworkConfiguration[]
{
new DataNetworkConfiguration(new WritableSubResource()
{
Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/dataNetworks/testdataNetwork"),
},new Ambr("500 Mbps","1 Gbps"),new WritableSubResource[]
{
new WritableSubResource()
{
Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/services/testService"),
}
})
{
FiveQi = 9,
AllocationAndRetentionPriorityLevel = 9,
PreemptionCapability = MobileNetworkPreemptionCapability.NotPreempt,
PreemptionVulnerability = MobileNetworkPreemptionVulnerability.Preemptable,
DefaultSessionType = MobileNetworkPduSessionType.IPv4,
AdditionalAllowedSessionTypes =
{
},
MaximumNumberOfBufferedPackets = 200,
}
})
})
{
    RegistrationTimer = 3240,
};
ArmOperation<MobileNetworkSimPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, simPolicyName, data);
MobileNetworkSimPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MobileNetworkSimPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
