
/**
 * Samples for ExadbVmClusters Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExadbVmClusters_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: ExadbVmClusters_Delete_MinimumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        exadbVmClustersDeleteMinimumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exadbVmClusters().delete("rgopenapi", "exadaVmClusterName1", com.azure.core.util.Context.NONE);
    }
}
