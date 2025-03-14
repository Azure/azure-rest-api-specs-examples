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

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/BenefitUtilizationSummaries/Async/GenerateBenefitUtilizationSummariesReportByBillingProfile.json
// this example is just showing the usage of "BillingProfileScope_GenerateBenefitUtilizationSummariesReport" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
string billingAccountId = "00000000-0000-0000-0000-000000000000";
string billingProfileId = "CZSFR-SDFXC-DSDF";
BenefitUtilizationSummariesContent content = new BenefitUtilizationSummariesContent(BenefitRecommendationUsageGrain.Daily, DateTimeOffset.Parse("2022-06-01T00:00:00Z"), DateTimeOffset.Parse("2022-08-31T00:00:00Z"))
{
    Kind = BillingAccountBenefitKind.Reservation,
};
ArmOperation<BenefitUtilizationSummariesOperationStatus> lro = await tenantResource.GenerateBenefitUtilizationSummariesReportBillingProfileScopeAsync(WaitUntil.Completed, billingAccountId, billingProfileId, content);
BenefitUtilizationSummariesOperationStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
