using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/machine/Machines_CreateOrUpdate.json
// this example is just showing the usage of "Machines_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HybridComputeMachineResource
HybridComputeMachineCollection collection = resourceGroupResource.GetHybridComputeMachines();

// invoke the operation
string machineName = "myMachine";
HybridComputeMachineData data = new HybridComputeMachineData(new AzureLocation("eastus2euap"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    LocationData = new HybridComputeLocation("Redmond"),
    OSProfile = new HybridComputeOSProfile
    {
        WindowsConfiguration = new HybridComputeWindowsConfiguration
        {
            IsHotpatchingEnabled = true,
        },
    },
    VmId = Guid.Parse("b7a098cc-b0b8-46e8-a205-62f301a62a8f"),
    ClientPublicKey = "string",
    PrivateLinkScopeResourceId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
    ParentClusterResourceId = new ResourceIdentifier("{AzureStackHCIResourceId}"),
};
ArmOperation<HybridComputeMachineResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, machineName, data);
HybridComputeMachineResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridComputeMachineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
