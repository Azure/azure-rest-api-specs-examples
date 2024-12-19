
/**
 * Samples for SupportTicketsNoSubscription List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/
     * ListSupportTicketsInUpdatingState.json
     */
    /**
     * Sample code: List support tickets in updating state.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void listSupportTicketsInUpdatingState(com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTicketsNoSubscriptions().list(null, "status eq 'Updating'", com.azure.core.util.Context.NONE);
    }
}
