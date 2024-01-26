
import com.azure.core.util.Context;

/** Samples for DatabaseSecurityAlertPolicies ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseSecurityAlertListByDatabase.
     * json
     */
    /**
     * Sample code: Get the database's threat detection policies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheDatabaseSThreatDetectionPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseSecurityAlertPolicies()
            .listByDatabase("securityalert-6852", "securityalert-2080", "testdb", Context.NONE);
    }
}
