using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;
using Azure.ResourceManager.CustomerInsights.Models;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorMappingsDelete.json
// this example is just showing the usage of "ConnectorMappings_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConnectorMappingResourceFormatResource created on azure
// for more information of creating ConnectorMappingResourceFormatResource, please refer to the document of ConnectorMappingResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string connectorName = "testConnector8858";
string mappingName = "testMapping12491";
ResourceIdentifier connectorMappingResourceFormatResourceId = ConnectorMappingResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, connectorName, mappingName);
ConnectorMappingResourceFormatResource connectorMappingResourceFormat = client.GetConnectorMappingResourceFormatResource(connectorMappingResourceFormatResourceId);

// invoke the operation
await connectorMappingResourceFormat.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
