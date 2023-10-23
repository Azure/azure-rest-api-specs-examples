/** Samples for SupportTicketCommunicationsNoSubscription List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListWebCommunicationsForSupportTicket.json
     */
    /**
     * Sample code: List web communications for a no-subscription support ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listWebCommunicationsForANoSubscriptionSupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .supportTicketCommunicationsNoSubscriptions()
            .list("testticket", null, "communicationType eq 'web'", com.azure.core.util.Context.NONE);
    }
}
