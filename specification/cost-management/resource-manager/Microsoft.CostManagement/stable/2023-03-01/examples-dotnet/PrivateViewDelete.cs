using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;
using Azure.ResourceManager.CostManagement.Models;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/PrivateViewDelete.json
// this example is just showing the usage of "Views_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantsCostManagementViewsResource created on azure
// for more information of creating TenantsCostManagementViewsResource, please refer to the document of TenantsCostManagementViewsResource
string viewName = "TestView";
ResourceIdentifier tenantsCostManagementViewsResourceId = TenantsCostManagementViewsResource.CreateResourceIdentifier(viewName);
TenantsCostManagementViewsResource tenantsCostManagementViews = client.GetTenantsCostManagementViewsResource(tenantsCostManagementViewsResourceId);

// invoke the operation
await tenantsCostManagementViews.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
