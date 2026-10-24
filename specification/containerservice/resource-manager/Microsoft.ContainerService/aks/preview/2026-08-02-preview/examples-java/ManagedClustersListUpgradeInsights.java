
/**
 * Samples for UpgradeInsights List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/ManagedClustersListUpgradeInsights.json
     */
    /**
     * Sample code: List the stored upgrade insights for a managed cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listTheStoredUpgradeInsightsForAManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getUpgradeInsights().list("rg1", "clustername1", com.azure.core.util.Context.NONE);
    }
}
