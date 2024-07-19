using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadGraphQLApiResolverPolicy.json
// this example is just showing the usage of "GraphQLApiResolverPolicy_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceApiResolverPolicyResource created on azure
// for more information of creating ServiceApiResolverPolicyResource, please refer to the document of ServiceApiResolverPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "5600b539c53f5b0062040001";
string resolverId = "5600b53ac53f5b0062080006";
PolicyName policyId = PolicyName.Policy;
ResourceIdentifier serviceApiResolverPolicyResourceId = ServiceApiResolverPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, resolverId, policyId);
ServiceApiResolverPolicyResource serviceApiResolverPolicy = client.GetServiceApiResolverPolicyResource(serviceApiResolverPolicyResourceId);

// invoke the operation
bool result = await serviceApiResolverPolicy.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
