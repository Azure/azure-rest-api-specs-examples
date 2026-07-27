
/**
 * Samples for TrustedHostSubscriptions ListBySubscriptionLocationResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-31/TrustedHostSubscriptions_List.json
     */
    /**
     * Sample code: List trusted host subscriptions for a scope.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void
        listTrustedHostSubscriptionsForAScope(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.trustedHostSubscriptions().listBySubscriptionLocationResource("eastus",
            com.azure.core.util.Context.NONE);
    }
}
