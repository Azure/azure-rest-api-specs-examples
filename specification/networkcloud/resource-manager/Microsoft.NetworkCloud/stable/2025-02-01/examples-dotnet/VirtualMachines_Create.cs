using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/VirtualMachines_Create.json
// this example is just showing the usage of "VirtualMachines_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetworkCloudVirtualMachineResource
NetworkCloudVirtualMachineCollection collection = resourceGroupResource.GetNetworkCloudVirtualMachines();

// invoke the operation
string virtualMachineName = "virtualMachineName";
NetworkCloudVirtualMachineData data = new NetworkCloudVirtualMachineData(
    new AzureLocation("location"),
    new ExtendedLocation("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName", "CustomLocation"),
    "username",
    new NetworkAttachment(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName"), VirtualMachineIPAllocationMethod.Dynamic),
    2L,
    8L,
    new NetworkCloudStorageProfile(new NetworkCloudOSDisk(120L)
    {
        CreateOption = OSDiskCreateOption.Ephemeral,
        DeleteOption = OSDiskDeleteOption.Delete,
    })
    {
        VolumeAttachments = { new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/volumes/volumeName") },
    },
    "myacr.azurecr.io/foobar:latest")
{
    BootMethod = VirtualMachineBootMethod.Uefi,
    NetworkAttachments = {new NetworkAttachment(new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"), VirtualMachineIPAllocationMethod.Dynamic)
    {
    DefaultGateway = DefaultGateway.True,
    IPv4Address = "198.51.100.1",
    IPv6Address = "2001:0db8:0000:0000:0000:0000:0000:0000",
    NetworkAttachmentName = "netAttachName01",
    }},
    NetworkData = "bmV0d29ya0RhdGVTYW1wbGU=",
    PlacementHints = { new VirtualMachinePlacementHint(VirtualMachinePlacementHintType.Affinity, new ResourceIdentifier("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"), VirtualMachineSchedulingExecution.Hard, new VirtualMachinePlacementHintPodAffinityScope("")) },
    SshPublicKeys = { new NetworkCloudSshPublicKey("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm") },
    UserData = "dXNlckRhdGVTYW1wbGU=",
    VmDeviceModel = VirtualMachineDeviceModelType.T2,
    VmImageRepositoryCredentials = new ImageRepositoryCredentials("myacr.azurecr.io", "myuser")
    {
        Password = "{password}",
    },
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2"
    },
};
ArmOperation<NetworkCloudVirtualMachineResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, virtualMachineName, data);
NetworkCloudVirtualMachineResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudVirtualMachineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
