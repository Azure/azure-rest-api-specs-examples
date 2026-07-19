
/**
 * Samples for LongTermRetentionPolicies ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/LongTermRetentionPolicyListByDatabase.json
     */
    /**
     * Sample code: Get the long term retention policy for the database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheLongTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionPolicies().listByDatabase("resourceGroup", "testserver",
            "testDatabase", com.azure.core.util.Context.NONE);
    }
}
