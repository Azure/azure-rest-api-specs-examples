using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CostManagement.Models;
using Azure.ResourceManager.CostManagement;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/ExportRunBySubscription.json
// this example is just showing the usage of "Exports_Execute" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CostManagementExportResource created on azure
// for more information of creating CostManagementExportResource, please refer to the document of CostManagementExportResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000";
string exportName = "TestExport";
ResourceIdentifier costManagementExportResourceId = CostManagementExportResource.CreateResourceIdentifier(scope, exportName);
CostManagementExportResource costManagementExport = client.GetCostManagementExportResource(costManagementExportResourceId);

// invoke the operation
await costManagementExport.ExecuteAsync();

Console.WriteLine("Succeeded");
