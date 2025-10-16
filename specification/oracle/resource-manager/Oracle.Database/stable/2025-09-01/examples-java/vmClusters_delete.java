
/**
 * Samples for CloudVmClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/vmClusters_delete.json
     */
    /**
     * Sample code: CloudVmClusters_delete.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void cloudVmClustersDelete(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().delete("rg000", "cluster1", com.azure.core.util.Context.NONE);
    }
}
