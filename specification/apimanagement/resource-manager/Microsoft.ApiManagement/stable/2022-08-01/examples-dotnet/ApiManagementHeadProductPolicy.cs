using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadProductPolicy.json
// this example is just showing the usage of "ProductPolicy_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementProductPolicyResource created on azure
// for more information of creating ApiManagementProductPolicyResource, please refer to the document of ApiManagementProductPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string productId = "unlimited";
PolicyName policyId = PolicyName.Policy;
ResourceIdentifier apiManagementProductPolicyResourceId = ApiManagementProductPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, productId, policyId);
ApiManagementProductPolicyResource apiManagementProductPolicy = client.GetApiManagementProductPolicyResource(apiManagementProductPolicyResourceId);

// invoke the operation
bool result = await apiManagementProductPolicy.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
