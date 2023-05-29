using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Consumption;

// Generated from example definition: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/DeleteBudget.json
// this example is just showing the usage of "Budgets_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ConsumptionBudgetResource created on azure
// for more information of creating ConsumptionBudgetResource, please refer to the document of ConsumptionBudgetResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000";
string budgetName = "TestBudget";
ResourceIdentifier consumptionBudgetResourceId = ConsumptionBudgetResource.CreateResourceIdentifier(scope, budgetName);
ConsumptionBudgetResource consumptionBudget = client.GetConsumptionBudgetResource(consumptionBudgetResourceId);

// invoke the operation
await consumptionBudget.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
