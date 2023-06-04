using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/SingleResourceGroupAlert.json
// this example is just showing the usage of "Alerts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this CostManagementAlertResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ScreenSharingTest-peer";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
CostManagementAlertCollection collection = client.GetCostManagementAlerts(scopeId);

// invoke the operation
string alertId = "22222222-2222-2222-2222-222222222222";
bool result = await collection.ExistsAsync(alertId);

Console.WriteLine($"Succeeded: {result}");
