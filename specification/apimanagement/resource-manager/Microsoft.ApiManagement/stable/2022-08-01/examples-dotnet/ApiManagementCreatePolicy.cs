using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreatePolicy.json
// this example is just showing the usage of "Policy_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementPolicyResource created on azure
// for more information of creating ApiManagementPolicyResource, please refer to the document of ApiManagementPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
PolicyName policyId = PolicyName.Policy;
ResourceIdentifier apiManagementPolicyResourceId = ApiManagementPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, policyId);
ApiManagementPolicyResource apiManagementPolicy = client.GetApiManagementPolicyResource(apiManagementPolicyResourceId);

// invoke the operation
PolicyContractData data = new PolicyContractData()
{
    Value = "<policies>\r\n  <inbound />\r\n  <backend>\r\n    <forward-request />\r\n  </backend>\r\n  <outbound />\r\n</policies>",
    Format = PolicyContentFormat.Xml,
};
ArmOperation<ApiManagementPolicyResource> lro = await apiManagementPolicy.UpdateAsync(WaitUntil.Completed, data);
ApiManagementPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
