using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Chaos;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/DeleteCapability.json
// this example is just showing the usage of "Capabilities_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ChaosCapabilityResource created on azure
// for more information of creating ChaosCapabilityResource, please refer to the document of ChaosCapabilityResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string resourceGroupName = "exampleRG";
string parentProviderNamespace = "Microsoft.Compute";
string parentResourceType = "virtualMachines";
string parentResourceName = "exampleVM";
string targetName = "Microsoft-VirtualMachine";
string capabilityName = "Shutdown-1.0";
ResourceIdentifier chaosCapabilityResourceId = ChaosCapabilityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, capabilityName);
ChaosCapabilityResource chaosCapability = client.GetChaosCapabilityResource(chaosCapabilityResourceId);

// invoke the operation
await chaosCapability.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
