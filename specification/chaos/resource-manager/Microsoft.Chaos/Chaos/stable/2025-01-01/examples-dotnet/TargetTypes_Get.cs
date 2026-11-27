using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Chaos;

// Generated from example definition: 2025-01-01/TargetTypes_Get.json
// this example is just showing the usage of "TargetType_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this ChaosTargetMetadataResource
AzureLocation location = new AzureLocation("westus2");
ChaosTargetMetadataCollection collection = subscriptionResource.GetAllChaosTargetMetadata(location);

// invoke the operation
string targetTypeName = "Microsoft-Agent";
NullableResponse<ChaosTargetMetadataResource> response = await collection.GetIfExistsAsync(targetTypeName);
ChaosTargetMetadataResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ChaosTargetMetadataData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
