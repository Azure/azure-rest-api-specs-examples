
/**
 * Samples for BackupShortTermRetentionPolicies ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListShortTermRetentionPoliciesByDatabase.json
     */
    /**
     * Sample code: Get the short term retention policy for the database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheShortTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getBackupShortTermRetentionPolicies().listByDatabase("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", com.azure.core.util.Context.NONE);
    }
}
