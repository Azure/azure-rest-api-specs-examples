using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/unregisterFeature.json
// this example is just showing the usage of "Features_Unregister" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FeatureResource created on azure
// for more information of creating FeatureResource, please refer to the document of FeatureResource
string subscriptionId = "ff23096b-f5a2-46ea-bd62-59c3e93fef9a";
string resourceProviderNamespace = "Resource Provider Namespace";
string featureName = "feature";
ResourceIdentifier featureResourceId = FeatureResource.CreateResourceIdentifier(subscriptionId, resourceProviderNamespace, featureName);
FeatureResource feature = client.GetFeatureResource(featureResourceId);

// invoke the operation
FeatureResource result = await feature.UnregisterAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FeatureData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
