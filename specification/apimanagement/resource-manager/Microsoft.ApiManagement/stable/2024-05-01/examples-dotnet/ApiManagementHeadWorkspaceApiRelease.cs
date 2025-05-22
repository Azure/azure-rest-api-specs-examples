using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementHeadWorkspaceApiRelease.json
// this example is just showing the usage of "WorkspaceApiRelease_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceApiReleaseResource created on azure
// for more information of creating ServiceWorkspaceApiReleaseResource, please refer to the document of ServiceWorkspaceApiReleaseResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string apiId = "a1";
string releaseId = "5a7cb545298324c53224a799";
ResourceIdentifier serviceWorkspaceApiReleaseResourceId = ServiceWorkspaceApiReleaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, apiId, releaseId);
ServiceWorkspaceApiReleaseResource serviceWorkspaceApiRelease = client.GetServiceWorkspaceApiReleaseResource(serviceWorkspaceApiReleaseResourceId);

// invoke the operation
bool result = await serviceWorkspaceApiRelease.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
