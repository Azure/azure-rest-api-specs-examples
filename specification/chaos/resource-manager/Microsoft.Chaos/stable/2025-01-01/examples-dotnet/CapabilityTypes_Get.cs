using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Chaos;

// Generated from example definition: 2025-01-01/CapabilityTypes_Get.json
// this example is just showing the usage of "CapabilityType_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ChaosTargetMetadataResource created on azure
// for more information of creating ChaosTargetMetadataResource, please refer to the document of ChaosTargetMetadataResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
AzureLocation location = new AzureLocation("westus2");
string targetTypeName = "Microsoft-VirtualMachine";
ResourceIdentifier chaosTargetMetadataResourceId = ChaosTargetMetadataResource.CreateResourceIdentifier(subscriptionId, location, targetTypeName);
ChaosTargetMetadataResource chaosTargetMetadata = client.GetChaosTargetMetadataResource(chaosTargetMetadataResourceId);

// get the collection of this ChaosCapabilityMetadataResource
ChaosCapabilityMetadataCollection collection = chaosTargetMetadata.GetAllChaosCapabilityMetadata();

// invoke the operation
string capabilityTypeName = "Shutdown-1.0";
NullableResponse<ChaosCapabilityMetadataResource> response = await collection.GetIfExistsAsync(capabilityTypeName);
ChaosCapabilityMetadataResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ChaosCapabilityMetadataData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
