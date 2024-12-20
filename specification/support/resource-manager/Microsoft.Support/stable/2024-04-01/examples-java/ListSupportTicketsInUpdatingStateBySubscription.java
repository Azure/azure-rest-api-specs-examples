
/**
 * Samples for SupportTickets List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * ListSupportTicketsInUpdatingStateBySubscription.json
     */
    /**
     * Sample code: List support tickets in updating state for a subscription.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void
        listSupportTicketsInUpdatingStateForASubscription(com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTickets().list(null, "status eq 'Updating'", com.azure.core.util.Context.NONE);
    }
}
