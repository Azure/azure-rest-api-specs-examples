using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementCreateWorkspaceApiOperationPolicy.json
// this example is just showing the usage of "WorkspaceApiOperationPolicy_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceApiOperationResource created on azure
// for more information of creating ServiceWorkspaceApiOperationResource, please refer to the document of ServiceWorkspaceApiOperationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string apiId = "5600b57e7e8880006a040001";
string operationId = "5600b57e7e8880006a080001";
ResourceIdentifier serviceWorkspaceApiOperationResourceId = ServiceWorkspaceApiOperationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, apiId, operationId);
ServiceWorkspaceApiOperationResource serviceWorkspaceApiOperation = client.GetServiceWorkspaceApiOperationResource(serviceWorkspaceApiOperationResourceId);

// get the collection of this ServiceWorkspaceApiOperationPolicyResource
ServiceWorkspaceApiOperationPolicyCollection collection = serviceWorkspaceApiOperation.GetServiceWorkspaceApiOperationPolicies();

// invoke the operation
PolicyName policyId = PolicyName.Policy;
PolicyContractData data = new PolicyContractData
{
    Value = "<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>",
    Format = PolicyContentFormat.Xml,
};
ETag? ifMatch = new ETag("*");
ArmOperation<ServiceWorkspaceApiOperationPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, policyId, data, ifMatch: ifMatch);
ServiceWorkspaceApiOperationPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
