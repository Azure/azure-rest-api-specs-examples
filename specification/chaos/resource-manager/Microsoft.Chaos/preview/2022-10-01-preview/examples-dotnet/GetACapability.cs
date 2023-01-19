using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Chaos;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-10-01-preview/examples/GetACapability.json
// this example is just showing the usage of "Capabilities_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TargetResource created on azure
// for more information of creating TargetResource, please refer to the document of TargetResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string resourceGroupName = "exampleRG";
string parentProviderNamespace = "Microsoft.Compute";
string parentResourceType = "virtualMachines";
string parentResourceName = "exampleVM";
string targetName = "Microsoft-VirtualMachine";
ResourceIdentifier targetResourceId = TargetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName);
TargetResource target = client.GetTargetResource(targetResourceId);

// get the collection of this CapabilityResource
CapabilityCollection collection = target.GetCapabilities();

// invoke the operation
string capabilityName = "Shutdown-1.0";
bool result = await collection.ExistsAsync(capabilityName);

Console.WriteLine($"Succeeded: {result}");
