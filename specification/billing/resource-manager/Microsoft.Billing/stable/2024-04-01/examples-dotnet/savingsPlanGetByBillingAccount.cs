using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/savingsPlanGetByBillingAccount.json
// this example is just showing the usage of "SavingsPlans_GetByBillingAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingSavingsPlanModelResource created on azure
// for more information of creating BillingSavingsPlanModelResource, please refer to the document of BillingSavingsPlanModelResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string savingsPlanOrderId = "20000000-0000-0000-0000-000000000000";
string savingsPlanId = "30000000-0000-0000-0000-000000000000";
ResourceIdentifier billingSavingsPlanModelResourceId = BillingSavingsPlanModelResource.CreateResourceIdentifier(billingAccountName, savingsPlanOrderId, savingsPlanId);
BillingSavingsPlanModelResource billingSavingsPlanModel = client.GetBillingSavingsPlanModelResource(billingSavingsPlanModelResourceId);

// invoke the operation
BillingSavingsPlanModelResource result = await billingSavingsPlanModel.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingSavingsPlanModelData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
