using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StandbyPool;

// Generated from example definition: 2025-03-01/StandbyVirtualMachinePoolRuntimeViews_ListByStandbyPool.json
// this example is just showing the usage of "StandbyVirtualMachinePoolRuntimeViewResource_ListByStandbyPool" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StandbyVirtualMachinePoolResource created on azure
// for more information of creating StandbyVirtualMachinePoolResource, please refer to the document of StandbyVirtualMachinePoolResource
string subscriptionId = "00000000-0000-0000-0000-000000000009";
string resourceGroupName = "rgstandbypool";
string standbyVirtualMachinePoolName = "pool";
ResourceIdentifier standbyVirtualMachinePoolResourceId = StandbyVirtualMachinePoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, standbyVirtualMachinePoolName);
StandbyVirtualMachinePoolResource standbyVirtualMachinePool = client.GetStandbyVirtualMachinePoolResource(standbyVirtualMachinePoolResourceId);

// get the collection of this StandbyVirtualMachinePoolRuntimeViewResource
StandbyVirtualMachinePoolRuntimeViewCollection collection = standbyVirtualMachinePool.GetStandbyVirtualMachinePoolRuntimeViews();

// invoke the operation and iterate over the result
await foreach (StandbyVirtualMachinePoolRuntimeViewResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    StandbyVirtualMachinePoolRuntimeViewData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
