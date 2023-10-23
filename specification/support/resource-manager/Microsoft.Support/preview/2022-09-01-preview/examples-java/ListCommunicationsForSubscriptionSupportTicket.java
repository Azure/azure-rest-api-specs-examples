/** Samples for Communications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListCommunicationsForSubscriptionSupportTicket.json
     */
    /**
     * Sample code: List communications for a subscription support ticket.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listCommunicationsForASubscriptionSupportTicket(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.communications().list("testticket", null, null, com.azure.core.util.Context.NONE);
    }
}
