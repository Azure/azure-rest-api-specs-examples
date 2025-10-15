
/**
 * Samples for CloudVmClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/vmClusters_listByResourceGroup.json
     */
    /**
     * Sample code: CloudVmClusters_listByResourceGroup.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        cloudVmClustersListByResourceGroup(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().listByResourceGroup("rg000", com.azure.core.util.Context.NONE);
    }
}
