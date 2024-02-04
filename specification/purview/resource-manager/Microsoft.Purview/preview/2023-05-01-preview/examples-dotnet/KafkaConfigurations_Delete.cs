using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Purview;
using Azure.ResourceManager.Purview.Models;

// Generated from example definition: specification/purview/resource-manager/Microsoft.Purview/preview/2023-05-01-preview/examples/KafkaConfigurations_Delete.json
// this example is just showing the usage of "KafkaConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PurviewKafkaConfigurationResource created on azure
// for more information of creating PurviewKafkaConfigurationResource, please refer to the document of PurviewKafkaConfigurationResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rgpurview";
string accountName = "account1";
string kafkaConfigurationName = "kafkaConfigName";
ResourceIdentifier purviewKafkaConfigurationResourceId = PurviewKafkaConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, kafkaConfigurationName);
PurviewKafkaConfigurationResource purviewKafkaConfiguration = client.GetPurviewKafkaConfigurationResource(purviewKafkaConfigurationResourceId);

// invoke the operation
await purviewKafkaConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
