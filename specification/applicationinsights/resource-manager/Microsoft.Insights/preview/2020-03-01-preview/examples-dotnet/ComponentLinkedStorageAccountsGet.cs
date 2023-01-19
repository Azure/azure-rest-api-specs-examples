using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApplicationInsights;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/ComponentLinkedStorageAccountsGet.json
// this example is just showing the usage of "ComponentLinkedStorageAccounts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4918";
string resourceGroupName = "someResourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ComponentLinkedStorageAccountResource
ComponentLinkedStorageAccountCollection collection = resourceGroupResource.GetComponentLinkedStorageAccounts();

// invoke the operation
string resourceName = "myComponent";
StorageType storageType = StorageType.ServiceProfiler;
bool result = await collection.ExistsAsync(resourceName, storageType);

Console.WriteLine($"Succeeded: {result}");
