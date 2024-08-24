
/**
 * Samples for LongTermRetentionPolicies ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/LongTermRetentionPolicyListByDatabase
     * .json
     */
    /**
     * Sample code: Get the long term retention policy for the database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getTheLongTermRetentionPolicyForTheDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getLongTermRetentionPolicies().listByDatabase("resourceGroup",
            "testserver", "testDatabase", com.azure.core.util.Context.NONE);
    }
}
