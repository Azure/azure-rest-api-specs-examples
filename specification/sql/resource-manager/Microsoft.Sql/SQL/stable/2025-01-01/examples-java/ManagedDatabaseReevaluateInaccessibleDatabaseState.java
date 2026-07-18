
/**
 * Samples for ManagedDatabases ReevaluateInaccessibleDatabaseState.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseReevaluateInaccessibleDatabaseState.json
     */
    /**
     * Sample code: Reevaluate the inaccessibility state of a managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        reevaluateTheInaccessibilityStateOfAManagedDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabases().reevaluateInaccessibleDatabaseState("Test1", "managedInstance",
            "managedDatabase", com.azure.core.util.Context.NONE);
    }
}
