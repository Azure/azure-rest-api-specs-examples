import com.azure.resourcemanager.support.models.SeverityLevel;
import com.azure.resourcemanager.support.models.UpdateSupportTicket;

/** Samples for SupportTicketsNoSubscription Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/UpdateSeverityOfSupportTicket.json
     */
    /**
     * Sample code: Update severity of a support ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void updateSeverityOfASupportTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTicketsNoSubscriptions()
            .updateWithResponse(
                "testticket",
                new UpdateSupportTicket().withSeverity(SeverityLevel.CRITICAL),
                com.azure.core.util.Context.NONE);
    }
}
