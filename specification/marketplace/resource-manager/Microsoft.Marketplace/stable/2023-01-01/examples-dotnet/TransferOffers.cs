using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Marketplace.Models;
using Azure.ResourceManager.Marketplace;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/TransferOffers.json
// this example is just showing the usage of "PrivateStoreCollection_TransferOffers" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateStoreCollectionInfoResource created on azure
// for more information of creating PrivateStoreCollectionInfoResource, please refer to the document of PrivateStoreCollectionInfoResource
Guid privateStoreId = Guid.Parse("a0e28e55-90c4-41d8-8e34-bb7ef7775406");
Guid collectionId = Guid.Parse("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d");
ResourceIdentifier privateStoreCollectionInfoResourceId = PrivateStoreCollectionInfoResource.CreateResourceIdentifier(privateStoreId, collectionId);
PrivateStoreCollectionInfoResource privateStoreCollectionInfo = client.GetPrivateStoreCollectionInfoResource(privateStoreCollectionInfoResourceId);

// invoke the operation
TransferOffersContent content = new TransferOffersContent
{
    TargetCollections = { "c752f021-1c37-4af5-b82f-74c51c27b44a", "f47ef1c7-e908-4f39-ae29-db181634ad8d" },
    Operation = "copy",
    OfferIdsList = { "marketplacetestthirdparty.md-test-third-party-2", "marketplacetestthirdparty.md-test-third-party-3" },
};
TransferOffersResult result = await privateStoreCollectionInfo.TransferOffersAsync(content: content);

Console.WriteLine($"Succeeded: {result}");
