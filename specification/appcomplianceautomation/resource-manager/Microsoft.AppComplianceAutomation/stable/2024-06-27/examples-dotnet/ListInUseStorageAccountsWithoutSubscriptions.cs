using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppComplianceAutomation.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppComplianceAutomation;

// Generated from example definition: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/ListInUseStorageAccountsWithoutSubscriptions.json
// this example is just showing the usage of "ProviderActions_ListInUseStorageAccounts" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
ReportListInUseStorageAccountsContent content = new ReportListInUseStorageAccountsContent();
ReportListInUseStorageAccountsResult result = await tenantResource.GetInUseStorageAccountsProviderActionAsync(content);

Console.WriteLine($"Succeeded: {result}");
