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

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/ExternalBillingAccountAlerts.json
// this example is just showing the usage of "Alerts_ListExternal" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation and iterate over the result
ExternalCloudProviderType externalCloudProviderType = ExternalCloudProviderType.ExternalBillingAccounts;
string externalCloudProviderId = "100";
await foreach (CostManagementAlertResource item in tenantResource.GetCostManagementAlertsAsync(externalCloudProviderType, externalCloudProviderId))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CostManagementAlertData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
