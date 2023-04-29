using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/LoadBalancerFrontendIPConfigurationGet.json
// this example is just showing the usage of "LoadBalancerFrontendIPConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LoadBalancerResource created on azure
// for more information of creating LoadBalancerResource, please refer to the document of LoadBalancerResource
string subscriptionId = "subid";
string resourceGroupName = "testrg";
string loadBalancerName = "lb";
ResourceIdentifier loadBalancerResourceId = LoadBalancerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, loadBalancerName);
LoadBalancerResource loadBalancer = client.GetLoadBalancerResource(loadBalancerResourceId);

// get the collection of this FrontendIPConfigurationResource
FrontendIPConfigurationCollection collection = loadBalancer.GetFrontendIPConfigurations();

// invoke the operation
string frontendIPConfigurationName = "frontend";
bool result = await collection.ExistsAsync(frontendIPConfigurationName);

Console.WriteLine($"Succeeded: {result}");
