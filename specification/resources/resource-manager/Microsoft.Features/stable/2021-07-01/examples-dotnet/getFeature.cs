using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/getFeature.json
// this example is just showing the usage of "Features_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceProviderResource created on azure
// for more information of creating ResourceProviderResource, please refer to the document of ResourceProviderResource
string subscriptionId = "subid";
string resourceProviderNamespace = "Resource Provider Namespace";
ResourceIdentifier resourceProviderResourceId = ResourceProviderResource.CreateResourceIdentifier(subscriptionId, resourceProviderNamespace);
ResourceProviderResource resourceProvider = client.GetResourceProviderResource(resourceProviderResourceId);

// get the collection of this FeatureResource
FeatureCollection collection = resourceProvider.GetFeatures();

// invoke the operation
string featureName = "feature";
NullableResponse<FeatureResource> response = await collection.GetIfExistsAsync(featureName);
FeatureResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    FeatureData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
