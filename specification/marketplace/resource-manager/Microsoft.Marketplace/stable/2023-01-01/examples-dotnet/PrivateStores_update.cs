using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Marketplace;
using Azure.ResourceManager.Marketplace.Models;

// Generated from example definition: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2023-01-01/examples/PrivateStores_update.json
// this example is just showing the usage of "PrivateStore_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
PrivateStoreData data = new PrivateStoreData()
{
    Availability = PrivateStoreAvailability.Disabled,
    ETag = new ETag("\"9301f4fd-0000-0100-0000-5e248b350345\""),
};
await privateStore.UpdateAsync(WaitUntil.Completed, data);

Console.WriteLine($"Succeeded");
