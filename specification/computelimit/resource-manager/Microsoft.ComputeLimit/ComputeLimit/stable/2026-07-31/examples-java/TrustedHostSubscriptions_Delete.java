
/**
 * Samples for TrustedHostSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-31/TrustedHostSubscriptions_Delete.json
     */
    /**
     * Sample code: Revoke trust in a host subscription.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void
        revokeTrustInAHostSubscription(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.trustedHostSubscriptions().deleteByResourceGroupWithResponse("eastus",
            "22222222-2222-2222-2222-222222222222", com.azure.core.util.Context.NONE);
    }
}
