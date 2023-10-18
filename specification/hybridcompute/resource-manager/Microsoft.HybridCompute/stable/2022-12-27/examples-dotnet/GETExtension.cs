using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-12-27/examples/GETExtension.json
// this example is just showing the usage of "MachineExtensions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputeMachineResource created on azure
// for more information of creating HybridComputeMachineResource, please refer to the document of HybridComputeMachineResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string machineName = "myMachine";
ResourceIdentifier hybridComputeMachineResourceId = HybridComputeMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName);
HybridComputeMachineResource hybridComputeMachine = client.GetHybridComputeMachineResource(hybridComputeMachineResourceId);

// get the collection of this HybridComputeMachineExtensionResource
HybridComputeMachineExtensionCollection collection = hybridComputeMachine.GetHybridComputeMachineExtensions();

// invoke the operation
string extensionName = "CustomScriptExtension";
NullableResponse<HybridComputeMachineExtensionResource> response = await collection.GetIfExistsAsync(extensionName);
HybridComputeMachineExtensionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HybridComputeMachineExtensionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
