using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGraphQLApiResolverPolicy.json
// this example is just showing the usage of "GraphQLApiResolverPolicy_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceApiResolverPolicyResource created on azure
// for more information of creating ServiceApiResolverPolicyResource, please refer to the document of ServiceApiResolverPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "5600b57e7e8880006a040001";
string resolverId = "5600b57e7e8880006a080001";
PolicyName policyId = PolicyName.Policy;
ResourceIdentifier serviceApiResolverPolicyResourceId = ServiceApiResolverPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, resolverId, policyId);
ServiceApiResolverPolicyResource serviceApiResolverPolicy = client.GetServiceApiResolverPolicyResource(serviceApiResolverPolicyResourceId);

// invoke the operation
PolicyContractData data = new PolicyContractData()
{
    Value = "<http-data-source><http-request><set-method>GET</set-method><set-backend-service base-url=\"https://some.service.com\" /><set-url>/api/users</set-url></http-request></http-data-source>",
    Format = PolicyContentFormat.Xml,
};
ETag? ifMatch = new ETag("*");
ArmOperation<ServiceApiResolverPolicyResource> lro = await serviceApiResolverPolicy.UpdateAsync(WaitUntil.Completed, data, ifMatch: ifMatch);
ServiceApiResolverPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
