
/**
 * Samples for VirtualClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/VirtualClusterDelete.json
     */
    /**
     * Sample code: Delete virtual cluster.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteVirtualCluster(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getVirtualClusters().delete("testrg", "vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2",
            com.azure.core.util.Context.NONE);
    }
}
