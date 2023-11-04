using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ConfigurationGroupValueDelete.json
// this example is just showing the usage of "ConfigurationGroupValues_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConfigurationGroupValueResource created on azure
// for more information of creating ConfigurationGroupValueResource, please refer to the document of ConfigurationGroupValueResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string configurationGroupValueName = "testConfigurationGroupValue";
ResourceIdentifier configurationGroupValueResourceId = ConfigurationGroupValueResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, configurationGroupValueName);
ConfigurationGroupValueResource configurationGroupValue = client.GetConfigurationGroupValueResource(configurationGroupValueResourceId);

// invoke the operation
await configurationGroupValue.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
