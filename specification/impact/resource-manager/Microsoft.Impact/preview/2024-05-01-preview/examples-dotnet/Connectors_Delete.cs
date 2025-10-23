using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ImpactReporting.Models;
using Azure.ResourceManager.ImpactReporting;

// Generated from example definition: 2024-05-01-preview/Connectors_Delete.json
// this example is just showing the usage of "Connector_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ImpactConnectorResource created on azure
// for more information of creating ImpactConnectorResource, please refer to the document of ImpactConnectorResource
string subscriptionId = "8F74B371-8573-4773-9BDA-D546505BDB3A";
string connectorName = "testconnector1";
ResourceIdentifier impactConnectorResourceId = ImpactConnectorResource.CreateResourceIdentifier(subscriptionId, connectorName);
ImpactConnectorResource impactConnector = client.GetImpactConnectorResource(impactConnectorResourceId);

// invoke the operation
await impactConnector.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
