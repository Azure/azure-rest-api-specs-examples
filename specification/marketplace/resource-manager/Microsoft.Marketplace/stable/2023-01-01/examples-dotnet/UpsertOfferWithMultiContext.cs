using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Marketplace.Models;
using Azure.ResourceManager.Marketplace;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/UpsertOfferWithMultiContext.json
// this example is just showing the usage of "PrivateStoreCollectionOffer_UpsertOfferWithMultiContext" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateStoreOfferResource created on azure
// for more information of creating PrivateStoreOfferResource, please refer to the document of PrivateStoreOfferResource
Guid privateStoreId = Guid.Parse("a0e28e55-90c4-41d8-8e34-bb7ef7775406");
Guid collectionId = Guid.Parse("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d");
string offerId = "contoso.logger";
ResourceIdentifier privateStoreOfferResourceId = PrivateStoreOfferResource.CreateResourceIdentifier(privateStoreId, collectionId, offerId);
PrivateStoreOfferResource privateStoreOffer = client.GetPrivateStoreOfferResource(privateStoreOfferResourceId);

// invoke the operation
MultiContextAndPlansContent content = new MultiContextAndPlansContent
{
    OfferId = "contoso.logger",
    ETag = new ETag("\"9301f4fd-0000-0100-0000-5e248b350332\""),
    PlansContext = {new ContextAndPlansDetails
    {
    Context = "a5edbe7d-9f73-47fd-834a-0d6142f4c7a1",
    PlanIds = {"log4db", "log4file"},
    }, new ContextAndPlansDetails
    {
    Context = "45b604af-19bb-448e-a761-4a6be7374b2f",
    PlanIds = {"log4web"},
    }},
};
PrivateStoreOfferResource result = await privateStoreOffer.UpsertOfferWithMultiContextAsync(content: content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PrivateStoreOfferData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
