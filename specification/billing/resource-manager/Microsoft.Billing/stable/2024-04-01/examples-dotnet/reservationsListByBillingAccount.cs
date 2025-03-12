using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/reservationsListByBillingAccount.json
// this example is just showing the usage of "Reservations_ListByBillingAccount" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingAccountResource created on azure
// for more information of creating BillingAccountResource, please refer to the document of BillingAccountResource
string billingAccountName = "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31";
ResourceIdentifier billingAccountResourceId = BillingAccountResource.CreateResourceIdentifier(billingAccountName);
BillingAccountResource billingAccount = client.GetBillingAccountResource(billingAccountResourceId);

// invoke the operation and iterate over the result
BillingAccountResourceGetReservationsOptions options = new BillingAccountResourceGetReservationsOptions { SelectedState = "Succeeded" };
await foreach (BillingReservationResource item in billingAccount.GetReservationsAsync(options))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BillingReservationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
