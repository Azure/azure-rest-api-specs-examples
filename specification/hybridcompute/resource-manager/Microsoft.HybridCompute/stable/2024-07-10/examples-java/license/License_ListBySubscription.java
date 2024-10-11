
/**
 * Samples for Licenses List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/license/
     * License_ListBySubscription.json
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
