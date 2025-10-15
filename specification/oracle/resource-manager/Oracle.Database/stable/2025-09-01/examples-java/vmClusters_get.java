
/**
 * Samples for CloudVmClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/vmClusters_get.json
     */
    /**
     * Sample code: CloudVmClusters_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void cloudVmClustersGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().getByResourceGroupWithResponse("rg000", "cluster1", com.azure.core.util.Context.NONE);
    }
}
