
import com.azure.resourcemanager.support.models.Status;
import com.azure.resourcemanager.support.models.SupportTicketDetails;

/**
 * Samples for SupportTickets Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * UpdateStatusOfSupportTicketForSubscription.json
     */
    /**
     * Sample code: Update status of a subscription support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        updateStatusOfASubscriptionSupportTicket(com.azure.resourcemanager.support.SupportManager manager) {
        SupportTicketDetails resource
            = manager.supportTickets().getWithResponse("testticket", com.azure.core.util.Context.NONE).getValue();
        resource.update().withStatus(Status.CLOSED).apply();
    }
}
