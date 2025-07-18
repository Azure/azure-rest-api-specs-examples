using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Avs;

// Generated from example definition: 2024-09-01/Authorizations_CreateOrUpdate.json
// this example is just showing the usage of "ExpressRouteAuthorization_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
ExpressRouteAuthorizationData data = new ExpressRouteAuthorizationData();
ArmOperation<ExpressRouteAuthorizationResource> lro = await expressRouteAuthorization.UpdateAsync(WaitUntil.Completed, data);
ExpressRouteAuthorizationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExpressRouteAuthorizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
