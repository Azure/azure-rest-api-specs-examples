using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;
using Azure.ResourceManager.ApiManagement.Models;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiOperationPolicy.json
// this example is just showing the usage of "ApiOperationPolicy_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiOperationResource created on azure
// for more information of creating ApiOperationResource, please refer to the document of ApiOperationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "5600b539c53f5b0062040001";
string operationId = "5600b53ac53f5b0062080006";
ResourceIdentifier apiOperationResourceId = ApiOperationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, operationId);
ApiOperationResource apiOperation = client.GetApiOperationResource(apiOperationResourceId);

// get the collection of this ApiOperationPolicyResource
ApiOperationPolicyCollection collection = apiOperation.GetApiOperationPolicies();

// invoke the operation
PolicyName policyId = PolicyName.Policy;
NullableResponse<ApiOperationPolicyResource> response = await collection.GetIfExistsAsync(policyId);
ApiOperationPolicyResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PolicyContractData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
