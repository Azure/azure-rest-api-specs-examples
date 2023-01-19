using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApiManagement;
using Azure.ResourceManager.ApiManagement.Models;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiPolicyNonXmlEncoded.json
// this example is just showing the usage of "ApiPolicy_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiPolicyResource created on azure
// for more information of creating ApiPolicyResource, please refer to the document of ApiPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "5600b57e7e8880006a040001";
PolicyName policyId = PolicyName.Policy;
ResourceIdentifier apiPolicyResourceId = ApiPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId, policyId);
ApiPolicyResource apiPolicy = client.GetApiPolicyResource(apiPolicyResourceId);

// invoke the operation
PolicyContractData data = new PolicyContractData()
{
    Value = "<policies>\r\n     <inbound>\r\n     <base />\r\n  <set-header name=\"newvalue\" exists-action=\"override\">\r\n   <value>\"@(context.Request.Headers.FirstOrDefault(h => h.Ke==\"Via\"))\" </value>\r\n    </set-header>\r\n  </inbound>\r\n      </policies>",
    Format = PolicyContentFormat.RawXml,
};
ETag? ifMatch = new ETag("*");
ArmOperation<ApiPolicyResource> lro = await apiPolicy.UpdateAsync(WaitUntil.Completed, data, ifMatch: ifMatch);
ApiPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
