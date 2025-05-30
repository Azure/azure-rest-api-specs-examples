using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Chaos;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/GetCapabilityType.json
// this example is just showing the usage of "CapabilityTypes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ChaosCapabilityTypeResource created on azure
// for more information of creating ChaosCapabilityTypeResource, please refer to the document of ChaosCapabilityTypeResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
string locationName = "westus2";
string targetTypeName = "Microsoft-VirtualMachine";
string capabilityTypeName = "Shutdown-1.0";
ResourceIdentifier chaosCapabilityTypeResourceId = ChaosCapabilityTypeResource.CreateResourceIdentifier(subscriptionId, locationName, targetTypeName, capabilityTypeName);
ChaosCapabilityTypeResource chaosCapabilityType = client.GetChaosCapabilityTypeResource(chaosCapabilityTypeResourceId);

// invoke the operation
ChaosCapabilityTypeResource result = await chaosCapabilityType.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ChaosCapabilityTypeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
