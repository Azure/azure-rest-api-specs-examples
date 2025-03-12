using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementDeleteWorkspaceApiOperation.json
// this example is just showing the usage of "WorkspaceApiOperation_Delete" operation, for the dependent resources, they will have to be created separately.

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
string apiId = "57d2ef278aa04f0888cba3f3";
string operationId = "57d2ef278aa04f0ad01d6cdc";
ResourceIdentifier serviceWorkspaceApiOperationResourceId = ServiceWorkspaceApiOperationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, workspaceId, apiId, operationId);
ServiceWorkspaceApiOperationResource serviceWorkspaceApiOperation = client.GetServiceWorkspaceApiOperationResource(serviceWorkspaceApiOperationResourceId);

// invoke the operation
ETag ifMatch = new ETag("*");
await serviceWorkspaceApiOperation.DeleteAsync(WaitUntil.Completed, ifMatch);

Console.WriteLine("Succeeded");
