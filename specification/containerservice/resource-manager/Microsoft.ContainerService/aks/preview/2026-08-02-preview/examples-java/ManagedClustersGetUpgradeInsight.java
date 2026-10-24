
/**
 * Samples for UpgradeInsights Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/ManagedClustersGetUpgradeInsight.json
     */
    /**
     * Sample code: Get a stored upgrade insight.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getAStoredUpgradeInsight(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getUpgradeInsights().getWithResponse("rg1", "clustername1", "controlPlane-1.33.5",
            com.azure.core.util.Context.NONE);
    }
}
