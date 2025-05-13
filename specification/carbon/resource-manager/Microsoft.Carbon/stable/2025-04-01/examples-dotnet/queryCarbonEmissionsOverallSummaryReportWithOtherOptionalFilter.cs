using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CarbonOptimization.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CarbonOptimization;

// Generated from example definition: 2025-04-01/queryCarbonEmissionsOverallSummaryReportWithOtherOptionalFilter.json
// this example is just showing the usage of "CarbonService_QueryCarbonEmissionReports" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation and iterate over the result
CarbonEmissionQueryFilter queryParameters = new OverallSummaryReportQueryFilter(new CarbonEmissionQueryDateRange(DateTimeOffset.Parse("2023-06-01"), DateTimeOffset.Parse("2023-06-01")), new string[] { "00000000-0000-0000-0000-000000000000" }, new CarbonEmissionScope[] { CarbonEmissionScope.Scope1, CarbonEmissionScope.Scope3 })
{
    ResourceGroupUrlList = { "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-name" },
    ResourceTypeList = { new ResourceType("microsoft.storage/storageaccounts"), new ResourceType("microsoft.databricks/workspaces") },
    LocationList = { new AzureLocation("east us"), new AzureLocation("west us") },
};
var result = await tenantResource.QueryCarbonEmissionReportsAsync(queryParameters);

Console.WriteLine("Succeeded");
