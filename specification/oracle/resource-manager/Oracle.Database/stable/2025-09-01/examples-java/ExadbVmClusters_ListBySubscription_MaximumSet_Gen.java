
/**
 * Samples for ExadbVmClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExadbVmClusters_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExadbVmClusters_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void exadbVmClustersListBySubscriptionMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exadbVmClusters().list(com.azure.core.util.Context.NONE);
    }
}
