using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ExpressRouteCircuitAuthorizationCreate.json
// this example is just showing the usage of "ExpressRouteCircuitAuthorizations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteCircuitResource created on azure
// for more information of creating ExpressRouteCircuitResource, please refer to the document of ExpressRouteCircuitResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string circuitName = "circuitName";
ResourceIdentifier expressRouteCircuitResourceId = ExpressRouteCircuitResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, circuitName);
ExpressRouteCircuitResource expressRouteCircuit = client.GetExpressRouteCircuitResource(expressRouteCircuitResourceId);

// get the collection of this ExpressRouteCircuitAuthorizationResource
ExpressRouteCircuitAuthorizationCollection collection = expressRouteCircuit.GetExpressRouteCircuitAuthorizations();

// invoke the operation
string authorizationName = "authorizatinName";
ExpressRouteCircuitAuthorizationData data = new ExpressRouteCircuitAuthorizationData();
ArmOperation<ExpressRouteCircuitAuthorizationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, authorizationName, data);
ExpressRouteCircuitAuthorizationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExpressRouteCircuitAuthorizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
