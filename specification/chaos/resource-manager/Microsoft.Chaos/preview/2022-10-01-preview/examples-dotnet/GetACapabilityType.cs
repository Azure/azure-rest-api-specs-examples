using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Chaos;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/preview/2022-10-01-preview/examples/GetACapabilityType.json
// this example is just showing the usage of "CapabilityTypes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TargetTypeResource created on azure
// for more information of creating TargetTypeResource, please refer to the document of TargetTypeResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string locationName = "westus2";
string targetTypeName = "Microsoft-VirtualMachine";
ResourceIdentifier targetTypeResourceId = TargetTypeResource.CreateResourceIdentifier(subscriptionId, locationName, targetTypeName);
TargetTypeResource targetType = client.GetTargetTypeResource(targetTypeResourceId);

// get the collection of this CapabilityTypeResource
CapabilityTypeCollection collection = targetType.GetCapabilityTypes();

// invoke the operation
string capabilityTypeName = "Shutdown-1.0";
bool result = await collection.ExistsAsync(capabilityTypeName);

Console.WriteLine($"Succeeded: {result}");
