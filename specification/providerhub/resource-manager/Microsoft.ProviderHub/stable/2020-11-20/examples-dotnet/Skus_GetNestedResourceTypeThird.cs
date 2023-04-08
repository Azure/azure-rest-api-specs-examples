using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ProviderHub;
using Azure.ResourceManager.ProviderHub.Models;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_GetNestedResourceTypeThird.json
// this example is just showing the usage of "Skus_GetNestedResourceTypeThird" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceTypeRegistrationResource created on azure
// for more information of creating ResourceTypeRegistrationResource, please refer to the document of ResourceTypeRegistrationResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
string resourceType = "testResourceType";
ResourceIdentifier resourceTypeRegistrationResourceId = ResourceTypeRegistrationResource.CreateResourceIdentifier(subscriptionId, providerNamespace, resourceType);
ResourceTypeRegistrationResource resourceTypeRegistration = client.GetResourceTypeRegistrationResource(resourceTypeRegistrationResourceId);

// get the collection of this NestedResourceTypeThirdSkuResource
string nestedResourceTypeFirst = "nestedResourceTypeFirst";
string nestedResourceTypeSecond = "nestedResourceTypeSecond";
string nestedResourceTypeThird = "nestedResourceTypeThird";
NestedResourceTypeThirdSkuCollection collection = resourceTypeRegistration.GetNestedResourceTypeThirdSkus(nestedResourceTypeFirst, nestedResourceTypeSecond, nestedResourceTypeThird);

// invoke the operation
string sku = "testSku";
bool result = await collection.ExistsAsync(sku);

Console.WriteLine($"Succeeded: {result}");
