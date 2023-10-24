import com.azure.resourcemanager.support.models.SeverityLevel;
import com.azure.resourcemanager.support.models.SupportTicketDetails;

/** Samples for SupportTickets Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/UpdateSeverityOfSupportTicketForSubscription.json
     */
    /**
     * Sample code: Update severity of a support ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void updateSeverityOfASupportTicket(com.azure.resourcemanager.support.SupportManager manager) {
        SupportTicketDetails resource =
            manager.supportTickets().getWithResponse("testticket", com.azure.core.util.Context.NONE).getValue();
        resource.update().withSeverity(SeverityLevel.CRITICAL).apply();
    }
}
