using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementDeleteWorkspaceProduct.json
// this example is just showing the usage of "WorkspaceProduct_Delete" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
ETag ifMatch = new ETag("*");
bool? deleteSubscriptions = true;
await serviceWorkspaceProduct.DeleteAsync(WaitUntil.Completed, ifMatch, deleteSubscriptions: deleteSubscriptions);

Console.WriteLine($"Succeeded");
