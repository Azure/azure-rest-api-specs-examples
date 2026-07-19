
/**
 * Samples for ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies ListByRestorableDroppedDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetListManagedShortTermRetentionPolicyRestorableDropped.json
     */
    /**
     * Sample code: Get the short term retention policy list for the database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheShortTermRetentionPolicyListForTheDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies()
            .listByRestorableDroppedDatabase("Default-SQL-SouthEastAsia", "testsvr", "testdb,131403269876900000",
                com.azure.core.util.Context.NONE);
    }
}
