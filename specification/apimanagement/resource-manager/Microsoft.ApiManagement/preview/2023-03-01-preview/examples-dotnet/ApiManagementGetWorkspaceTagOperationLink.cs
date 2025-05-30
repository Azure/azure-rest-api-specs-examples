using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementGetWorkspaceTagOperationLink.json
// this example is just showing the usage of "WorkspaceTagOperationLink_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceTagOperationLinkResource created on azure
// for more information of creating ServiceWorkspaceTagOperationLinkResource, please refer to the document of ServiceWorkspaceTagOperationLinkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string tagId = "tag1";
string operationLinkId = "link1";
ResourceIdentifier serviceWorkspaceTagOperationLinkResourceId = ServiceWorkspaceTagOperationLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, tagId, operationLinkId);
ServiceWorkspaceTagOperationLinkResource serviceWorkspaceTagOperationLink = client.GetServiceWorkspaceTagOperationLinkResource(serviceWorkspaceTagOperationLinkResourceId);

// invoke the operation
ServiceWorkspaceTagOperationLinkResource result = await serviceWorkspaceTagOperationLink.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TagOperationLinkContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
