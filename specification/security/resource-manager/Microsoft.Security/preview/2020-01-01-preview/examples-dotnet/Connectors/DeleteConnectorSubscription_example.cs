using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/DeleteConnectorSubscription_example.json
// this example is just showing the usage of "Connectors_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityCloudConnectorResource created on azure
// for more information of creating SecurityCloudConnectorResource, please refer to the document of SecurityCloudConnectorResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string connectorName = "aws_dev1";
ResourceIdentifier securityCloudConnectorResourceId = SecurityCloudConnectorResource.CreateResourceIdentifier(subscriptionId, connectorName);
SecurityCloudConnectorResource securityCloudConnector = client.GetSecurityCloudConnectorResource(securityCloudConnectorResourceId);

// invoke the operation
await securityCloudConnector.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
