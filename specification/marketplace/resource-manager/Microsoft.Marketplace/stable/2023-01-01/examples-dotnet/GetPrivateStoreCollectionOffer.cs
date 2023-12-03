using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Marketplace;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/GetPrivateStoreCollectionOffer.json
// this example is just showing the usage of "PrivateStoreCollectionOffer_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this PrivateStoreOfferResource
PrivateStoreOfferCollection collection = privateStoreCollectionInfo.GetPrivateStoreOffers();

// invoke the operation
string offerId = "marketplacetestthirdparty.md-test-third-party-2";
NullableResponse<PrivateStoreOfferResource> response = await collection.GetIfExistsAsync(offerId);
PrivateStoreOfferResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PrivateStoreOfferData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
