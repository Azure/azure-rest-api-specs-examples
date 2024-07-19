using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApiOperationPolicy.json
// this example is just showing the usage of "ApiOperationPolicy_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiOperationPolicyResource created on azure
// for more information of creating ApiOperationPolicyResource, please refer to the document of ApiOperationPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "testapi";
string operationId = "testoperation";
PolicyName policyId = PolicyName.Policy;
ResourceIdentifier apiOperationPolicyResourceId = ApiOperationPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, operationId, policyId);
ApiOperationPolicyResource apiOperationPolicy = client.GetApiOperationPolicyResource(apiOperationPolicyResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
await apiOperationPolicy.DeleteAsync(WaitUntil.Completed, ifMatch);

Console.WriteLine($"Succeeded");
