using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingRuleGet.json
// this example is just showing the usage of "DataMaskingRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseDataMaskingPolicyResource created on azure
// for more information of creating SynapseDataMaskingPolicyResource, please refer to the document of SynapseDataMaskingPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-6852";
string workspaceName = "sqlcrudtest-2080";
string sqlPoolName = "sqlcrudtest-331";
ResourceIdentifier synapseDataMaskingPolicyResourceId = SynapseDataMaskingPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName);
SynapseDataMaskingPolicyResource synapseDataMaskingPolicy = client.GetSynapseDataMaskingPolicyResource(synapseDataMaskingPolicyResourceId);

// get the collection of this SynapseDataMaskingRuleResource
SynapseDataMaskingRuleCollection collection = synapseDataMaskingPolicy.GetSynapseDataMaskingRules();

// invoke the operation
string dataMaskingRuleName = "rule1";
NullableResponse<SynapseDataMaskingRuleResource> response = await collection.GetIfExistsAsync(dataMaskingRuleName);
SynapseDataMaskingRuleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SynapseDataMaskingRuleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
