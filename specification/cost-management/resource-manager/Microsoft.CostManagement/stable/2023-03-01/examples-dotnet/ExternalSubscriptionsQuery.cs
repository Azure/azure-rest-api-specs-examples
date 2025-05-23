using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CostManagement.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CostManagement;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/ExternalSubscriptionsQuery.json
// this example is just showing the usage of "Query_UsageByExternalCloudProviderType" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
ExternalCloudProviderType externalCloudProviderType = ExternalCloudProviderType.ExternalSubscriptions;
string externalCloudProviderId = "100";
QueryDefinition queryDefinition = new QueryDefinition(ExportType.Usage, TimeframeType.MonthToDate, new QueryDataset
{
    Granularity = GranularityType.Daily,
    Filter = new QueryFilter
    {
        And = {new QueryFilter
        {
        Or = {new QueryFilter
        {
        Dimensions = new QueryComparisonExpression("ResourceLocation", QueryOperatorType.In, new string[]{"East US", "West Europe"}),
        }, new QueryFilter
        {
        Tags = new QueryComparisonExpression("Environment", QueryOperatorType.In, new string[]{"UAT", "Prod"}),
        }},
        }, new QueryFilter
        {
        Dimensions = new QueryComparisonExpression("ResourceGroup", QueryOperatorType.In, new string[]{"API"}),
        }},
    },
});
QueryResult result = await tenantResource.UsageByExternalCloudProviderTypeQueryAsync(externalCloudProviderType, externalCloudProviderId, queryDefinition);

Console.WriteLine($"Succeeded: {result}");
