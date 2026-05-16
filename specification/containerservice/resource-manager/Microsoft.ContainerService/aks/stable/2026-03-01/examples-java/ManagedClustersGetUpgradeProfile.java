
/**
 * Samples for ManagedClusters GetUpgradeProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ManagedClustersGetUpgradeProfile.json
     */
    /**
     * Sample code: Get Upgrade Profile for Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getUpgradeProfileForManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().getUpgradeProfileWithResponse("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
