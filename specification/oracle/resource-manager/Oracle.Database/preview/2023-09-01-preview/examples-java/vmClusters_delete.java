
/**
 * Samples for CloudVmClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/vmClusters_delete.json
     */
    /**
     * Sample code: Delete VM Cluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void deleteVMCluster(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().delete("rg000", "cluster1", com.azure.core.util.Context.NONE);
    }
}
