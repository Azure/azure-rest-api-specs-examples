using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceBus;
using Azure.ResourceManager.ServiceBus.Models;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Migrationconfigurations/SBMigrationconfigurationGet.json
// this example is just showing the usage of "MigrationConfigs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusNamespaceResource created on azure
// for more information of creating ServiceBusNamespaceResource, please refer to the document of ServiceBusNamespaceResource
string subscriptionId = "SubscriptionId";
string resourceGroupName = "ResourceGroup";
string namespaceName = "sdk-Namespace-41";
ResourceIdentifier serviceBusNamespaceResourceId = ServiceBusNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
ServiceBusNamespaceResource serviceBusNamespace = client.GetServiceBusNamespaceResource(serviceBusNamespaceResourceId);

// get the collection of this MigrationConfigurationResource
MigrationConfigurationCollection collection = serviceBusNamespace.GetMigrationConfigurations();

// invoke the operation
MigrationConfigurationName configName = MigrationConfigurationName.Default;
bool result = await collection.ExistsAsync(configName);

Console.WriteLine($"Succeeded: {result}");
