using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateWorkspaceTagProductLink.json
// this example is just showing the usage of "WorkspaceTagProductLink_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceTagResource created on azure
// for more information of creating ServiceWorkspaceTagResource, please refer to the document of ServiceWorkspaceTagResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string tagId = "tag1";
ResourceIdentifier serviceWorkspaceTagResourceId = ServiceWorkspaceTagResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, tagId);
ServiceWorkspaceTagResource serviceWorkspaceTag = client.GetServiceWorkspaceTagResource(serviceWorkspaceTagResourceId);

// get the collection of this ServiceWorkspaceTagProductLinkResource
ServiceWorkspaceTagProductLinkCollection collection = serviceWorkspaceTag.GetServiceWorkspaceTagProductLinks();

// invoke the operation
string productLinkId = "link1";
TagProductLinkContractData data = new TagProductLinkContractData
{
    ProductId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/products/product1",
};
ArmOperation<ServiceWorkspaceTagProductLinkResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, productLinkId, data);
ServiceWorkspaceTagProductLinkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TagProductLinkContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
