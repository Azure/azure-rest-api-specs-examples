
/**
 * Samples for CommunicationsNoSubscription List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * ListCommunicationsForSupportTicket.json
     */
    /**
     * Sample code: List communications for a no-subscription support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        listCommunicationsForANoSubscriptionSupportTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager.communicationsNoSubscriptions().list("testticket", null, null, com.azure.core.util.Context.NONE);
    }
}
