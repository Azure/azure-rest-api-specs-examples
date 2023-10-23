import com.azure.resourcemanager.support.fluent.models.SupportTicketDetailsInner;
import com.azure.resourcemanager.support.models.ContactProfile;
import com.azure.resourcemanager.support.models.PreferredContactMethod;
import com.azure.resourcemanager.support.models.SeverityLevel;

/** Samples for SupportTicketsNoSubscription Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/CreateBillingSupportTicket.json
     */
    /**
     * Sample code: Create a ticket for Billing related issues.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketForBillingRelatedIssues(com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTicketsNoSubscriptions()
            .create(
                "testticket",
                new SupportTicketDetailsInner()
                    .withDescription("my description")
                    .withProblemClassificationId(
                        "/providers/Microsoft.Support/services/billing_service_guid/problemClassifications/billing_problemClassification_guid")
                    .withSeverity(SeverityLevel.MODERATE)
                    .withSupportPlanId(
                        "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
                    .withContactDetails(
                        new ContactProfile()
                            .withFirstName("abc")
                            .withLastName("xyz")
                            .withPreferredContactMethod(PreferredContactMethod.EMAIL)
                            .withPrimaryEmailAddress("abc@contoso.com")
                            .withPreferredTimeZone("Pacific Standard Time")
                            .withCountry("usa")
                            .withPreferredSupportLanguage("en-US"))
                    .withTitle("my title")
                    .withServiceId("/providers/Microsoft.Support/services/billing_service_guid")
                    .withFileWorkspaceName("6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066"),
                com.azure.core.util.Context.NONE);
    }
}
