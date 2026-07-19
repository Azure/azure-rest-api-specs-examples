
/**
 * Samples for ManagedDatabaseSecurityAlertPolicies ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseSecurityAlertListByDatabase.json
     */
    /**
     * Sample code: Get a list of the database's threat detection policies.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAListOfTheDatabaseSThreatDetectionPolicies(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSecurityAlertPolicies().listByDatabase("securityalert-6852",
            "securityalert-2080", "testdb", com.azure.core.util.Context.NONE);
    }
}
