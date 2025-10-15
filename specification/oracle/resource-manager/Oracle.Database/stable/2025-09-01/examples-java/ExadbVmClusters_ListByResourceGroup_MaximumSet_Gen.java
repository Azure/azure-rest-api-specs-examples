
/**
 * Samples for ExadbVmClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExadbVmClusters_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExadbVmClusters_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void exadbVmClustersListByResourceGroupMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exadbVmClusters().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
