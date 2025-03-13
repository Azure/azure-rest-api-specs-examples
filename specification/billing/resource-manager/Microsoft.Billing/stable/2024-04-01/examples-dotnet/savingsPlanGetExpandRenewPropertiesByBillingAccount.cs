using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/savingsPlanGetExpandRenewPropertiesByBillingAccount.json
// this example is just showing the usage of "SavingsPlans_GetByBillingAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SavingsPlanOrderModelResource created on azure
// for more information of creating SavingsPlanOrderModelResource, please refer to the document of SavingsPlanOrderModelResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
string savingsPlanOrderId = "20000000-0000-0000-0000-000000000000";
ResourceIdentifier savingsPlanOrderModelResourceId = SavingsPlanOrderModelResource.CreateResourceIdentifier(billingAccountName, savingsPlanOrderId);
SavingsPlanOrderModelResource savingsPlanOrderModel = client.GetSavingsPlanOrderModelResource(savingsPlanOrderModelResourceId);

// get the collection of this BillingSavingsPlanModelResource
BillingSavingsPlanModelCollection collection = savingsPlanOrderModel.GetBillingSavingsPlanModels();

// invoke the operation
string savingsPlanId = "30000000-0000-0000-0000-000000000000";
NullableResponse<BillingSavingsPlanModelResource> response = await collection.GetIfExistsAsync(savingsPlanId);
BillingSavingsPlanModelResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BillingSavingsPlanModelData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
