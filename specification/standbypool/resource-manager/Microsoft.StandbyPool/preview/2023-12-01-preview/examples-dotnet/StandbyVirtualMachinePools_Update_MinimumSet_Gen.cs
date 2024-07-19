using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StandbyPool.Models;
using Azure.ResourceManager.StandbyPool;

// Generated from example definition: specification/standbypool/resource-manager/Microsoft.StandbyPool/preview/2023-12-01-preview/examples/StandbyVirtualMachinePools_Update_MinimumSet_Gen.json
// this example is just showing the usage of "StandbyVirtualMachinePools_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
StandbyVirtualMachinePoolPatch patch = new StandbyVirtualMachinePoolPatch();
StandbyVirtualMachinePoolResource result = await standbyVirtualMachinePool.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StandbyVirtualMachinePoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
