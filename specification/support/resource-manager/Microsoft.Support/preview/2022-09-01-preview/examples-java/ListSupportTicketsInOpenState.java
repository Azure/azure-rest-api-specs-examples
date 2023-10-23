/** Samples for SupportTicketsNoSubscription List. */
public final class Main {
    /*
     * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/ListSupportTicketsInOpenState.json
     */
    /**
     * Sample code: List support tickets in open state for a subscription.
     *
     * @param manager Entry point to SupportManager.
     */
    public static void listSupportTicketsInOpenStateForASubscription(
        com.azure.resourcemanager.support.SupportManager manager) {
        manager.supportTicketsNoSubscriptions().list(null, "status eq 'Open'", com.azure.core.util.Context.NONE);
    }
}
