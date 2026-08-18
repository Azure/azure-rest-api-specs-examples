
/**
 * Samples for Licenses List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15/license/License_ListBySubscription.json
     */
    /**
     * Sample code: List Licenses by Subscription.
     * 
     * @param manager Entry point to HybridComputeManager.
     */
    public static void
        listLicensesBySubscription(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.licenses().list(com.azure.core.util.Context.NONE);
    }
}
