using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/Consoles_Patch.json
// this example is just showing the usage of "Consoles_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudVirtualMachineConsoleResource created on azure
// for more information of creating NetworkCloudVirtualMachineConsoleResource, please refer to the document of NetworkCloudVirtualMachineConsoleResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string virtualMachineName = "virtualMachineName";
string consoleName = "default";
ResourceIdentifier networkCloudVirtualMachineConsoleResourceId = NetworkCloudVirtualMachineConsoleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineName, consoleName);
NetworkCloudVirtualMachineConsoleResource networkCloudVirtualMachineConsole = client.GetNetworkCloudVirtualMachineConsoleResource(networkCloudVirtualMachineConsoleResourceId);

// invoke the operation
NetworkCloudVirtualMachineConsolePatch patch = new NetworkCloudVirtualMachineConsolePatch
{
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2"
    },
    Enabled = ConsoleEnabled.True,
    ExpireOn = DateTimeOffset.Parse("2022-06-01T01:27:03.008Z"),
    KeyData = "ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm",
};
ArmOperation<NetworkCloudVirtualMachineConsoleResource> lro = await networkCloudVirtualMachineConsole.UpdateAsync(WaitUntil.Completed, patch);
NetworkCloudVirtualMachineConsoleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudVirtualMachineConsoleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
