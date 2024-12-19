
/**
 * Samples for SupportTicketsNoSubscription List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListSupportTickets.json
     */
    /**
     * Sample code: List support tickets.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void listSupportTickets(com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTicketsNoSubscriptions().list(null, null, com.azure.core.util.Context.NONE);
    }
}
