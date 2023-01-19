using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Avs;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/Authorizations_Delete.json
// this example is just showing the usage of "Authorizations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteAuthorizationResource created on azure
// for more information of creating ExpressRouteAuthorizationResource, please refer to the document of ExpressRouteAuthorizationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
string authorizationName = "authorization1";
ResourceIdentifier expressRouteAuthorizationResourceId = ExpressRouteAuthorizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, authorizationName);
ExpressRouteAuthorizationResource expressRouteAuthorization = client.GetExpressRouteAuthorizationResource(expressRouteAuthorizationResourceId);

// invoke the operation
await expressRouteAuthorization.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
