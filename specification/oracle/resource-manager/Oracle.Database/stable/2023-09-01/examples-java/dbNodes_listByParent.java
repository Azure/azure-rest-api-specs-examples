
/**
 * Samples for DbNodes ListByCloudVmCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dbNodes_listByParent.json
     */
    /**
     * Sample code: List DbNodes by VM Cluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void listDbNodesByVMCluster(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbNodes().listByCloudVmCluster("rg000", "cluster1", com.azure.core.util.Context.NONE);
    }
}
