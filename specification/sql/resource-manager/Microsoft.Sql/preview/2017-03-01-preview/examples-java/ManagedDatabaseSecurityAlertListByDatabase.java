import com.azure.core.util.Context;

/** Samples for ManagedDatabaseSecurityAlertPolicies ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ManagedDatabaseSecurityAlertListByDatabase.json
     */
    /**
     * Sample code: Get a list of the database's threat detection policies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAListOfTheDatabaseSThreatDetectionPolicies(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedDatabaseSecurityAlertPolicies()
            .listByDatabase("securityalert-6852", "securityalert-2080", "testdb", Context.NONE);
    }
}
