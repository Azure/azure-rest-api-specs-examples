using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/productValidateMoveSuccess.json
// this example is just showing the usage of "Products_ValidateMoveEligibility" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BillingProductResource created on azure
// for more information of creating BillingProductResource, please refer to the document of BillingProductResource
string billingAccountName = "a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31";
string productName = "6b96d3f2-9008-4a9d-912f-f87744185aa3";
ResourceIdentifier billingProductResourceId = BillingProductResource.CreateResourceIdentifier(billingAccountName, productName);
BillingProductResource billingProduct = client.GetBillingProductResource(billingProductResourceId);

// invoke the operation
MoveProductContent content = new MoveProductContent(new ResourceIdentifier("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998/invoiceSections/Q7GV-UUVA-PJA-TGB"));
MoveProductEligibilityResult result = await billingProduct.ValidateMoveEligibilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
