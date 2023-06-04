using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;
using Azure.ResourceManager.CostManagement.Models;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/PricesheetDownloadByBillingProfile.json
// this example is just showing the usage of "PriceSheet_DownloadByBillingProfile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
string billingAccountName = "7c05a543-80ff-571e-9f98-1063b3b53cf2:99ad03ad-2d1b-4889-a452-090ad407d25f_2019-05-31";
string billingProfileName = "2USN-TPCD-BG7-TGB";
ArmOperation<DownloadURL> lro = await tenantResource.DownloadByBillingProfilePriceSheetAsync(WaitUntil.Completed, billingAccountName, billingProfileName);
DownloadURL result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
