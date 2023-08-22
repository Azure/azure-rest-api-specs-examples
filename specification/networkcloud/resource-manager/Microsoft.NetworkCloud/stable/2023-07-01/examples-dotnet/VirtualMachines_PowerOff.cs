using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/VirtualMachines_PowerOff.json
// this example is just showing the usage of "VirtualMachines_PowerOff" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
VirtualMachinePowerOffContent content = new VirtualMachinePowerOffContent()
{
    SkipShutdown = SkipShutdown.True,
};
ArmOperation<NetworkCloudOperationStatusResult> lro = await networkCloudVirtualMachine.PowerOffAsync(WaitUntil.Completed, content: content);
NetworkCloudOperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
