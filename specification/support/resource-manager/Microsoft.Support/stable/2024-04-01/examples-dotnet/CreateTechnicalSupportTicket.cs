using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Support.Models;
using Azure.ResourceManager.Support;

// Generated from example definition: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/CreateTechnicalSupportTicket.json
// this example is just showing the usage of "SupportTicketsNoSubscription_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this TenantSupportTicketResource
TenantSupportTicketCollection collection = tenantResource.GetTenantSupportTickets();

// invoke the operation
string supportTicketName = "testticket";
SupportTicketData data = new SupportTicketData(
    "my description",
    "/providers/Microsoft.Support/services/virtual_machine_running_linux_service_guid/problemClassifications/problemClassification_guid",
    SupportSeverityLevel.Moderate,
    AdvancedDiagnosticConsent.Yes,
    new SupportContactProfile(
    "abc",
    "xyz",
    PreferredContactMethod.Email,
    "abc@contoso.com",
    "Pacific Standard Time",
    "usa",
    "en-US"),
    "my title",
    "/providers/Microsoft.Support/services/cddd3eb5-1830-b494-44fd-782f691479dc")
{
    ProblemScopingQuestions = "{\"articleId\":\"076846c1-4c0b-4b21-91c6-1a30246b3867\",\"scopingDetails\":[{\"question\":\"When did the problem begin?\",\"controlId\":\"problem_start_time\",\"orderId\":1,\"inputType\":\"static\",\"answer\":{\"displayValue\":\"2023-08-31T18:55:00.739Z\",\"value\":\"2023-08-31T18:55:00.739Z\",\"type\":\"datetime\"}},{\"question\":\"API Type of the Cosmos DB account\",\"controlId\":\"api_type\",\"orderId\":2,\"inputType\":\"static\",\"answer\":{\"displayValue\":\"Table\",\"value\":\"tables\",\"type\":\"string\"}},{\"question\":\"Table name\",\"controlId\":\"collection_name_table\",\"orderId\":11,\"inputType\":\"nonstatic\",\"answer\":{\"displayValue\":\"Select Table Name\",\"value\":\"dont_know_answer\",\"type\":\"string\"}},{\"question\":\"Provide additional details about the issue you're facing\",\"controlId\":\"problem_description\",\"orderId\":12,\"inputType\":\"nonstatic\",\"answer\":{\"displayValue\":\"test ticket, please ignore and close\",\"value\":\"test ticket, please ignore and close\",\"type\":\"string\"}}]}",
    SupportPlanId = "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=",
    FileWorkspaceName = "6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066",
    SecondaryConsent = {new SecondaryConsent
    {
    UserConsent = UserConsent.Yes,
    SecondaryConsentType = "virtualmachinerunninglinuxservice",
    }},
};
ArmOperation<TenantSupportTicketResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, supportTicketName, data);
TenantSupportTicketResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SupportTicketData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
