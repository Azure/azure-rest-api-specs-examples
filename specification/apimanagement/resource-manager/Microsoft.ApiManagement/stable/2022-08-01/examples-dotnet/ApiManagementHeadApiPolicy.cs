using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiPolicy.json
// this example is just showing the usage of "ApiPolicy_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiPolicyResource created on azure
// for more information of creating ApiPolicyResource, please refer to the document of ApiPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "57d1f7558aa04f15146d9d8a";
PolicyName policyId = PolicyName.Policy;
ResourceIdentifier apiPolicyResourceId = ApiPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, policyId);
ApiPolicyResource apiPolicy = client.GetApiPolicyResource(apiPolicyResourceId);

// invoke the operation
bool result = await apiPolicy.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
