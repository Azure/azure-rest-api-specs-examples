using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;
using Azure.ResourceManager.CostManagement.Models;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/ExternalBillingAccountsQuery.json
// this example is just showing the usage of "Query_UsageByExternalCloudProviderType" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
ExternalCloudProviderType externalCloudProviderType = ExternalCloudProviderType.ExternalBillingAccounts;
string externalCloudProviderId = "100";
QueryDefinition queryDefinition = new QueryDefinition(ExportType.Usage, TimeframeType.MonthToDate, new QueryDataset()
{
    Granularity = GranularityType.Daily,
    Filter = new QueryFilter()
    {
        And =
        {
        new QueryFilter()
        {
        Or =
        {
        new QueryFilter()
        {
        Dimensions = new QueryComparisonExpression("ResourceLocation",QueryOperatorType.In,new string[]
        {
        "East US","West Europe"
        }),
        },new QueryFilter()
        {
        Tags = new QueryComparisonExpression("Environment",QueryOperatorType.In,new string[]
        {
        "UAT","Prod"
        }),
        }
        },
        },new QueryFilter()
        {
        Dimensions = new QueryComparisonExpression("ResourceGroup",QueryOperatorType.In,new string[]
        {
        "API"
        }),
        }
        },
    },
});
QueryResult result = await tenantResource.UsageByExternalCloudProviderTypeQueryAsync(externalCloudProviderType, externalCloudProviderId, queryDefinition);

Console.WriteLine($"Succeeded: {result}");
