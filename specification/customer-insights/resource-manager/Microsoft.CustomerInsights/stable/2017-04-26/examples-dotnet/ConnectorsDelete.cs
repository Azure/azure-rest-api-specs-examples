using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;
using Azure.ResourceManager.CustomerInsights.Models;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorsDelete.json
// this example is just showing the usage of "Connectors_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConnectorResourceFormatResource created on azure
// for more information of creating ConnectorResourceFormatResource, please refer to the document of ConnectorResourceFormatResource
string subscriptionId = "subid";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string connectorName = "testConnector";
ResourceIdentifier connectorResourceFormatResourceId = ConnectorResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, connectorName);
ConnectorResourceFormatResource connectorResourceFormat = client.GetConnectorResourceFormatResource(connectorResourceFormatResourceId);

// invoke the operation
await connectorResourceFormat.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
