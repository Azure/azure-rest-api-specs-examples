using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetWorkspaceProductGroupLink.json
// this example is just showing the usage of "WorkspaceProductGroupLink_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceProductResource created on azure
// for more information of creating ServiceWorkspaceProductResource, please refer to the document of ServiceWorkspaceProductResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string productId = "testproduct";
ResourceIdentifier serviceWorkspaceProductResourceId = ServiceWorkspaceProductResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, productId);
ServiceWorkspaceProductResource serviceWorkspaceProduct = client.GetServiceWorkspaceProductResource(serviceWorkspaceProductResourceId);

// get the collection of this ServiceWorkspaceProductGroupLinkResource
ServiceWorkspaceProductGroupLinkCollection collection = serviceWorkspaceProduct.GetServiceWorkspaceProductGroupLinks();

// invoke the operation
string groupLinkId = "link1";
NullableResponse<ServiceWorkspaceProductGroupLinkResource> response = await collection.GetIfExistsAsync(groupLinkId);
ServiceWorkspaceProductGroupLinkResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ProductGroupLinkContractData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
