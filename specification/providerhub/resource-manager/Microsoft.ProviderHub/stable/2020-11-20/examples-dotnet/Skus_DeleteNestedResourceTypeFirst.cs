using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ProviderHub;
using Azure.ResourceManager.ProviderHub.Models;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_DeleteNestedResourceTypeFirst.json
// this example is just showing the usage of "Skus_DeleteNestedResourceTypeFirst" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NestedResourceTypeFirstSkuResource created on azure
// for more information of creating NestedResourceTypeFirstSkuResource, please refer to the document of NestedResourceTypeFirstSkuResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
string resourceType = "testResourceType";
string nestedResourceTypeFirst = "nestedResourceTypeFirst";
string sku = "testSku";
ResourceIdentifier nestedResourceTypeFirstSkuResourceId = NestedResourceTypeFirstSkuResource.CreateResourceIdentifier(subscriptionId, providerNamespace, resourceType, nestedResourceTypeFirst, sku);
NestedResourceTypeFirstSkuResource nestedResourceTypeFirstSku = client.GetNestedResourceTypeFirstSkuResource(nestedResourceTypeFirstSkuResourceId);

// invoke the operation
await nestedResourceTypeFirstSku.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
