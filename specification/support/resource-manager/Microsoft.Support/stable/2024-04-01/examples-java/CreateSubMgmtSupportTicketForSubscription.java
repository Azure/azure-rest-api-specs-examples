
import com.azure.resourcemanager.support.models.Consent;
import com.azure.resourcemanager.support.models.ContactProfile;
import com.azure.resourcemanager.support.models.PreferredContactMethod;
import com.azure.resourcemanager.support.models.SeverityLevel;

/**
 * Samples for SupportTickets Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * CreateSubMgmtSupportTicketForSubscription.json
     */
    /**
     * Sample code: Create a ticket for Subscription Management related issues for a subscription.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void createATicketForSubscriptionManagementRelatedIssuesForASubscription(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket").withDescription("my description").withProblemClassificationId(
            "/providers/Microsoft.Support/services/subscription_management_service_guid/problemClassifications/subscription_management_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.NO)
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title")
            .withServiceId("/providers/Microsoft.Support/services/subscription_management_service_guid")
            .withSupportPlanId(
                "U291cmNlOlNDTSxDbGFyaWZ5SW5zdGFsbGF0aW9uU2l0ZUlkOjcsTGluZUl0ZW1JZDo5ODY1NzIyOSxDb250cmFjdElkOjk4NjU5MTk0LFN1YnNjcmlwdGlvbklkOjc2Y2I3N2ZhLThiMTctNGVhYi05NDkzLWI2NWRhY2U5OTgxMyw=")
            .withFileWorkspaceName("6f16735c-1530836f-e9970f1a-2e49-47b7-96cd-9746b83aa066").create();
    }
}
