import com.azure.resourcemanager.support.models.Consent;
import com.azure.resourcemanager.support.models.UpdateSupportTicket;

/** Samples for SupportTicketsNoSubscription Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/UpdateAdvancedDiagnosticConsentOfSupportTicket.json
     */
    /**
     * Sample code: Update advanced diagnostic consent of a support ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void updateAdvancedDiagnosticConsentOfASupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTicketsNoSubscriptions()
            .updateWithResponse(
                "testticket",
                new UpdateSupportTicket().withAdvancedDiagnosticConsent(Consent.YES),
                com.azure.core.util.Context.NONE);
    }
}
