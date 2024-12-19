
/**
 * Samples for CommunicationsNoSubscription List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * ListWebCommunicationsForSupportTicket.json
     */
    /**
     * Sample code: List web communications for a no-subscription support ticket.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        listWebCommunicationsForANoSubscriptionSupportTicket(com.azure.resourcemanager.support.SupportManager manager) {
        manager.communicationsNoSubscriptions().list("testticket", null, "communicationType eq 'web'",
            com.azure.core.util.Context.NONE);
    }
}
