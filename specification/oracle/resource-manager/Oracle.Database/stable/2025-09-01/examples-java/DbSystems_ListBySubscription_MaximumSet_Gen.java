
/**
 * Samples for DbSystems List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/DbSystems_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: DbSystems_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        dbSystemsListBySubscriptionMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbSystems().list(com.azure.core.util.Context.NONE);
    }
}
