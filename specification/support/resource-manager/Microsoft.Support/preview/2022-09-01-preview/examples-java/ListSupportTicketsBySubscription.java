/** Samples for SupportTickets List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListSupportTicketsBySubscription.json
     */
    /**
     * Sample code: List support tickets for a subscription.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listSupportTicketsForASubscription(com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().list(null, null, com.azure.core.util.Context.NONE);
    }
}
