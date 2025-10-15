
/**
 * Samples for ExadbVmClusters GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExadbVmClusters_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExadbVmClusters_Get_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        exadbVmClustersGetMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exadbVmClusters().getByResourceGroupWithResponse("rgopenapi", "exadbVmClusterName1",
            com.azure.core.util.Context.NONE);
    }
}
