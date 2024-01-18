using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2023-10-01-preview/examples/SecurityConnectors/DeleteSecurityConnector_example.json
// this example is just showing the usage of "SecurityConnectors_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityConnectorResource created on azure
// for more information of creating SecurityConnectorResource, please refer to the document of SecurityConnectorResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "myRg";
string securityConnectorName = "mySecurityConnectorName";
ResourceIdentifier securityConnectorResourceId = SecurityConnectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, securityConnectorName);
SecurityConnectorResource securityConnector = client.GetSecurityConnectorResource(securityConnectorResourceId);

// invoke the operation
await securityConnector.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
