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

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/PricesheetDownload.json
// this example is just showing the usage of "PriceSheet_Download" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
string billingAccountName = "7c05a543-80ff-571e-9f98-1063b3b53cf2:99ad03ad-2d1b-4889-a452-090ad407d25f_2019-05-31";
string billingProfileName = "2USN-TPCD-BG7-TGB";
string invoiceName = "T000940677";
ArmOperation<DownloadURL> lro = await tenantResource.DownloadPriceSheetAsync(WaitUntil.Completed, billingAccountName, billingProfileName, invoiceName);
DownloadURL result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
