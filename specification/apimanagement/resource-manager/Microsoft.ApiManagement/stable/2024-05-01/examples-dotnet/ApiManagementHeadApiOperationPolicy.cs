using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementHeadApiOperationPolicy.json
// this example is just showing the usage of "ApiOperationPolicy_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiOperationPolicyResource created on azure
// for more information of creating ApiOperationPolicyResource, please refer to the document of ApiOperationPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "5600b539c53f5b0062040001";
string operationId = "5600b53ac53f5b0062080006";
PolicyName policyId = PolicyName.Policy;
ResourceIdentifier apiOperationPolicyResourceId = ApiOperationPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, operationId, policyId);
ApiOperationPolicyResource apiOperationPolicy = client.GetApiOperationPolicyResource(apiOperationPolicyResourceId);

// invoke the operation
bool result = await apiOperationPolicy.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
