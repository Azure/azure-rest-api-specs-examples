
/**
 * Samples for ManagedInstanceLongTermRetentionPolicies ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedInstanceLongTermRetentionPolicyListByDatabase.json
     */
    /**
     * Sample code: Get the long term retention policies for the managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheLongTermRetentionPoliciesForTheManagedDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceLongTermRetentionPolicies().listByDatabase("testResourceGroup",
            "testInstance", "testDatabase", com.azure.core.util.Context.NONE);
    }
}
