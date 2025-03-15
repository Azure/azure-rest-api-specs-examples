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

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/GenerateReservationDetailsReportByBillingAccount.json
// this example is just showing the usage of "GenerateReservationDetailsReport_ByBillingAccountId" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
string billingAccountId = "9845612";
string startDate = "2020-01-01";
string endDate = "2020-01-30";
ArmOperation<OperationStatus> lro = await tenantResource.ByBillingAccountIdGenerateReservationDetailsReportAsync(WaitUntil.Completed, billingAccountId, startDate, endDate);
OperationStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
