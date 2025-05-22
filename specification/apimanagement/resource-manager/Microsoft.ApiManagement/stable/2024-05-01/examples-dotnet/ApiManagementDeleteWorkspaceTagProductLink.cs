using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeleteWorkspaceTagProductLink.json
// this example is just showing the usage of "WorkspaceTagProductLink_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceTagProductLinkResource created on azure
// for more information of creating ServiceWorkspaceTagProductLinkResource, please refer to the document of ServiceWorkspaceTagProductLinkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string tagId = "tag1";
string productLinkId = "link1";
ResourceIdentifier serviceWorkspaceTagProductLinkResourceId = ServiceWorkspaceTagProductLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, tagId, productLinkId);
ServiceWorkspaceTagProductLinkResource serviceWorkspaceTagProductLink = client.GetServiceWorkspaceTagProductLinkResource(serviceWorkspaceTagProductLinkResourceId);

// invoke the operation
await serviceWorkspaceTagProductLink.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
