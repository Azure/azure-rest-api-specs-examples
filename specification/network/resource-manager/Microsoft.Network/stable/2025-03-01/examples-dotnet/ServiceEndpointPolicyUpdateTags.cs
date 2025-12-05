using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ServiceEndpointPolicyUpdateTags.json
// this example is just showing the usage of "ServiceEndpointPolicies_UpdateTags" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceEndpointPolicyResource created on azure
// for more information of creating ServiceEndpointPolicyResource, please refer to the document of ServiceEndpointPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceEndpointPolicyName = "testServiceEndpointPolicy";
ResourceIdentifier serviceEndpointPolicyResourceId = ServiceEndpointPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceEndpointPolicyName);
ServiceEndpointPolicyResource serviceEndpointPolicy = client.GetServiceEndpointPolicyResource(serviceEndpointPolicyResourceId);

// invoke the operation
NetworkTagsObject networkTagsObject = new NetworkTagsObject
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
ServiceEndpointPolicyResource result = await serviceEndpointPolicy.UpdateAsync(networkTagsObject);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceEndpointPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
