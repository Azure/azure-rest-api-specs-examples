using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppConfiguration;

// Generated from example definition: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-05-01/examples/DeletedConfigurationStoresPurge.json
// this example is just showing the usage of "ConfigurationStores_PurgeDeleted" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeletedAppConfigurationStoreResource created on azure
// for more information of creating DeletedAppConfigurationStoreResource, please refer to the document of DeletedAppConfigurationStoreResource
string subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
AzureLocation location = new AzureLocation("westus");
string configStoreName = "contoso";
ResourceIdentifier deletedAppConfigurationStoreResourceId = DeletedAppConfigurationStoreResource.CreateResourceIdentifier(subscriptionId, location, configStoreName);
DeletedAppConfigurationStoreResource deletedAppConfigurationStore = client.GetDeletedAppConfigurationStoreResource(deletedAppConfigurationStoreResourceId);

// invoke the operation
await deletedAppConfigurationStore.PurgeDeletedAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
