using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/Consoles_Get.json
// this example is just showing the usage of "Consoles_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudVirtualMachineResource created on azure
// for more information of creating NetworkCloudVirtualMachineResource, please refer to the document of NetworkCloudVirtualMachineResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string virtualMachineName = "virtualMachineName";
ResourceIdentifier networkCloudVirtualMachineResourceId = NetworkCloudVirtualMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineName);
NetworkCloudVirtualMachineResource networkCloudVirtualMachine = client.GetNetworkCloudVirtualMachineResource(networkCloudVirtualMachineResourceId);

// get the collection of this NetworkCloudVirtualMachineConsoleResource
NetworkCloudVirtualMachineConsoleCollection collection = networkCloudVirtualMachine.GetNetworkCloudVirtualMachineConsoles();

// invoke the operation
string consoleName = "default";
NullableResponse<NetworkCloudVirtualMachineConsoleResource> response = await collection.GetIfExistsAsync(consoleName);
NetworkCloudVirtualMachineConsoleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkCloudVirtualMachineConsoleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
