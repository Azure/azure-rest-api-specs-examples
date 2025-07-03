using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Avs.Models;
using Azure.ResourceManager.Avs;

// Generated from example definition: 2024-09-01/VirtualMachines_RestrictMovement.json
// this example is just showing the usage of "VirtualMachines_RestrictMovement" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvsPrivateCloudClusterVirtualMachineResource created on azure
// for more information of creating AvsPrivateCloudClusterVirtualMachineResource, please refer to the document of AvsPrivateCloudClusterVirtualMachineResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string clusterName = "cluster1";
string virtualMachineId = "vm-209";
ResourceIdentifier avsPrivateCloudClusterVirtualMachineResourceId = AvsPrivateCloudClusterVirtualMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, clusterName, virtualMachineId);
AvsPrivateCloudClusterVirtualMachineResource avsPrivateCloudClusterVirtualMachine = client.GetAvsPrivateCloudClusterVirtualMachineResource(avsPrivateCloudClusterVirtualMachineResourceId);

// invoke the operation
AvsPrivateCloudClusterVirtualMachineRestrictMovement restrictMovement = new AvsPrivateCloudClusterVirtualMachineRestrictMovement
{
    RestrictMovement = VirtualMachineRestrictMovementState.Enabled,
};
await avsPrivateCloudClusterVirtualMachine.RestrictMovementAsync(WaitUntil.Completed, restrictMovement);

Console.WriteLine("Succeeded");
