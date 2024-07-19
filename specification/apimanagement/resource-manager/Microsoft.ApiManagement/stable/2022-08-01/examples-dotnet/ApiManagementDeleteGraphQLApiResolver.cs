using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteGraphQLApiResolver.json
// this example is just showing the usage of "GraphQLApiResolver_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResolverContractResource created on azure
// for more information of creating ResolverContractResource, please refer to the document of ResolverContractResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "57d2ef278aa04f0888cba3f3";
string resolverId = "57d2ef278aa04f0ad01d6cdc";
ResourceIdentifier resolverContractResourceId = ResolverContractResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, resolverId);
ResolverContractResource resolverContract = client.GetResolverContractResource(resolverContractResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
await resolverContract.DeleteAsync(WaitUntil.Completed, ifMatch);

Console.WriteLine($"Succeeded");
