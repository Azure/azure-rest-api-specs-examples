using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/transactionsListByBillingProfile.json
// this example is just showing the usage of "Transactions_ListByBillingProfile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProfileResource created on azure
// for more information of creating BillingProfileResource, please refer to the document of BillingProfileResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string billingProfileName = "xxxx-xxxx-xxx-xxx";
ResourceIdentifier billingProfileResourceId = BillingProfileResource.CreateResourceIdentifier(billingAccountName, billingProfileName);
BillingProfileResource billingProfile = client.GetBillingProfileResource(billingProfileResourceId);

// invoke the operation and iterate over the result
DateTimeOffset periodStartDate = DateTimeOffset.Parse("2024-04-01");
DateTimeOffset periodEndDate = DateTimeOffset.Parse("2023-05-30");
TransactionType type = TransactionType.Billed;
BillingProfileResourceGetTransactionsOptions options = new BillingProfileResourceGetTransactionsOptions(periodStartDate, periodEndDate, type) { Filter = "properties/date gt '2020-10-01'", Search = "storage" };
await foreach (BillingTransactionData item in billingProfile.GetTransactionsAsync(options))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
