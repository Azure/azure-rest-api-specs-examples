using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateTagOperationLink.json
// this example is just showing the usage of "TagOperationLink_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceTagOperationLinkResource created on azure
// for more information of creating ServiceTagOperationLinkResource, please refer to the document of ServiceTagOperationLinkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string tagId = "tag1";
string operationLinkId = "link1";
ResourceIdentifier serviceTagOperationLinkResourceId = ServiceTagOperationLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, tagId, operationLinkId);
ServiceTagOperationLinkResource serviceTagOperationLink = client.GetServiceTagOperationLinkResource(serviceTagOperationLinkResourceId);

// invoke the operation
TagOperationLinkContractData data = new TagOperationLinkContractData
{
    OperationId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api/operations/op1",
};
ArmOperation<ServiceTagOperationLinkResource> lro = await serviceTagOperationLink.UpdateAsync(WaitUntil.Completed, data);
ServiceTagOperationLinkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TagOperationLinkContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
