using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRoutePortAuthorizationDelete.json
// this example is just showing the usage of "ExpressRoutePortAuthorizations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRoutePortAuthorizationResource created on azure
// for more information of creating ExpressRoutePortAuthorizationResource, please refer to the document of ExpressRoutePortAuthorizationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string expressRoutePortName = "expressRoutePortName";
string authorizationName = "authorizationName";
ResourceIdentifier expressRoutePortAuthorizationResourceId = ExpressRoutePortAuthorizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, expressRoutePortName, authorizationName);
ExpressRoutePortAuthorizationResource expressRoutePortAuthorization = client.GetExpressRoutePortAuthorizationResource(expressRoutePortAuthorizationResourceId);

// invoke the operation
await expressRoutePortAuthorization.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
