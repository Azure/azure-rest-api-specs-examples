using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementHeadWorkspaceBackend.json
// this example is just showing the usage of "WorkspaceBackend_GetEntityTag" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceWorkspaceBackendResource created on azure
// for more information of creating ServiceWorkspaceBackendResource, please refer to the document of ServiceWorkspaceBackendResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string workspaceId = "wks1";
string backendId = "sfbackend";
ResourceIdentifier serviceWorkspaceBackendResourceId = ServiceWorkspaceBackendResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, backendId);
ServiceWorkspaceBackendResource serviceWorkspaceBackend = client.GetServiceWorkspaceBackendResource(serviceWorkspaceBackendResourceId);

// invoke the operation
bool result = await serviceWorkspaceBackend.GetEntityTagAsync();

Console.WriteLine($"Succeeded: {result}");
