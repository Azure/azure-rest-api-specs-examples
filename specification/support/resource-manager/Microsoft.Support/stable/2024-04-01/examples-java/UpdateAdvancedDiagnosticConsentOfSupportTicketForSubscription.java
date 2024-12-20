
import com.azure.resourcemanager.support.models.Consent;
import com.azure.resourcemanager.support.models.SupportTicketDetails;

/**
 * Samples for SupportTickets Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * UpdateAdvancedDiagnosticConsentOfSupportTicketForSubscription.json
     */
    /**
     * Sample code: Update advanced diagnostic consent of a subscription support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void updateAdvancedDiagnosticConsentOfASubscriptionSupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        SupportTicketDetails resource
            = manager.supportTickets().getWithResponse("testticket", com.azure.core.util.Context.NONE).getValue();
        resource.update().withAdvancedDiagnosticConsent(Consent.YES).apply();
    }
}
