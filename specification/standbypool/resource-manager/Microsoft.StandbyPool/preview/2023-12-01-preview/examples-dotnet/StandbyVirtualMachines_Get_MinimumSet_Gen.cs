using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StandbyPool;

// Generated from example definition: specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyVirtualMachines_Get_MinimumSet_Gen.json
// this example is just showing the usage of "StandbyVirtualMachines_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StandbyVirtualMachinePoolResource created on azure
// for more information of creating StandbyVirtualMachinePoolResource, please refer to the document of StandbyVirtualMachinePoolResource
string subscriptionId = "8CC31D61-82D7-4B2B-B9DC-6B924DE7D229";
string resourceGroupName = "rgstandbypool";
string standbyVirtualMachinePoolName = "pool";
ResourceIdentifier standbyVirtualMachinePoolResourceId = StandbyVirtualMachinePoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, standbyVirtualMachinePoolName);
StandbyVirtualMachinePoolResource standbyVirtualMachinePool = client.GetStandbyVirtualMachinePoolResource(standbyVirtualMachinePoolResourceId);

// get the collection of this StandbyVirtualMachineResource
StandbyVirtualMachineCollection collection = standbyVirtualMachinePool.GetStandbyVirtualMachines();

// invoke the operation
string standbyVirtualMachineName = "v-m3mbxY7j5yRugj347t4sQ-1_C7sD7-_H-IbUG26s__B8pE_bGqZm";
NullableResponse<StandbyVirtualMachineResource> response = await collection.GetIfExistsAsync(standbyVirtualMachineName);
StandbyVirtualMachineResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    StandbyVirtualMachineData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
