using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ImpactReporting.Models;
using Azure.ResourceManager.ImpactReporting;

// Generated from example definition: 2024-05-01-preview/Connectors_Update.json
// this example is just showing the usage of "Connector_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ImpactConnectorResource created on azure
// for more information of creating ImpactConnectorResource, please refer to the document of ImpactConnectorResource
string subscriptionId = "74f5e23f-d4d9-410a-bb4d-8f0608adb10d";
string connectorName = "testconnector1";
ResourceIdentifier impactConnectorResourceId = ImpactConnectorResource.CreateResourceIdentifier(subscriptionId, connectorName);
ImpactConnectorResource impactConnector = client.GetImpactConnectorResource(impactConnectorResourceId);

// invoke the operation
ImpactConnectorPatch patch = new ImpactConnectorPatch
{
    ConnectorType = ImpactConnectorType.AzureMonitor,
};
ImpactConnectorResource result = await impactConnector.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ImpactConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
