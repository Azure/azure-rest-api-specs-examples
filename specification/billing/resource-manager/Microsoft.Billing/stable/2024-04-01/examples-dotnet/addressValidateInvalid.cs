using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Billing.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Billing;

// Generated from example definition: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/addressValidateInvalid.json
// this example is just showing the usage of "Address_Validate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
BillingAddressDetails details = new BillingAddressDetails("1 Test", "us")
{
    City = "bellevue",
    PostalCode = "12345",
    Region = "wa",
};
BillingAddressValidationResult result = await tenantResource.ValidateAddresAsync(details);

Console.WriteLine($"Succeeded: {result}");
