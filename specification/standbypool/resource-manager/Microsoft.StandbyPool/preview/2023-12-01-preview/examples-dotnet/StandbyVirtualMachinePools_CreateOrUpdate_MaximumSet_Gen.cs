using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StandbyPool.Models;
using Azure.ResourceManager.StandbyPool;

// Generated from example definition: specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyVirtualMachinePools_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "StandbyVirtualMachinePools_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "8CC31D61-82D7-4B2B-B9DC-6B924DE7D229";
string resourceGroupName = "rgstandbypool";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this StandbyVirtualMachinePoolResource
StandbyVirtualMachinePoolCollection collection = resourceGroupResource.GetStandbyVirtualMachinePools();

// invoke the operation
string standbyVirtualMachinePoolName = "pool";
StandbyVirtualMachinePoolData data = new StandbyVirtualMachinePoolData(new AzureLocation("West US"))
{
    ElasticityMaxReadyCapacity = 304,
    VirtualMachineState = StandbyVirtualMachineState.Running,
    AttachedVirtualMachineScaleSetId = new ResourceIdentifier("/subscriptions/8CC31D61-82D7-4B2B-B9DC-6B924DE7D229/resourceGroups/vmssRg/providers/Microsoft.Compute/virtualMachineScaleSets/myVmss"),
    Tags =
    {
    },
};
ArmOperation<StandbyVirtualMachinePoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, standbyVirtualMachinePoolName, data);
StandbyVirtualMachinePoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StandbyVirtualMachinePoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
