using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Support;
using Azure.ResourceManager.Support.Models;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/CheckNameAvailability.json
// this example is just showing the usage of "SupportTicketsNoSubscription_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
SupportNameAvailabilityContent content = new SupportNameAvailabilityContent("sampleName", SupportResourceType.MicrosoftSupportSupportTickets);
SupportNameAvailabilityResult result = await tenantResource.CheckNameAvailabilitySupportTicketsNoSubscriptionAsync(content);

Console.WriteLine($"Succeeded: {result}");
