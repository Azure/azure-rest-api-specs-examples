using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppConfiguration;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/DeletedConfigurationStoresGet.json
// this example is just showing the usage of "ConfigurationStores_GetDeleted" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this DeletedAppConfigurationStoreResource
DeletedAppConfigurationStoreCollection collection = subscriptionResource.GetDeletedAppConfigurationStores();

// invoke the operation
AzureLocation location = new AzureLocation("westus");
string configStoreName = "contoso";
NullableResponse<DeletedAppConfigurationStoreResource> response = await collection.GetIfExistsAsync(location, configStoreName);
DeletedAppConfigurationStoreResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DeletedAppConfigurationStoreData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
