
/**
 * Samples for ExascaleDbStorageVaults List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExascaleDbStorageVaults_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExascaleDbStorageVaults_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void exascaleDbStorageVaultsListBySubscriptionMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exascaleDbStorageVaults().list(com.azure.core.util.Context.NONE);
    }
}
