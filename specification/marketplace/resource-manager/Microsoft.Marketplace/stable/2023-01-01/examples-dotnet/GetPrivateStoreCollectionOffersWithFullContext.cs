using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Marketplace;
using Azure.ResourceManager.Marketplace.Models;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/GetPrivateStoreCollectionOffersWithFullContext.json
// this example is just showing the usage of "PrivateStoreCollectionOffer_ListByContexts" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
CollectionOffersByAllContextsPayload payload = new CollectionOffersByAllContextsPayload()
{
    SubscriptionIds =
    {
    "b340914e-353d-453a-85fb-8f9b65b51f91","f2baa04d-5bfc-461b-b6d8-61b403c9ec48"
    },
};
await foreach (CollectionOffersByContext item in privateStoreCollectionInfo.GetPrivateStoreCollectionOffersByContextsAsync(payload: payload))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
