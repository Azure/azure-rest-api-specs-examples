using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/ImpactedResources_Get.json
// this example is just showing the usage of "ImpactedResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceHealthEventImpactedResource created on azure
// for more information of creating ResourceHealthEventImpactedResource, please refer to the document of ResourceHealthEventImpactedResource
string subscriptionId = "subscriptionId";
string eventTrackingId = "BC_1-FXZ";
string impactedResourceName = "abc-123-ghj-456";
ResourceIdentifier resourceHealthEventImpactedResourceId = ResourceHealthEventImpactedResource.CreateResourceIdentifier(subscriptionId, eventTrackingId, impactedResourceName);
ResourceHealthEventImpactedResource resourceHealthEventImpactedResource = client.GetResourceHealthEventImpactedResource(resourceHealthEventImpactedResourceId);

// invoke the operation
ResourceHealthEventImpactedResource result = await resourceHealthEventImpactedResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ResourceHealthEventImpactedResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
