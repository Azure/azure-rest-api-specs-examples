/** Samples for Communications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListWebCommunicationsForSubscriptionSupportTicket.json
     */
    /**
     * Sample code: List web communications for a subscription support ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listWebCommunicationsForASubscriptionSupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager
            .communications()
            .list("testticket", null, "communicationType eq 'web'", com.azure.core.util.Context.NONE);
    }
}
