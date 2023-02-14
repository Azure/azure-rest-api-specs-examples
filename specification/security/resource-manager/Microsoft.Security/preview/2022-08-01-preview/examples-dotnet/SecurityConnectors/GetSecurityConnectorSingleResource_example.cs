using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2022-08-01-preview/examples/SecurityConnectors/GetSecurityConnectorSingleResource_example.json
// this example is just showing the usage of "SecurityConnectors_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "a5caac9c-5c04-49af-b3d0-e204f40345d5";
string resourceGroupName = "exampleResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SecurityConnectorResource
SecurityConnectorCollection collection = resourceGroupResource.GetSecurityConnectors();

// invoke the operation
string securityConnectorName = "exampleSecurityConnectorName";
bool result = await collection.ExistsAsync(securityConnectorName);

Console.WriteLine($"Succeeded: {result}");
