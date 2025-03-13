using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/enrollmentAccountGet.json
// this example is just showing the usage of "EnrollmentAccounts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingEnrollmentAccountResource created on azure
// for more information of creating BillingEnrollmentAccountResource, please refer to the document of BillingEnrollmentAccountResource
string billingAccountName = "6564892";
string enrollmentAccountName = "257698";
ResourceIdentifier billingEnrollmentAccountResourceId = BillingEnrollmentAccountResource.CreateResourceIdentifier(billingAccountName, enrollmentAccountName);
BillingEnrollmentAccountResource billingEnrollmentAccount = client.GetBillingEnrollmentAccountResource(billingEnrollmentAccountResourceId);

// invoke the operation
BillingEnrollmentAccountResource result = await billingEnrollmentAccount.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BillingEnrollmentAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
