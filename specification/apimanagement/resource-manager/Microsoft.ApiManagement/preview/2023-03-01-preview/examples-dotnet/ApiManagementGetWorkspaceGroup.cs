using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementGetWorkspaceGroup.json
// this example is just showing the usage of "WorkspaceGroup_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkspaceContractResource created on azure
// for more information of creating WorkspaceContractResource, please refer to the document of WorkspaceContractResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
ResourceIdentifier workspaceContractResourceId = WorkspaceContractResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId);
WorkspaceContractResource workspaceContract = client.GetWorkspaceContractResource(workspaceContractResourceId);

// get the collection of this ServiceWorkspaceGroupResource
ServiceWorkspaceGroupCollection collection = workspaceContract.GetServiceWorkspaceGroups();

// invoke the operation
string groupId = "59306a29e4bbd510dc24e5f9";
NullableResponse<ServiceWorkspaceGroupResource> response = await collection.GetIfExistsAsync(groupId);
ServiceWorkspaceGroupResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ApiManagementGroupData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
