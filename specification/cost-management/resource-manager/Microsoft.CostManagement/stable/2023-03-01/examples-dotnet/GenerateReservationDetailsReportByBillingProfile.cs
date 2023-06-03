using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;
using Azure.ResourceManager.CostManagement.Models;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/GenerateReservationDetailsReportByBillingProfile.json
// this example is just showing the usage of "GenerateReservationDetailsReport_ByBillingProfileId" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
string billingAccountId = "00000000-0000-0000-0000-000000000000";
string billingProfileId = "CZSFR-SDFXC-DSDF";
string startDate = "2020-01-01";
string endDate = "2020-01-30";
ArmOperation<OperationStatus> lro = await tenantResource.ByBillingProfileIdGenerateReservationDetailsReportAsync(WaitUntil.Completed, billingAccountId, billingProfileId, startDate, endDate);
OperationStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
