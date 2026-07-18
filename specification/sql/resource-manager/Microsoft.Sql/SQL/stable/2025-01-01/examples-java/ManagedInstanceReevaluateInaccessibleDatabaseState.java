
/**
 * Samples for ManagedInstances ReevaluateInaccessibleDatabaseState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceReevaluateInaccessibleDatabaseState.json
     */
    /**
     * Sample code: Reevaluate inaccessibility states of all managed databases.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        reevaluateInaccessibilityStatesOfAllManagedDatabases(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().reevaluateInaccessibleDatabaseState("testrg", "testinstance",
            com.azure.core.util.Context.NONE);
    }
}
