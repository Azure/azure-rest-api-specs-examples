using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Marketplace;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/CreatePrivateStoreCollection.json
// this example is just showing the usage of "PrivateStoreCollection_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateStoreResource created on azure
// for more information of creating PrivateStoreResource, please refer to the document of PrivateStoreResource
Guid privateStoreId = Guid.Parse("a0e28e55-90c4-41d8-8e34-bb7ef7775406");
ResourceIdentifier privateStoreResourceId = PrivateStoreResource.CreateResourceIdentifier(privateStoreId);
PrivateStoreResource privateStore = client.GetPrivateStoreResource(privateStoreResourceId);

// get the collection of this PrivateStoreCollectionInfoResource
PrivateStoreCollectionInfoCollection collection = privateStore.GetPrivateStoreCollectionInfos();

// invoke the operation
Guid collectionId = Guid.Parse("d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1");
PrivateStoreCollectionInfoData info = new PrivateStoreCollectionInfoData
{
    CollectionName = "Test Collection",
    Claim = "",
    AreAllSubscriptionsSelected = false,
    SubscriptionsList = { "b340914e-353d-453a-85fb-8f9b65b51f91", "f2baa04d-5bfc-461b-b6d8-61b403c9ec48" },
};
ArmOperation<PrivateStoreCollectionInfoResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, collectionId, info);
PrivateStoreCollectionInfoResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PrivateStoreCollectionInfoData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
