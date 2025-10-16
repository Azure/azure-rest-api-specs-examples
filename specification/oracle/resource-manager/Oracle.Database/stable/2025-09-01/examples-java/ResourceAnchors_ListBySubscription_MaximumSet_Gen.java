
/**
 * Samples for ResourceAnchors List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ResourceAnchors_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: ResourceAnchors_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void resourceAnchorsListBySubscriptionMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.resourceAnchors().list(com.azure.core.util.Context.NONE);
    }
}
