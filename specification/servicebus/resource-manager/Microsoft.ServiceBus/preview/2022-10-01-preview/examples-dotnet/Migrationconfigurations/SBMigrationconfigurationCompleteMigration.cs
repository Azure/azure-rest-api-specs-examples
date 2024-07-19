using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Migrationconfigurations/SBMigrationconfigurationCompleteMigration.json
// this example is just showing the usage of "MigrationConfigs_CompleteMigration" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MigrationConfigurationResource created on azure
// for more information of creating MigrationConfigurationResource, please refer to the document of MigrationConfigurationResource
string subscriptionId = "SubscriptionId";
string resourceGroupName = "ResourceGroup";
string namespaceName = "sdk-Namespace-41";
MigrationConfigurationName configName = MigrationConfigurationName.Default;
ResourceIdentifier migrationConfigurationResourceId = MigrationConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, configName);
MigrationConfigurationResource migrationConfiguration = client.GetMigrationConfigurationResource(migrationConfigurationResourceId);

// invoke the operation
await migrationConfiguration.CompleteMigrationAsync();

Console.WriteLine($"Succeeded");
