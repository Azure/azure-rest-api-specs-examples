using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimBulkUploadEncrypted.json
// this example is just showing the usage of "Sims_BulkUploadEncrypted" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
EncryptedSimUploadList encryptedSimUploadList = new EncryptedSimUploadList(1, 1, "ABC123", "ABC123", "ABC123", new SimNameAndEncryptedProperties[]
{
new SimNameAndEncryptedProperties("testSim","00000")
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
EncryptedCredentials = "ABC123",
},new SimNameAndEncryptedProperties("testSim2","00000")
{
IntegratedCircuitCardIdentifier = "8900000000000000001",
DeviceType = "Video camera",
SimPolicyId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
StaticIPConfiguration =
{
new SimStaticIPProperties()
{
AttachedDataNetworkId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
SliceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
StaticIPIPv4Address = "2.4.0.2",
}
},
EncryptedCredentials = "ABC123",
}
});
ArmOperation<AsyncOperationStatus> lro = await mobileNetworkSimGroup.BulkUploadEncryptedSimAsync(WaitUntil.Completed, encryptedSimUploadList);
AsyncOperationStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
