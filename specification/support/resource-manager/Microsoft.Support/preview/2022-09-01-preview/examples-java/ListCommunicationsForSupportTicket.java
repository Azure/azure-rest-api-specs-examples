/** Samples for SupportTicketCommunicationsNoSubscription List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListCommunicationsForSupportTicket.json
     */
    /**
     * Sample code: List communications for a no-subscription support ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listCommunicationsForANoSubscriptionSupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTicketCommunicationsNoSubscriptions()
            .list("testticket", null, null, com.azure.core.util.Context.NONE);
    }
}
