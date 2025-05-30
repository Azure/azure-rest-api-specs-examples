
import com.azure.resourcemanager.support.fluent.models.SupportTicketDetailsInner;
import com.azure.resourcemanager.support.models.Consent;
import com.azure.resourcemanager.support.models.ContactProfile;
import com.azure.resourcemanager.support.models.PreferredContactMethod;
import com.azure.resourcemanager.support.models.SecondaryConsent;
import com.azure.resourcemanager.support.models.SeverityLevel;
import com.azure.resourcemanager.support.models.UserConsent;
import java.util.Arrays;

/**
 * Samples for SupportTicketsNoSubscription Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/CreateTechnicalSupportTicket.
     * json
     */
    /**
     * Sample code: Create a ticket for Technical issue related to a specific resource.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketForTechnicalIssueRelatedToASpecificResource(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTicketsNoSubscriptions().create("testticket", new SupportTicketDetailsInner()
            .withDescription("my description")
            .withProblemClassificationId(
                "/providers/Microsoft.Support/services/virtual_machine_running_linux_service_guid/problemClassifications/problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withProblemScopingQuestions(
                "{\"articleId\":\"076846c1-4c0b-4b21-91c6-1a30246b3867\",\"scopingDetails\":[{\"question\":\"When did the problem begin?\",\"controlId\":\"problem_start_time\",\"orderId\":1,\"inputType\":\"static\",\"answer\":{\"displayValue\":\"2023-08-31T18:55:00.739Z\",\"value\":\"2023-08-31T18:55:00.739Z\",\"type\":\"datetime\"}},{\"question\":\"API Type of the Cosmos DB account\",\"controlId\":\"api_type\",\"orderId\":2,\"inputType\":\"static\",\"answer\":{\"displayValue\":\"Table\",\"value\":\"tables\",\"type\":\"string\"}},{\"question\":\"Table name\",\"controlId\":\"collection_name_table\",\"orderId\":11,\"inputType\":\"nonstatic\",\"answer\":{\"displayValue\":\"Select Table Name\",\"value\":\"dont_know_answer\",\"type\":\"string\"}},{\"question\":\"Provide additional details about the issue you're facing\",\"controlId\":\"problem_description\",\"orderId\":12,\"inputType\":\"nonstatic\",\"answer\":{\"displayValue\":\"test ticket, please ignore and close\",\"value\":\"test ticket, please ignore and close\",\"type\":\"string\"}}]}")
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title")
            .withServiceId("/providers/Microsoft.Support/services/cddd3eb5-1830-b494-44fd-782f691479dc")
            .withFileWorkspaceName("6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066")
            .withSecondaryConsent(Arrays.asList(
                new SecondaryConsent().withUserConsent(UserConsent.YES).withType("virtualmachinerunninglinuxservice"))),
            com.azure.core.util.Context.NONE);
    }
}
