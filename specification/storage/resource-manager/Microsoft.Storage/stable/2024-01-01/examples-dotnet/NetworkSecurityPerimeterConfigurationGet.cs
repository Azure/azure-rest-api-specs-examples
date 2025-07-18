using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/NetworkSecurityPerimeterConfigurationGet.json
// this example is just showing the usage of "NetworkSecurityPerimeterConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageAccountResource created on azure
// for more information of creating StorageAccountResource, please refer to the document of StorageAccountResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "res4410";
string accountName = "sto8607";
ResourceIdentifier storageAccountResourceId = StorageAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
StorageAccountResource storageAccount = client.GetStorageAccountResource(storageAccountResourceId);

// get the collection of this NetworkSecurityPerimeterConfigurationResource
NetworkSecurityPerimeterConfigurationCollection collection = storageAccount.GetNetworkSecurityPerimeterConfigurations();

// invoke the operation
string networkSecurityPerimeterConfigurationName = "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1";
NullableResponse<NetworkSecurityPerimeterConfigurationResource> response = await collection.GetIfExistsAsync(networkSecurityPerimeterConfigurationName);
NetworkSecurityPerimeterConfigurationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkSecurityPerimeterConfigurationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
