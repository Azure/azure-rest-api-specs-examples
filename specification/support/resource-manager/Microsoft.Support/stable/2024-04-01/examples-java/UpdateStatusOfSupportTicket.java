
import com.azure.resourcemanager.support.models.Status;
import com.azure.resourcemanager.support.models.UpdateSupportTicket;

/**
 * Samples for SupportTicketsNoSubscription Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/UpdateStatusOfSupportTicket.
     * json
     */
    /**
     * Sample code: Update status of a support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void updateStatusOfASupportTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTicketsNoSubscriptions().updateWithResponse("testticket",
            new UpdateSupportTicket().withStatus(Status.CLOSED), com.azure.core.util.Context.NONE);
    }
}
