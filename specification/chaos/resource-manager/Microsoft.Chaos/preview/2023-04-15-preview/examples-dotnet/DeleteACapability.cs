using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Chaos;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/preview/2023-04-15-preview/examples/DeleteACapability.json
// this example is just showing the usage of "Capabilities_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CapabilityResource created on azure
// for more information of creating CapabilityResource, please refer to the document of CapabilityResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string resourceGroupName = "exampleRG";
string parentProviderNamespace = "Microsoft.Compute";
string parentResourceType = "virtualMachines";
string parentResourceName = "exampleVM";
string targetName = "Microsoft-VirtualMachine";
string capabilityName = "Shutdown-1.0";
ResourceIdentifier capabilityResourceId = CapabilityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, capabilityName);
CapabilityResource capability = client.GetCapabilityResource(capabilityResourceId);

// invoke the operation
await capability.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
