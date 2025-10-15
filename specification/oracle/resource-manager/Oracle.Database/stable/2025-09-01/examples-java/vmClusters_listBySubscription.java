
/**
 * Samples for CloudVmClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/vmClusters_listBySubscription.json
     */
    /**
     * Sample code: CloudVmClusters_listBySubscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        cloudVmClustersListBySubscription(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().list(com.azure.core.util.Context.NONE);
    }
}
