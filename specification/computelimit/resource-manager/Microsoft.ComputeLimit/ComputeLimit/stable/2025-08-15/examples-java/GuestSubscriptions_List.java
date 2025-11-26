
/**
 * Samples for GuestSubscriptions ListBySubscriptionLocationResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-15/GuestSubscriptions_List.json
     */
    /**
     * Sample code: List guest subscriptions for a scope.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void
        listGuestSubscriptionsForAScope(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.guestSubscriptions().listBySubscriptionLocationResource("eastus", com.azure.core.util.Context.NONE);
    }
}
