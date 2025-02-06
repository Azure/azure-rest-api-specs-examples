using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceNetworking;

// Generated from example definition: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/stable/2025-01-01/examples/SecurityPolicyPut.json
// this example is just showing the usage of "SecurityPoliciesInterface_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TrafficControllerResource created on azure
// for more information of creating TrafficControllerResource, please refer to the document of TrafficControllerResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string trafficControllerName = "tc1";
ResourceIdentifier trafficControllerResourceId = TrafficControllerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, trafficControllerName);
TrafficControllerResource trafficController = client.GetTrafficControllerResource(trafficControllerResourceId);

// get the collection of this ApplicationGatewayForContainersSecurityPolicyResource
ApplicationGatewayForContainersSecurityPolicyCollection collection = trafficController.GetApplicationGatewayForContainersSecurityPolicies();

// invoke the operation
string securityPolicyName = "sp1";
ApplicationGatewayForContainersSecurityPolicyData data = new ApplicationGatewayForContainersSecurityPolicyData(new AzureLocation("NorthCentralUS"))
{
    WafPolicyId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Networking/applicationGatewayWebApplicationFirewallPolicies/wp-0"),
};
ArmOperation<ApplicationGatewayForContainersSecurityPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, securityPolicyName, data);
ApplicationGatewayForContainersSecurityPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApplicationGatewayForContainersSecurityPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
