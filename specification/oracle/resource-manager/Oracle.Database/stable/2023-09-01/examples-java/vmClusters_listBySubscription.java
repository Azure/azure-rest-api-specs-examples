
/**
 * Samples for CloudVmClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/vmClusters_listBySubscription.
     * json
     */
    /**
     * Sample code: List VM Clusters by subscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listVMClustersBySubscription(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().list(com.azure.core.util.Context.NONE);
    }
}
