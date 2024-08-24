
/**
 * Samples for ManagedDatabaseSecurityAlertPolicies ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedDatabaseSecurityAlertListByDatabase.json
     */
    /**
     * Sample code: Get a list of the database's threat detection policies.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAListOfTheDatabaseSThreatDetectionPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSecurityAlertPolicies()
            .listByDatabase("securityalert-6852", "securityalert-2080", "testdb", com.azure.core.util.Context.NONE);
    }
}
