using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Consumption;

// Generated from example definition: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/BudgetsList.json
// this example is just showing the usage of "Budgets_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this ConsumptionBudgetResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000";
ConsumptionBudgetCollection collection = client.GetConsumptionBudgets(new ResourceIdentifier(scope));

// invoke the operation and iterate over the result
await foreach (ConsumptionBudgetResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ConsumptionBudgetData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
