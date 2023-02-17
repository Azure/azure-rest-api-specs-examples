using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Marketplace;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2022-03-01/examples/GetPrivateStoreCollection.json
// this example is just showing the usage of "PrivateStoreCollection_Get" operation, for the dependent resources, they will have to be created separately.

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
Guid collectionId = Guid.Parse("56a1a02d-8cf8-45df-bf37-d5f7120fcb3d");
bool result = await collection.ExistsAsync(collectionId);

Console.WriteLine($"Succeeded: {result}");
