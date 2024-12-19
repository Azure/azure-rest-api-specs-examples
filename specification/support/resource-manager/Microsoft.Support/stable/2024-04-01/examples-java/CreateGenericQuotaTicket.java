
import com.azure.resourcemanager.support.models.Consent;
import com.azure.resourcemanager.support.models.ContactProfile;
import com.azure.resourcemanager.support.models.PreferredContactMethod;
import com.azure.resourcemanager.support.models.SeverityLevel;

/**
 * Samples for SupportTickets Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/CreateGenericQuotaTicket.json
     */
    /**
     * Sample code: Create a ticket to request Quota increase for services that do not require additional details in the
     * quotaTicketDetails object.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        createATicketToRequestQuotaIncreaseForServicesThatDoNotRequireAdditionalDetailsInTheQuotaTicketDetailsObject(
            com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().define("testticket")
            .withDescription("Increase the maximum throughput per container limit to 10000 for account foo bar")
            .withProblemClassificationId(
                "/providers/Microsoft.Support/services/quota_service_guid/problemClassifications/cosmosdb_problemClassification_guid")
            .withSeverity(SeverityLevel.MODERATE).withAdvancedDiagnosticConsent(Consent.YES)
            .withContactDetails(new ContactProfile().withFirstName("abc").withLastName("xyz")
                .withPreferredContactMethod(PreferredContactMethod.EMAIL).withPrimaryEmailAddress("abc@contoso.com")
                .withPreferredTimeZone("Pacific Standard Time").withCountry("usa")
                .withPreferredSupportLanguage("en-US"))
            .withTitle("my title").withServiceId("/providers/Microsoft.Support/services/quota_service_guid").create();
    }
}
