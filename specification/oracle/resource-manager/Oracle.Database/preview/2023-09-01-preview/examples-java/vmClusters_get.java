
/**
 * Samples for CloudVmClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/vmClusters_get.json
     */
    /**
     * Sample code: Get VM Cluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void getVMCluster(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().getByResourceGroupWithResponse("rg000", "cluster1", com.azure.core.util.Context.NONE);
    }
}
