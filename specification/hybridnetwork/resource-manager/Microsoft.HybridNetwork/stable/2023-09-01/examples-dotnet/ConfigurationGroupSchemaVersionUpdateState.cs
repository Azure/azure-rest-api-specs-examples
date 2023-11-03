using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ConfigurationGroupSchemaVersionUpdateState.json
// this example is just showing the usage of "ConfigurationGroupSchemas_updateState" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConfigurationGroupSchemaResource created on azure
// for more information of creating ConfigurationGroupSchemaResource, please refer to the document of ConfigurationGroupSchemaResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string publisherName = "testPublisher";
string configurationGroupSchemaName = "testConfigurationGroupSchema";
ResourceIdentifier configurationGroupSchemaResourceId = ConfigurationGroupSchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName, configurationGroupSchemaName);
ConfigurationGroupSchemaResource configurationGroupSchema = client.GetConfigurationGroupSchemaResource(configurationGroupSchemaResourceId);

// invoke the operation
ConfigurationGroupSchemaVersionUpdateState configurationGroupSchemaVersionUpdateState = new ConfigurationGroupSchemaVersionUpdateState()
{
    VersionState = VersionState.Active,
};
ArmOperation<ConfigurationGroupSchemaVersionUpdateState> lro = await configurationGroupSchema.UpdateStateAsync(WaitUntil.Completed, configurationGroupSchemaVersionUpdateState);
ConfigurationGroupSchemaVersionUpdateState result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
