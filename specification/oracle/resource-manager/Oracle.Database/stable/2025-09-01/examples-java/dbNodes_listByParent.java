
/**
 * Samples for DbNodes ListByCloudVmCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/dbNodes_listByParent.json
     */
    /**
     * Sample code: DbNodes_listByCloudVmCluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        dbNodesListByCloudVmCluster(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbNodes().listByCloudVmCluster("rg000", "cluster1", com.azure.core.util.Context.NONE);
    }
}
