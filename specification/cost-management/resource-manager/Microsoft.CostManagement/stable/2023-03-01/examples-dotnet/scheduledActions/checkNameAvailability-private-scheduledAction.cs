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

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/scheduledActions/checkNameAvailability-private-scheduledAction.json
// this example is just showing the usage of "ScheduledActions_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
CostManagementNameAvailabilityContent content = new CostManagementNameAvailabilityContent
{
    Name = "testName",
    ResourceType = "Microsoft.CostManagement/ScheduledActions",
};
CostManagementNameAvailabilityResult result = await tenantResource.CheckCostManagementNameAvailabilityByScheduledActionAsync(content);

Console.WriteLine($"Succeeded: {result}");
