using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Marketplace;
using Azure.ResourceManager.Marketplace.Models;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/BulkCollectionsAction.json
// this example is just showing the usage of "PrivateStore_BulkCollectionsAction" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateStoreResource created on azure
// for more information of creating PrivateStoreResource, please refer to the document of PrivateStoreResource
Guid privateStoreId = Guid.Parse("a0e28e55-90c4-41d8-8e34-bb7ef7775406");
ResourceIdentifier privateStoreResourceId = PrivateStoreResource.CreateResourceIdentifier(privateStoreId);
PrivateStoreResource privateStore = client.GetPrivateStoreResource(privateStoreResourceId);

// invoke the operation
BulkCollectionsActionContent content = new BulkCollectionsActionContent()
{
    CollectionIds =
    {
    Guid.Parse("c752f021-1c37-4af5-b82f-74c51c27b44a"),Guid.Parse("f47ef1c7-e908-4f39-ae29-db181634ad8d")
    },
    Action = "EnableCollections",
};
BulkCollectionsActionResult result = await privateStore.PerformActionOnBulkCollectionsAsync(content: content);

Console.WriteLine($"Succeeded: {result}");
