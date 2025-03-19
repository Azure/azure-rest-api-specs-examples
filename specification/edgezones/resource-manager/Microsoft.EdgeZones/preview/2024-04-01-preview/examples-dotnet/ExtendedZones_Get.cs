using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.EdgeZones;

// Generated from example definition: specification/edgezones/resource-manager/Microsoft.EdgeZones/preview/2024-04-01-preview/examples/ExtendedZones_Get.json
// this example is just showing the usage of "ExtendedZones_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "a1ffc958-d2c7-493e-9f1e-125a0477f536";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this ExtendedZoneResource
ExtendedZoneCollection collection = subscriptionResource.GetExtendedZones();

// invoke the operation
string extendedZoneName = "losangeles";
NullableResponse<ExtendedZoneResource> response = await collection.GetIfExistsAsync(extendedZoneName);
ExtendedZoneResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ExtendedZoneData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
