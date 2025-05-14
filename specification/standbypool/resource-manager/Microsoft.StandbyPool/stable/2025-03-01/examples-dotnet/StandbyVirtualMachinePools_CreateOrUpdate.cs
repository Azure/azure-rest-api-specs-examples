using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StandbyPool.Models;
using Azure.ResourceManager.StandbyPool;

// Generated from example definition: 2025-03-01/StandbyVirtualMachinePools_CreateOrUpdate.json
// this example is just showing the usage of "StandbyVirtualMachinePoolResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000009";
string resourceGroupName = "rgstandbypool";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this StandbyVirtualMachinePoolResource
StandbyVirtualMachinePoolCollection collection = resourceGroupResource.GetStandbyVirtualMachinePools();

// invoke the operation
string standbyVirtualMachinePoolName = "pool";
StandbyVirtualMachinePoolData data = new StandbyVirtualMachinePoolData(new AzureLocation("West US"))
{
    Properties = new StandbyVirtualMachinePoolProperties(StandbyVirtualMachineState.Running)
    {
        ElasticityProfile = new StandbyVirtualMachinePoolElasticityProfile(304L)
        {
            MinReadyCapacity = 300L,
        },
        AttachedVirtualMachineScaleSetId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
    },
    Tags = { },
};
ArmOperation<StandbyVirtualMachinePoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, standbyVirtualMachinePoolName, data);
StandbyVirtualMachinePoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StandbyVirtualMachinePoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
