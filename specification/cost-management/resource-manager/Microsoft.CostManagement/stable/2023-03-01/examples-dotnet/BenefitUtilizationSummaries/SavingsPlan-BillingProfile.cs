using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;
using Azure.ResourceManager.CostManagement.Models;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/BenefitUtilizationSummaries/SavingsPlan-BillingProfile.json
// this example is just showing the usage of "BenefitUtilizationSummaries_ListByBillingProfileId" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation and iterate over the result
string billingAccountId = "c0a00000-0e04-5ee3-000e-f0c6e00000ec:c0a00000-0e04-5ee3-000e-f0c6e00000ec";
string billingProfileId = "200e5e90-000e-4960-8dcd-8d00a02db000";
string filter = "properties/usageDate ge 2022-10-15 and properties/usageDate le 2022-10-18";
await foreach (BenefitUtilizationSummary item in tenantResource.GetBenefitUtilizationSummariesByBillingProfileIdAsync(billingAccountId, billingProfileId, filter: filter))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
