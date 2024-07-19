using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Support.Models;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/UpdateSeverityOfSupportTicket.json
// this example is just showing the usage of "SupportTicketsNoSubscription_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantSupportTicketResource created on azure
// for more information of creating TenantSupportTicketResource, please refer to the document of TenantSupportTicketResource
string supportTicketName = "testticket";
ResourceIdentifier tenantSupportTicketResourceId = TenantSupportTicketResource.CreateResourceIdentifier(supportTicketName);
TenantSupportTicketResource tenantSupportTicket = client.GetTenantSupportTicketResource(tenantSupportTicketResourceId);

// invoke the operation
UpdateSupportTicket updateSupportTicket = new UpdateSupportTicket()
{
    Severity = SupportSeverityLevel.Critical,
};
TenantSupportTicketResource result = await tenantSupportTicket.UpdateAsync(updateSupportTicket);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SupportTicketData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
