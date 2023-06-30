using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/RoutePolicies_Get_MaximumSet_Gen.json
// this example is just showing the usage of "RoutePolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RoutePolicyResource created on azure
// for more information of creating RoutePolicyResource, please refer to the document of RoutePolicyResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "rgRoutePolicies";
string routePolicyName = "routePolicyName";
ResourceIdentifier routePolicyResourceId = RoutePolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, routePolicyName);
RoutePolicyResource routePolicy = client.GetRoutePolicyResource(routePolicyResourceId);

// invoke the operation
RoutePolicyResource result = await routePolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RoutePolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
