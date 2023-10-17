using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridCompute;
using Azure.ResourceManager.HybridCompute.Models;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-12-27/examples/DELETEExtension.json
// this example is just showing the usage of "MachineExtensions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputeMachineExtensionResource created on azure
// for more information of creating HybridComputeMachineExtensionResource, please refer to the document of HybridComputeMachineExtensionResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string machineName = "myMachine";
string extensionName = "MMA";
ResourceIdentifier hybridComputeMachineExtensionResourceId = HybridComputeMachineExtensionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName, extensionName);
HybridComputeMachineExtensionResource hybridComputeMachineExtension = client.GetHybridComputeMachineExtensionResource(hybridComputeMachineExtensionResourceId);

// invoke the operation
await hybridComputeMachineExtension.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
