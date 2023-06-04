using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;
using Azure.ResourceManager.CostManagement.Models;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/BenefitUtilizationSummaries/SavingsPlan-SavingsPlanOrderId-Daily.json
// this example is just showing the usage of "BenefitUtilizationSummaries_ListBySavingsPlanOrder" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation and iterate over the result
string savingsPlanOrderId = "66cccc66-6ccc-6c66-666c-66cc6c6c66c6";
await foreach (BenefitUtilizationSummary item in tenantResource.GetBenefitUtilizationSummariesBySavingsPlanOrderAsync(savingsPlanOrderId))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
